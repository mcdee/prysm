package validator

import (
	"context"
	"errors"
	"math/big"
	"time"

	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/helpers"
	pbp2p "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	"github.com/prysmaticlabs/prysm/shared/params"
	"github.com/prysmaticlabs/prysm/shared/traceutil"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errPubkeyDoesNotExist = errors.New("pubkey does not exist")

// ValidatorStatus returns the validator status of the current epoch.
// The status response can be one of the following:
//	PENDING_ACTIVE - validator is waiting to get activated.
//	ACTIVE - validator is active.
//	INITIATED_EXIT - validator has initiated an an exit request.
//	WITHDRAWABLE - validator's deposit can be withdrawn after lock up period.
//	EXITED - validator has exited, means the deposit has been withdrawn.
//	EXITED_SLASHED - validator was forcefully exited due to slashing.
func (vs *Server) ValidatorStatus(
	ctx context.Context,
	req *ethpb.ValidatorStatusRequest) (*ethpb.ValidatorStatusResponse, error) {
	headState, err := vs.HeadFetcher.HeadState(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Could not get head state")
	}
	return vs.validatorStatus(ctx, req.PublicKey, headState), nil
}

// multipleValidatorStatus returns the validator status response for the set of validators
// requested by their pub keys.
func (vs *Server) multipleValidatorStatus(
	ctx context.Context,
	pubkeys [][]byte,
) (bool, []*ethpb.ValidatorActivationResponse_Status, error) {
	headState, err := vs.HeadFetcher.HeadState(ctx)
	if err != nil {
		return false, nil, err
	}
	activeValidatorExists := false
	statusResponses := make([]*ethpb.ValidatorActivationResponse_Status, len(pubkeys))
	for i, key := range pubkeys {
		if ctx.Err() != nil {
			return false, nil, ctx.Err()
		}
		status := vs.validatorStatus(ctx, key, headState)
		if status == nil {
			continue
		}
		resp := &ethpb.ValidatorActivationResponse_Status{
			Status:    status,
			PublicKey: key,
		}
		statusResponses[i] = resp
		if status.Status == ethpb.ValidatorState_ACTIVE {
			activeValidatorExists = true
		}
	}

	return activeValidatorExists, statusResponses, nil
}

func (vs *Server) validatorStatus(ctx context.Context, pubKey []byte, headState *pbp2p.BeaconState) *ethpb.ValidatorStatusResponse {
	ctx, span := trace.StartSpan(ctx, "validatorServer.validatorStatus")
	defer span.End()

	resp := &ethpb.ValidatorStatusResponse{
		Status:          ethpb.ValidatorState_UNKNOWN_STATE,
		ActivationEpoch: int64(params.BeaconConfig().FarFutureEpoch),
	}
	vStatus, idx, err := vs.retrieveStatusFromState(ctx, pubKey, headState)
	if err != nil && err != errPubkeyDoesNotExist {
		traceutil.AnnotateError(span, err)
		return resp
	}
	resp.Status = vStatus
	if err != errPubkeyDoesNotExist {
		resp.ActivationEpoch = int64(headState.Validators[idx].ActivationEpoch)
	}

	// If no connection to ETH1, the deposit block number or position in queue cannot be determined.
	if !vs.Eth1InfoFetcher.IsConnectedToETH1() {
		log.Warn("Not connected to ETH1. Cannot determine validator ETH1 deposit block number")
		return resp
	}

	_, eth1BlockNumBigInt := vs.DepositFetcher.DepositByPubkey(ctx, pubKey)
	if eth1BlockNumBigInt == nil { // No deposit found in ETH1.
		return resp
	}

	// TODO really?
	if resp.Status == ethpb.ValidatorState_UNKNOWN_STATE {
		resp.Status = ethpb.ValidatorState_PENDING
	}

	resp.Eth1DepositBlockNumber = eth1BlockNumBigInt.Uint64()

	depositBlockSlot, err := vs.depositBlockSlot(ctx, eth1BlockNumBigInt, headState)
	if err != nil {
		return resp
	}
	resp.DepositInclusionSlot = int64(depositBlockSlot)

	// If validator has been activated at any point, they are not in the queue so we can return
	// the request early. Additionally, if idx is zero (default return value) then we know this
	// validator cannot be in the queue either.
	if uint64(resp.ActivationEpoch) != params.BeaconConfig().FarFutureEpoch || idx == 0 {
		return resp
	}

	var lastActivatedValidatorIdx uint64
	for j := len(headState.Validators) - 1; j >= 0; j-- {
		if helpers.IsActiveValidator(headState.Validators[j], helpers.CurrentEpoch(headState)) {
			lastActivatedValidatorIdx = uint64(j)
			break
		}
	}
	// Our position in the activation queue is the above index - our validator index.
	if lastActivatedValidatorIdx > idx {
		resp.PositionInActivationQueue = int64(idx - lastActivatedValidatorIdx)
	}

	return resp
}

func (vs *Server) retrieveStatusFromState(
	ctx context.Context,
	pubKey []byte,
	headState *pbp2p.BeaconState,
) (ethpb.ValidatorState, uint64, error) {
	if headState == nil {
		return ethpb.ValidatorState(0), 0, errors.New("head state does not exist")
	}
	idx, ok, err := vs.BeaconDB.ValidatorIndex(ctx, pubKey)
	if err != nil {
		return ethpb.ValidatorState(0), 0, err
	}
	if !ok || int(idx) >= len(headState.Validators) {
		return ethpb.ValidatorState(0), 0, errPubkeyDoesNotExist
	}
	return vs.calculateStatus(idx, headState), idx, nil
}

func (vs *Server) calculateStatus(validatorIdx uint64, beaconState *pbp2p.BeaconState) ethpb.ValidatorState {
	var status ethpb.ValidatorState

	validator := beaconState.Validators[validatorIdx]
	currentEpoch := helpers.CurrentEpoch(beaconState)
	if validator.Slashed {
		if currentEpoch < validator.ExitEpoch {
			status = ethpb.ValidatorState_SLASHING
		} else {
			status = ethpb.ValidatorState_SLASHED
		}
	} else if validator.ExitEpoch != params.BeaconConfig().FarFutureEpoch {
		if currentEpoch < validator.ExitEpoch {
			status = ethpb.ValidatorState_EXITING
		} else {
			status = ethpb.ValidatorState_EXITED
		}
	} else if currentEpoch < validator.ActivationEpoch {
		status = ethpb.ValidatorState_PENDING
	} else {
		status = ethpb.ValidatorState_ACTIVE
	}

	return status
}

func (vs *Server) depositBlockSlot(ctx context.Context, eth1BlockNumBigInt *big.Int, beaconState *pbp2p.BeaconState) (uint64, error) {
	blockTimeStamp, err := vs.BlockFetcher.BlockTimeByHeight(ctx, eth1BlockNumBigInt)
	if err != nil {
		return 0, err
	}
	followTime := time.Duration(params.BeaconConfig().Eth1FollowDistance*params.BeaconConfig().GoerliBlockTime) * time.Second
	eth1UnixTime := time.Unix(int64(blockTimeStamp), 0).Add(followTime)

	votingPeriodSlots := helpers.StartSlot(params.BeaconConfig().SlotsPerEth1VotingPeriod / params.BeaconConfig().SlotsPerEpoch)
	votingPeriodSeconds := time.Duration(votingPeriodSlots*params.BeaconConfig().SecondsPerSlot) * time.Second
	timeToInclusion := eth1UnixTime.Add(votingPeriodSeconds)

	eth2Genesis := time.Unix(int64(beaconState.GenesisTime), 0)
	eth2TimeDifference := timeToInclusion.Sub(eth2Genesis).Seconds()
	depositBlockSlot := uint64(eth2TimeDifference) / params.BeaconConfig().SecondsPerSlot

	return depositBlockSlot, nil
}
