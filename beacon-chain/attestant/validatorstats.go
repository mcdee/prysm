package attestant

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// validatorStats are logged each epoch for each validator
type validatorStats struct {
	epoch            uint64
	pubKey           []byte
	state            string
	balance          uint64
	effectiveBalance uint64
	attested         bool
}

func prepareInsertValidatorStmt(tx *sql.Tx) (*sql.Stmt, error) {
	return tx.Prepare(`
	  INSERT INTO t_validatorstats(
	    f_epoch
	   ,f_public_key
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
	  )`)
}

func (s *Service) logValidatorStats(stmt *sql.Stmt, stats *validatorStats) error {
	_, err := stmt.Exec(stats.epoch, stats.pubKey, stats.state, stats.balance, stats.effectiveBalance, stats.attested)
	return err
}
