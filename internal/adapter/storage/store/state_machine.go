package store

import (
	"encoding/json"
	"fmt"

	"github.com/structx/lightnode/internal/core/domain"

	badger "github.com/dgraph-io/badger/v4"
)

// StateMachine
type StateMachine struct {
	db *badger.DB
}

// interface compliance
var _ domain.StateMachine = (*StateMachine)(nil)

// Store
func (sm *StateMachine) Store(log *domain.Log) error {
	return sm.db.Update(func(txn *badger.Txn) error {

		entrybytes, err := json.Marshal(log)
		if err != nil {
			return fmt.Errorf("failed to marshal log entry %v", err)
		}

		return txn.Set(entrybytes, log.Cmd)
	})
}
