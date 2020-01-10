package validator

import (
	"context"
	"fmt"

	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/feed"
	statefeed "github.com/prysmaticlabs/prysm/beacon-chain/core/feed/state"
	"github.com/prysmaticlabs/prysm/shared/params"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetInfo returns information for the validators.
func (vs *Server) GetInfo(ctx context.Context, req *ethpb.GetValidatorInfoRequest) (*ethpb.GetValidatorInfoResponse, error) {
	return vs.generateValidatorInfoResponse(ctx, req.PublicKeys)
}

// GetInfoStream returns a stream of information for the validators.
func (vs *Server) GetInfoStream(req *ethpb.GetValidatorInfoStreamRequest, stream ethpb.BeaconNodeValidator_GetInfoStreamServer) error {
	stateChannel := make(chan *feed.Event, 1)
	stateSub := vs.StateNotifier.StateFeed().Subscribe(stateChannel)
	defer stateSub.Unsubscribe()

	headState, err := vs.HeadFetcher.HeadState(vs.Ctx)
	if err != nil {
		return status.Error(codes.Internal, "Could not access head state")
	}
	if headState == nil {
		return status.Error(codes.Internal, "Not ready to serve information")
	}
	currentEpoch := headState.Slot / params.BeaconConfig().SlotsPerEpoch
	// Send a response for the current epoch immediately.
	resp, err := vs.generateValidatorInfoResponse(vs.Ctx, req.PublicKeys)
	if err != nil {
		return status.Errorf(codes.Internal, "Could not generate response: %v", err)
	}
	for _, validator := range resp.Validators {
		if err := stream.Send(validator); err != nil {
			return status.Errorf(codes.Unavailable, "Could not send over stream: %v", err)
		}
	}

	// Send additional responses every epoch.
	for {
		select {
		case event := <-stateChannel:
			if event.Type == statefeed.BlockProcessed {
				headState, err := vs.HeadFetcher.HeadState(vs.Ctx)
				if err != nil {
					return status.Error(codes.Internal, "Could not access head state")
				}
				if headState == nil {
					return status.Error(codes.Internal, "Not ready to serve information")
				}
				blockEpoch := headState.Slot / params.BeaconConfig().SlotsPerEpoch
				if blockEpoch == currentEpoch {
					// Epoch hasn't changed, nothing to report yet.
					continue
				}
				currentEpoch = blockEpoch
				resp, err := vs.generateValidatorInfoResponse(vs.Ctx, req.PublicKeys)
				if err != nil {
					return status.Errorf(codes.Internal, "Could not generate response: %v", err)
				}
				for _, validator := range resp.Validators {
					if err := stream.Send(validator); err != nil {
						return status.Errorf(codes.Unavailable, "Could not send over stream: %v", err)
					}
				}
			}
		case <-stateSub.Err():
			return status.Error(codes.Aborted, "Subscriber closed")
		case <-vs.Ctx.Done():
			return status.Error(codes.Canceled, "Service context canceled")
		case <-stream.Context().Done():
			return status.Error(codes.Canceled, "Stream context canceled")
		}
	}
}

func (vs *Server) generateValidatorInfoResponse(ctx context.Context, pubKeys [][]byte) (*ethpb.GetValidatorInfoResponse, error) {
	headState, err := vs.HeadFetcher.HeadState(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Could not access head state")
	}
	if headState == nil {
		return nil, status.Error(codes.Internal, "Not ready to serve information")
	}

	if len(pubKeys) == 0 {
		for _, validator := range headState.Validators {
			pubKeys = append(pubKeys, validator.PublicKey)
		}
	}

	validators := make([]*ethpb.ValidatorInfo, 0)
	for _, pubKey := range pubKeys {
		validator := &ethpb.ValidatorInfo{
			PublicKey: pubKey,
		}

		// Index
		validatorIndex, ok, err := vs.BeaconDB.ValidatorIndex(ctx, pubKey)
		if err != nil {
			// TODO log
			return nil, status.Error(codes.Internal, "Failed to obtain validator index")
		}
		if !ok {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("Unknown validator with public key %#x", pubKey))
		}

		// State
		validator.State = vs.calculateStatus(validatorIndex, headState)

		// Balance (for relevant states)
		if validator.State == ethpb.ValidatorState_PENDING ||
			validator.State == ethpb.ValidatorState_ACTIVE ||
			validator.State == ethpb.ValidatorState_SLASHING ||
			validator.State == ethpb.ValidatorState_EXITING {
			validator.Balance = headState.Balances[validatorIndex]
		}

		// Effective balance (for relevant states)
		if validator.State == ethpb.ValidatorState_ACTIVE ||
			validator.State == ethpb.ValidatorState_SLASHING ||
			validator.State == ethpb.ValidatorState_EXITING {
			validator.EffectiveBalance = headState.Validators[validatorIndex].EffectiveBalance
		}

		// Attested?
		// Proposed?
		validators = append(validators, validator)
	}
	resp := &ethpb.GetValidatorInfoResponse{
		Validators: validators,
	}
	return resp, nil
}
