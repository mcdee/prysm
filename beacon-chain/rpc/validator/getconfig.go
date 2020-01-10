package validator

import (
	"context"
	"math/big"

	ptypes "github.com/gogo/protobuf/types"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetConfig gets the configuration for the validators.
func (vs *Server) GetConfig(ctx context.Context, req *ethpb.GetValidatorConfigRequest) (*ethpb.GetValidatorConfigResponse, error) {
	headState, err := vs.HeadFetcher.HeadState(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Could not get head state")
	}

	validators := make([]*ethpb.ValidatorConfig, 0)

	if len(req.PublicKeys) == 0 {
		return nil, status.Error(codes.Internal, "Not implemented (all pubkeys)")
	}

	for _, pubKey := range req.PublicKeys {
		validator := &ethpb.ValidatorConfig{
			PublicKey: pubKey,
		}

		// Eth1DepositBlockNumber
		if vs.Eth1InfoFetcher.IsConnectedToETH1() {
			_, eth1BlockNumBigInt := vs.DepositFetcher.DepositByPubkey(ctx, pubKey)
			//			if eth1BlockNumBigInt != nil {
			validator.Eth1DepositBlockNumber = eth1BlockNumBigInt.Uint64()
			//			}
		}

		// Index
		idx, ok, err := vs.BeaconDB.ValidatorIndex(ctx, pubKey)
		if err != nil {
			// TODO log
			return nil, status.Error(codes.Internal, "Failed to obtain validator index")
		}
		if !ok && validator.Eth1DepositBlockNumber == 0 {
			// TODO more user-friendly error messages everywhere
			return nil, status.Error(codes.NotFound, "Unknown validator")
		}
		validator.Index = idx

		// DepositInclusionSlot
		if validator.Eth1DepositBlockNumber != 0 {
			depositBlockSlot, err := vs.depositBlockSlot(ctx, big.NewInt(int64(validator.Eth1DepositBlockNumber)), headState)
			if err == nil {
				validator.DepositInclusionSlot = int64(depositBlockSlot)
			}
		}

		// ActivationEpoch
		validator.ActivationEpoch = int64(headState.Validators[idx].ActivationEpoch)

		// WithdrawalEpoch
		validator.WithdrawalEpoch = int64(headState.Validators[idx].WithdrawableEpoch)

		// TODO Withdrawal credentials

		validators = append(validators, validator)
	}

	resp := &ethpb.GetValidatorConfigResponse{Validators: validators}
	return resp, nil
}

// GetConfigStream returns a stream of new configurations.
func (vs *Server) GetConfigStream(*ptypes.Empty, ethpb.BeaconNodeValidator_GetConfigStreamServer) error {
	return status.Error(codes.Internal, "Not implemented")
}
