package attestant

import (
	"database/sql"

	// Postgres
	_ "github.com/lib/pq"
)

// networkState is logged periodically
type networkState struct {
	timestamp          int64
	epoch              uint64
	justifiedEpoch     uint64
	lastJustifiedEpoch uint64
	finalizedEpoch     uint64
}

func (s *Service) logNetworkState(tx *sql.Tx, state *networkState) error {
	_, err := tx.Exec(`
	  INSERT INTO t_networkstate(
	    f_timestamp
	   ,f_epoch
	   ,f_justified_epoch
	   ,f_last_justified_epoch
	   ,f_finalized_epoch
	  ) VALUES (
	    $1
	   ,$2
	   ,$3
	   ,$4
	   ,$5
	  )`,
		state.timestamp,
		state.epoch,
		state.justifiedEpoch,
		state.lastJustifiedEpoch,
		state.finalizedEpoch)
	return err
}
