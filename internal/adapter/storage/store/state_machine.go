package store

import (
	"errors"
	"fmt"

	badger "github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/badger/v4/options"

	"github.com/structx/lightnode/internal/core/domain"
)

// StateMachine
type StateMachine struct {
	db *badger.DB
}

// interface compliance
var _ domain.StateMachine = (*StateMachine)(nil)

// NewStateMachine
func NewStateMachine() (*StateMachine, error) {

	opts := badger.Options{
		Dir:         "/opt/lightnode/data",
		Compression: options.Snappy,
	}
	db, err := badger.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to open badgerdb %v", err)
	}

	return &StateMachine{
		db: db,
	}, nil
}

// Set
func (sm *StateMachine) Put(key, value []byte) error {
	return sm.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
}

// Get
func (sm *StateMachine) Get(key []byte) ([]byte, error) {

	b := []byte{}
	if err := sm.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				return badger.ErrKeyNotFound
			}
		}

		_, err = item.ValueCopy(b)
		if err != nil {
			return fmt.Errorf("failed to copy value from db %v", err)
		}

		return nil
	}); err != nil {
		return []byte{}, err
	}

	return b, nil
}

// Close
func (sm *StateMachine) Close() error {
	return sm.db.Close()
}
