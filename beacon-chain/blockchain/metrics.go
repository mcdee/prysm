package blockchain

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prysmaticlabs/prysm/shared/bytesutil"
	"github.com/prysmaticlabs/prysm/shared/params"
)

var (
	beaconSlot = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "beacon_slot",
		Help: "Latest slot of the beacon chain state",
	})
	beaconHeadSlot = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "beacon_head_slot",
		Help: "Slot of the head block of the beacon chain",
	})
	beaconHeadRoot = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "beacon_head_root",
		Help: "Root of the head block of the beacon chain, it returns the lowest 8 bytes interpreted as little endian",
	})
	competingAtts = promauto.NewCounter(prometheus.CounterOpts{
		Name: "competing_attestations",
		Help: "The # of attestations received and processed from a competing chain",
	})
	competingBlks = promauto.NewCounter(prometheus.CounterOpts{
		Name: "competing_blocks",
		Help: "The # of blocks received and processed from a competing chain",
	})
	processedBlkNoPubsub = promauto.NewCounter(prometheus.CounterOpts{
		Name: "processed_no_pubsub_block_counter",
		Help: "The # of processed block without pubsub, this usually means the blocks from sync",
	})
	processedBlkNoPubsubForkchoice = promauto.NewCounter(prometheus.CounterOpts{
		Name: "processed_no_pubsub_forkchoice_block_counter",
		Help: "The # of processed block without pubsub and forkchoice, this means indicate blocks from initial sync",
	})
	processedBlk = promauto.NewCounter(prometheus.CounterOpts{
		Name: "processed_block_counter",
		Help: "The # of total processed in block chain service, with fork choice and pubsub",
	})
	processedAttNoPubsub = promauto.NewCounter(prometheus.CounterOpts{
		Name: "processed_no_pubsub_attestation_counter",
		Help: "The # of processed attestation without pubsub, this usually means the attestations from sync",
	})
	processedAtt = promauto.NewCounter(prometheus.CounterOpts{
		Name: "processed_attestation_counter",
		Help: "The # of processed attestation with pubsub and fork choice, this ususally means attestations from rpc",
	})
)

func (s *Service) reportSlotMetrics(currentSlot uint64) {
	beaconSlot.Set(float64(currentSlot))
	beaconHeadSlot.Set(float64(s.HeadSlot()))
	beaconHeadRoot.Set(float64(bytesutil.ToLowInt64(s.HeadRoot())))
}

var (
	eth2ValidatorInstancesGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "eth2beacon_validator_instances_total",
		Help: "Number of state={pending|active|withdrawn|slashed} validators",
	}, []string{"state"})

	eth2ValidatorBalancesGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "eth2beacon_validator_balance_gwei_total",
		Help: "GWei in total held by active validators",
	}, []string{"state"})

	eth2ValidatorEffectiveBalancesGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "eth2beacon_validator_effectivebalance_gwei_total",
		Help: "GWei in total used for staking by active validators",
	}, []string{"state"})
)

func (s *Service) reportEpochMetrics(currentSlot uint64) {
	var pending uint64
	var pendingBalances uint64
	var active uint64
	var activeBalances uint64
	var activeEffectiveBalances uint64
	var slashing uint64
	var slashingBalances uint64
	var slashingEffectiveBalances uint64
	var slashed uint64
	var exiting uint64
	var exitingBalances uint64
	var exitingEffectiveBalances uint64
	var exited uint64

	currentEpoch := currentSlot / params.BeaconConfig().SlotsPerEpoch
	for i, validator := range s.headState.Validators {
		var state string
		if validator.Slashed {
			if currentEpoch < validator.ExitEpoch {
				state = "slashing"
				slashing++
				slashingBalances += s.headState.Balances[i]
				slashingEffectiveBalances += validator.EffectiveBalance
			} else {
				state = "slashed"
				slashed++
			}
		} else if validator.ExitEpoch != params.BeaconConfig().FarFutureEpoch {
			if currentEpoch < validator.ExitEpoch {
				state = "exiting"
				exiting++
				exitingBalances += s.headState.Balances[i]
				exitingEffectiveBalances += validator.EffectiveBalance
			} else {
				state = "exited"
				exited++
			}
		} else if currentEpoch < validator.ActivationEpoch {
			state = "pending"
			pending++
			pendingBalances += s.headState.Balances[i]
		} else {
			state = "active"
			active++
			activeBalances += s.headState.Balances[i]
			activeEffectiveBalances += validator.EffectiveBalance
		}
		fmt.Printf("{\"epoch\":%d,\"id\":%d,\"state\":%q,\"balance\":%d,\"effectivebalance\":%d}\n", currentEpoch, i, state, s.headState.Balances[i], validator.EffectiveBalance)
	}

	eth2ValidatorInstancesGauge.WithLabelValues("pending").Set(float64(pending))
	eth2ValidatorInstancesGauge.WithLabelValues("active").Set(float64(active))
	eth2ValidatorInstancesGauge.WithLabelValues("slashing").Set(float64(slashing))
	eth2ValidatorInstancesGauge.WithLabelValues("slashed").Set(float64(slashed))
	eth2ValidatorInstancesGauge.WithLabelValues("exiting").Set(float64(exiting))
	eth2ValidatorInstancesGauge.WithLabelValues("exited").Set(float64(exited))
	eth2ValidatorBalancesGauge.WithLabelValues("pending").Set(float64(pendingBalances))
	eth2ValidatorBalancesGauge.WithLabelValues("active").Set(float64(activeBalances))
	eth2ValidatorBalancesGauge.WithLabelValues("exiting").Set(float64(exitingBalances))
	eth2ValidatorBalancesGauge.WithLabelValues("slashing").Set(float64(slashingBalances))
	eth2ValidatorEffectiveBalancesGauge.WithLabelValues("active").Set(float64(activeEffectiveBalances))
	eth2ValidatorEffectiveBalancesGauge.WithLabelValues("exiting").Set(float64(exitingEffectiveBalances))
	eth2ValidatorEffectiveBalancesGauge.WithLabelValues("slashing").Set(float64(slashingEffectiveBalances))
}
