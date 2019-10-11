package attestant

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/prysmaticlabs/prysm/beacon-chain/blockchain"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/helpers"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	eth "github.com/prysmaticlabs/prysm/proto/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/shared/params"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("prefix", "attestant")

// Service defining metrics functionality for storing metrics in a database backend.
type Service struct {
	ctx                context.Context
	cancel             context.CancelFunc
	genesisTimeFetcher blockchain.GenesisTimeFetcher
	headFetcher        blockchain.HeadFetcher
	newHeadNotifier    blockchain.NewHeadNotifier
	newHeadRootChan    chan [32]byte
	db                 *sql.DB
	lastEpoch          uint64
}

// Config options for the attestant service.
type Config struct {
	GenesisTimeFetcher blockchain.GenesisTimeFetcher
	HeadFetcher        blockchain.HeadFetcher
	NewHeadNotifier    blockchain.NewHeadNotifier
}

// NewService initializes the service from configuration options.
func NewService(ctx context.Context, cfg *Config) *Service {
	var err error
	db, err := sql.Open("postgres", "postgres://attestant:attestant@localhost:5432/attestant?sslmode=disable")
	if err != nil {
		log.WithError(err).Error("Failed to open attestant DB")
		return nil
	}

	// TODO obtain blocks from prior slots so we have stats for this epoch to date

	ctx, cancel := context.WithCancel(ctx)
	return &Service{
		db:                 db,
		ctx:                ctx,
		cancel:             cancel,
		genesisTimeFetcher: cfg.GenesisTimeFetcher,
		headFetcher:        cfg.HeadFetcher,
		newHeadNotifier:    cfg.NewHeadNotifier,
		newHeadRootChan:    make(chan [32]byte, 1),
	}
}

// Start the metrics service event loop.
func (s *Service) Start() {
	go s.run(s.ctx)
}

// Stop the metrics service event loop.
func (s *Service) Stop() error {
	defer s.cancel()
	return nil
}

// Status reports the healthy status of the metrics. Returning nil means service
// is correctly running without error.
func (s *Service) Status() error {
	return nil
}

// updateValidatorStats updates the validator stats.
// Validator stats are per-epoch.
func (s *Service) updateValidatorStats(headState *pb.BeaconState) error {
	epoch := helpers.SlotToEpoch(headState.Slot)
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stats := &epochStats{
		epoch: epoch,
	}

	// Per-validator
	for i, validator := range headState.Validators {
		validatorStats := &validatorStats{
			epoch:            epoch,
			id:               i,
			balance:          headState.Balances[i],
			effectiveBalance: validator.EffectiveBalance,
		}
		if validator.Slashed {
			if epoch < validator.ExitEpoch {
				validatorStats.state = "slashing"
				stats.slashingInstances++
				stats.slashingBalance += headState.Balances[i]
				stats.slashingEffectiveBalance += validator.EffectiveBalance
			} else {
				validatorStats.state = "slashed"
				stats.slashedInstances++
			}
		} else if validator.ExitEpoch != params.BeaconConfig().FarFutureEpoch {
			if epoch < validator.ExitEpoch {
				validatorStats.state = "exiting"
				stats.exitingInstances++
				stats.exitingBalance += headState.Balances[i]
				stats.exitingEffectiveBalance += validator.EffectiveBalance
			} else {
				validatorStats.state = "exited"
				stats.exitedInstances++
			}
		} else if epoch < validator.ActivationEpoch {
			validatorStats.state = "pending"
			stats.pendingInstances++
			stats.pendingBalance += headState.Balances[i]
		} else {
			validatorStats.state = "active"
			stats.activeInstances++
			stats.activeBalance += headState.Balances[i]
			stats.activeEffectiveBalance += validator.EffectiveBalance
		}
		err = s.logValidatorStats(tx, validatorStats)
		if err != nil {
			return err
		}
	}
	err = s.logEpochStats(tx, stats)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Service) updateBlockStats(headState *pb.BeaconState, blockHash [32]byte, block *eth.BeaconBlock) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	body := block.GetBody()
	stats := &blockStats{
		slot:              headState.Slot,
		hash:              blockHash,
		attestations:      len(body.GetAttestations()),
		deposits:          len(body.GetDeposits()),
		transfers:         len(body.GetTransfers()),
		exits:             len(body.GetVoluntaryExits()),
		attesterSlashings: len(body.GetAttesterSlashings()),
		proposerSlashings: len(body.GetProposerSlashings()),
	}

	err = s.logBlockStats(tx, stats)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Service) run(ctx context.Context) {
	sub := s.newHeadNotifier.HeadUpdatedFeed().Subscribe(s.newHeadRootChan)
	defer sub.Unsubscribe()
	for {
		select {
		case r := <-s.newHeadRootChan:
			log.WithField("headRoot", fmt.Sprintf("%#x", r)).Debug("New chain head event")
			headState := s.headFetcher.HeadState()

			// Update statistics
			err := s.updateBlockStats(headState, r, s.headFetcher.HeadBlock())
			if err != nil {
				log.WithError(err).Warn("failed to update block stats")
			}

			if helpers.IsEpochStart(headState.Slot) {
				err = s.updateValidatorStats(headState)
				if err != nil {
					log.WithError(err).Warn("failed to update validator stats")
				}
			}
			log.WithField("epoch", helpers.CurrentEpoch(headState)).Debug("Updated metrics")
		case <-s.ctx.Done():
			log.Debug("Context closed, exiting goroutine")
			return
		case err := <-sub.Err():
			log.WithError(err).Error("Subscription to new chain head notifier failed")
			return
		}
	}
}
