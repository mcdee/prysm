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
	attested         bool
}

func (s *Service) logValidatorStats(tx *sql.Tx, stats *validatorStats) error {
	_, err := tx.Exec(`
	  INSERT INTO t_validatorstats(
	    f_epoch
	   ,f_id
	   ,f_state
	   ,f_balance
	   ,f_effective_balance
	   ,f_attested
	  ) VALUES (
	    $1
	   ,$2
	   ,$3
	   ,$4
	   ,$5
	   ,$6
	  )`,
		stats.epoch,
		stats.id,
		stats.state,
		stats.balance,
		stats.effectiveBalance,
		stats.attested)
	return err
}
