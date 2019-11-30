package attestant

import (
	"database/sql"
	"fmt"

	// Postgres
	_ "github.com/lib/pq"
)

// blockStats are logged for every block
type blockStats struct {
	slot              uint64
	hash              [32]byte
	deposits          int
	attestations      int
	transfers         int
	exits             int
	attesterSlashings int
	proposerSlashings int
}

func (s *Service) logBlockStats(tx *sql.Tx, stats *blockStats) error {
	_, err := tx.Exec(`
	  INSERT INTO t_blockstats(
	    f_slot
	   ,f_hash
	   ,f_deposits
	   ,f_attestations
	   ,f_transfers
	   ,f_exits
	   ,f_attester_slashings
	   ,f_proposer_slashings
	  ) VALUES (
	    $1
	   ,$2
	   ,$3
	   ,$4
	   ,$5
	   ,$6
	   ,$7
	   ,$8
	  )`,
		stats.slot,
		fmt.Sprintf("%x", stats.hash),
		stats.deposits,
		stats.attestations,
		stats.transfers,
		stats.exits,
		stats.attesterSlashings,
		stats.proposerSlashings)
	return err
}
