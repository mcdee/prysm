package attestant

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// validatorStats are logged each epoch for each validator
type validatorStats struct {
	epoch            uint64
	id               int
	state            string
	balance          uint64
	effectiveBalance uint64
}

func (s *Service) logValidatorStats(tx *sql.Tx, stats *validatorStats) error {
	_, err := tx.Exec(`
	  INSERT INTO t_validatorstats(
	    f_epoch
	   ,f_id
	   ,f_state
	   ,f_balance
	   ,f_effective_balance
	  ) VALUES (
	    $1
	   ,$2
	   ,$3
	   ,$4
	   ,$5
	  )`,
		stats.epoch,
		stats.id,
		stats.state,
		stats.balance,
		stats.effectiveBalance)
	return err
}
