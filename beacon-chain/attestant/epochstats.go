package attestant

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// epochStats are logged for each epoch
type epochStats struct {
	epoch                    uint64
	pendingInstances         uint64
	pendingBalance           uint64
	activeInstances          uint64
	activeBalance            uint64
	activeEffectiveBalance   uint64
	slashingInstances        uint64
	slashingBalance          uint64
	slashingEffectiveBalance uint64
	exitingInstances         uint64
	exitingBalance           uint64
	exitingEffectiveBalance  uint64
	slashedInstances         uint64
	exitedInstances          uint64
}

func (s *Service) logEpochStats(tx *sql.Tx, stats *epochStats) error {
	_, err := tx.Exec(`
	  INSERT INTO t_epochstats(
	    f_epoch
	   ,f_pending_instances
	   ,f_pending_balance
	   ,f_active_instances
	   ,f_active_balance
	   ,f_active_effective_balance
	   ,f_slashing_instances
	   ,f_slashing_balance
	   ,f_slashing_effective_balance
	   ,f_exiting_instances
	   ,f_exiting_balance
	   ,f_exiting_effective_balance
	   ,f_slashed_instances
	   ,f_exited_instances
	  ) VALUES (
	    $1
	   ,$2
	   ,$3
	   ,$4
	   ,$5
	   ,$6
	   ,$7
	   ,$8
	   ,$9
	   ,$10
	   ,$11
	   ,$12
	   ,$13
	   ,$14
	  )`,
		stats.epoch,
		stats.pendingInstances,
		stats.pendingBalance,
		stats.activeInstances,
		stats.activeBalance,
		stats.activeEffectiveBalance,
		stats.slashingInstances,
		stats.slashingBalance,
		stats.slashingEffectiveBalance,
		stats.exitingInstances,
		stats.exitingBalance,
		stats.exitingEffectiveBalance,
		stats.slashedInstances,
		stats.exitedInstances)
	return err
}
