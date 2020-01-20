// Package operations provides in-memory pools for operations such as voluntary exits, that need to be kept in sync with their
// additions through RPC/network and removals through inclusion in blocks.
package operations

import (
	"context"

	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/go-ssz"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/feed"
	opfeed "github.com/prysmaticlabs/prysm/beacon-chain/core/feed/operation"
	statefeed "github.com/prysmaticlabs/prysm/beacon-chain/core/feed/state"
	"github.com/prysmaticlabs/prysm/beacon-chain/db"
	"github.com/prysmaticlabs/prysm/beacon-chain/operations/voluntaryexits"
	"github.com/sirupsen/logrus"
)

var log logrus.FieldLogger

func init() {
	log = logrus.WithField("prefix", "operations")
}

// Service defining an operations server for a beacon node.
type Service struct {
	ctx                context.Context
	cancel             context.CancelFunc
	db                 db.ReadOnlyDatabase
	voluntaryExitsPool voluntaryexits.Pool
	stateNotifier      statefeed.Notifier
	operationNotifier  opfeed.Notifier
}

// Config options for the beacon node operations server.
type Config struct {
	DB                db.ReadOnlyDatabase
	StateNotifier     statefeed.Notifier
	OperationNotifier opfeed.Notifier
}

// NewService instantiates a new operations service instance that will
// be registered into a running beacon node.
func NewService(ctx context.Context, cfg *Config) (*Service, error) {
	ctx, cancel := context.WithCancel(ctx)
	return &Service{
		ctx:                ctx,
		cancel:             cancel,
		db:                 cfg.DB,
		voluntaryExitsPool: voluntaryexits.NewPool(),
		stateNotifier:      cfg.StateNotifier,
		operationNotifier:  cfg.OperationNotifier,
	}, nil
}

// Start the operations service.
func (s *Service) Start() {
	// Listen for new voluntary exits received to add to pool.
	operationChannel := make(chan *feed.Event, 1)
	operationSub := s.operationNotifier.OperationFeed().Subscribe(operationChannel)
	defer operationSub.Unsubscribe()

	// Listen for new blocks processed to remove from pool.
	stateChannel := make(chan *feed.Event, 1)
	stateSub := s.stateNotifier.StateFeed().Subscribe(stateChannel)
	defer stateSub.Unsubscribe()

	for {
		select {
		case event := <-stateChannel:
			if event.Type == statefeed.BlockProcessed {
				data := event.Data.(*statefeed.BlockProcessedData)
				if block, err := s.db.Block(s.ctx, data.BlockRoot); err == nil {
					for _, exit := range block.Block.Body.VoluntaryExits {
						root, err := ssz.HashTreeRoot(exit)
						if err == nil && s.voluntaryExitsPool.HasVoluntaryExit(root) {
							log.WithField("index", exit.Exit.ValidatorIndex).Debug("Removing processed voluntary exit of validator from pool")
							s.voluntaryExitsPool.DeleteVoluntaryExit(exit)
						}
					}
				}
			}
		case event := <-operationChannel:
			if event.Type == opfeed.ExitReceived {
				data := event.Data.(*opfeed.ExitReceivedData)
				s.voluntaryExitsPool.SaveVoluntaryExit(data.Exit)
				// TODO error, logging?
			}
		case <-s.ctx.Done():
			log.Debug("Context closed, exiting goroutine")
			return
		case err := <-operationSub.Err():
			log.WithError(err).Error("Subscription to operation feed notifier failed")
			return
		case err := <-stateSub.Err():
			log.WithError(err).Error("Subscription to state feed notifier failed")
			return
		}
	}
}

// Stop the service.
func (s *Service) Stop() error {
	s.cancel()
	return nil
}

// Status returns nil.
func (s *Service) Status() error {
	return nil
}

// VoluntaryExits returns the list of voluntary exits that are currently in the pool.
func (s *Service) VoluntaryExits() []*ethpb.SignedVoluntaryExit {
	return s.voluntaryExitsPool.VoluntaryExits()
}
