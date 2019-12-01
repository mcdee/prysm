package attestant

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/beacon-chain/blockchain"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/helpers"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/statefeed"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	"github.com/prysmaticlabs/prysm/shared/params"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("prefix", "attestant")

// Service defining metrics functionality for storing metrics in a database backend.
type Service struct {
	mu                 *sync.RWMutex
	ctx                context.Context
	cancel             context.CancelFunc
	genesisTimeFetcher blockchain.GenesisTimeFetcher
	headFetcher        blockchain.HeadFetcher
	stateNotifier      statefeed.Notifier
	db                 *sql.DB
	currentEpoch       int64
	// Map of validator indices that have had attestations included in blocks this epoch
	epochAttestations map[uint64]bool
}

// Config options for the attestant service.
type Config struct {
	GenesisTimeFetcher blockchain.GenesisTimeFetcher
	HeadFetcher        blockchain.HeadFetcher
	StateNotifier      statefeed.Notifier
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
		mu:                 new(sync.RWMutex),
		db:                 db,
		ctx:                ctx,
		cancel:             cancel,
		genesisTimeFetcher: cfg.GenesisTimeFetcher,
		headFetcher:        cfg.HeadFetcher,
		stateNotifier:      cfg.StateNotifier,
		currentEpoch:       -1,
		epochAttestations:  make(map[uint64]bool),
	}
}

// Start the metrics service event loop.
func (s *Service) Start() {
	go s.run(s.ctx)
	go s.poll(s.ctx)
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

// run is the main service loop.
func (s *Service) run(ctx context.Context) {
	stateChannel := make(chan *statefeed.Event, 1)
	stateSub := s.stateNotifier.StateFeed().Subscribe(stateChannel)
	defer stateSub.Unsubscribe()
	for {
		select {
		case stateEvent := <-stateChannel:
			switch stateEvent.Type {
			case statefeed.BlockProcessed:
				data := stateEvent.Data.(*statefeed.BlockProcessedData)
				headState, err := s.headFetcher.HeadState(s.ctx)
				if err != nil {
					log.WithError(err).Warn("failed to obtain head state")
					continue
				}

				go func() {
					// Update block statistics
					err = s.onBlock(headState, data.BlockRoot, s.headFetcher.HeadBlock())
					if err != nil {
						log.WithError(err).Warn("failed to update block stats")
					}

					// Update epoch statistics if applicable
					epoch := helpers.SlotToEpoch(headState.Slot)
					if s.currentEpoch == -1 {
						// First time round; set the epoch
						s.currentEpoch = int64(epoch)
					} else if epoch > uint64(s.currentEpoch) {
						// Change of epoch; wrap up stats for the previous epoch
						// Update epoch first so future blocks coming in don't trigger this again
						finishedEpoch := uint64(s.currentEpoch)
						s.currentEpoch = int64(epoch)
						err := s.onEpoch(headState, finishedEpoch)
						if err != nil {
							log.WithError(err).Warn("failed to update validator stats")
						}
					}
				}()
			}
		case <-s.ctx.Done():
			log.Debug("Context closed, exiting goroutine")
			return
		case err := <-stateSub.Err():
			log.WithError(err).Error("Subscription to state feed failed")
			return
		}
	}
}

// onEpoch is called whenever a block is received in a new epoch
func (s *Service) onEpoch(headState *pb.BeaconState, finishedEpoch uint64) error {
	// Reset the epoch stats in the service now so writes can continue
	s.mu.Lock()
	epochAttestations := s.epochAttestations
	s.epochAttestations = make(map[uint64]bool)
	s.mu.Unlock()

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	logValidatorStatsStmt, err := prepareInsertValidatorStmt(tx)
	if err != nil {
		return err
	}

	stats := &epochStats{
		epoch:                 finishedEpoch,
		lastJustifiedEpoch:    headState.GetPreviousJustifiedCheckpoint().GetEpoch(),
		currentJustifiedEpoch: headState.GetCurrentJustifiedCheckpoint().GetEpoch(),
		finalizedEpoch:        headState.GetFinalizedCheckpoint().GetEpoch(),
	}

	// Per-validator
	for i, validator := range headState.Validators {
		validatorStats := &validatorStats{
			epoch:            finishedEpoch,
			pubKey:           headState.Validators[i].PublicKey,
			balance:          headState.Balances[i],
			effectiveBalance: validator.EffectiveBalance,
		}
		if _, exists := epochAttestations[uint64(i)]; exists {
			validatorStats.attested = true
			// TODO add proposers as well to count of live instances (but don't double count)
			stats.liveInstances++
		}
		if validator.Slashed {
			if finishedEpoch < validator.ExitEpoch {
				validatorStats.state = "slashing"
				stats.slashingInstances++
				stats.slashingBalance += headState.Balances[i]
				stats.slashingEffectiveBalance += validator.EffectiveBalance
			} else {
				validatorStats.state = "slashed"
				stats.slashedInstances++
			}
		} else if validator.ExitEpoch != params.BeaconConfig().FarFutureEpoch {
			if finishedEpoch < validator.ExitEpoch {
				validatorStats.state = "exiting"
				stats.exitingInstances++
				stats.exitingBalance += headState.Balances[i]
				stats.exitingEffectiveBalance += validator.EffectiveBalance
			} else {
				validatorStats.state = "exited"
				stats.exitedInstances++
			}
		} else if finishedEpoch < validator.ActivationEpoch {
			// TODO also headState.GetEth1Data().GetDepositCount() - headState.Eth1DepositIndex ?
			// as per https://github.com/ethereum/eth2.0-metrics/blob/master/metrics.md#additional-metrics
			validatorStats.state = "pending"
			stats.pendingInstances++
			stats.pendingBalance += headState.Balances[i]
		} else {
			validatorStats.state = "active"
			stats.activeInstances++
			stats.activeBalance += headState.Balances[i]
			stats.activeEffectiveBalance += validator.EffectiveBalance
		}
		err = s.logValidatorStats(logValidatorStatsStmt, validatorStats)
		if err != nil {
			log.WithError(err).WithField("validator", i).Warn("Failed to log validator statistics")
		}
	}
	err = s.logEpochStats(tx, stats)
	if err != nil {
		log.WithError(err).Warn("Failed to log epoch statistics")
	}

	log.WithField("epoch", finishedEpoch).Debug("Updated metrics")

	return tx.Commit()
}

// onBlock is called whenever the head block changes
func (s *Service) onBlock(headState *pb.BeaconState, blockHash [32]byte, block *ethpb.BeaconBlock) error {

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	body := block.GetBody()

	s.mu.Lock()
	for i := range body.GetAttestations() {
		indices, err := helpers.AttestingIndices(headState, body.GetAttestations()[i].GetData(), body.GetAttestations()[i].GetAggregationBits())
		if err == nil {
			for _, j := range indices {
				s.epochAttestations[j] = true
			}
		}
	}
	s.mu.Unlock()

	stats := &blockStats{
		slot:              headState.Slot,
		hash:              blockHash,
		attestations:      len(body.GetAttestations()),
		deposits:          len(body.GetDeposits()),
		transfers:         0,
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

func (s *Service) updateNetworkState(headState *pb.BeaconState) error {
	if headState == nil {
		fmt.Printf("Headstate nil, cannot update network state\n")
		// Can happen when the chain is not yet up and running
		return nil
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	state := &networkState{
		timestamp:      time.Now().Unix(),
		epoch:          helpers.CurrentEpoch(headState),
		justifiedEpoch: headState.GetPreviousJustifiedCheckpoint().GetEpoch(),
		finalizedEpoch: headState.GetFinalizedCheckpoint().GetEpoch(),
	}

	err = s.logNetworkState(tx, state)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Service) poll(ctx context.Context) {
	for {
		headState, err := s.headFetcher.HeadState(s.ctx)
		if err == nil {
			err = s.updateNetworkState(headState)
			if err != nil {
				log.WithError(err).Warn("failed to update network state")
			}
		}

		// TODO make this configurable?
		time.Sleep(60 * time.Second)
	}
}
