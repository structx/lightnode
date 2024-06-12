// Package chain blockchain implementation
package chain

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/structx/lightnode/internal/core/domain"
)

const (
	lastHash = "last_hash"
)

// SimpleChain chain implementation
type SimpleChain struct {
	mtx          sync.RWMutex
	stateMachine domain.StateMachine
	latestHash   []byte
}

// interface compliance
var _ domain.Chain = (*SimpleChain)(nil)

// New constructor
func New(stateMachine domain.StateMachine) (*SimpleChain, error) {

	gb := &domain.Block{
		Hash:          []byte("000000000000000000000000000"),
		Timestamp:     time.Now().Format(time.RFC3339Nano),
		Height:        1,
		Data:          []byte("genesis block"),
		PrevHash:      []byte{},
		Transactions:  []domain.Transaction{},
		AccessCtrlRef: "",
		AccessHash:    "",
	}

	genesisbytes, err := json.Marshal(gb)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal genesis block %v", err)
	}

	err = stateMachine.Put(gb.Hash, genesisbytes)
	if err != nil {
		return nil, fmt.Errorf("failed to put genesis block %v", err)
	}

	return &SimpleChain{
		latestHash:   gb.Hash,
		stateMachine: stateMachine,
		mtx:          sync.RWMutex{},
	}, nil
}

// AddBlock to chain
func (c *SimpleChain) AddBlock(block domain.Block) error {
	// TODO:
	// implement handler
	return nil
}

// GetBlockByHash ...
func (c *SimpleChain) GetBlockByHash(hash string) (*domain.Block, error) {

	blockbytes, err := c.stateMachine.Get([]byte(hash))
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil, ErrHashNotFound
		}

		return nil, fmt.Errorf("failed get block from store %v", err)
	}

	var b domain.Block
	err = json.Unmarshal(blockbytes, &b)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal block %v", err)
	}

	return &b, nil
}

// AddTransaction
func (c *SimpleChain) AddTransaction(tx domain.Transaction) error {

	blockbytes, err := c.stateMachine.Get([]byte(c.latestHash))
	if err != nil {
		return fmt.Errorf("unable to get block by latest hash %v", err)
	}

	var b domain.Block
	err = json.Unmarshal(blockbytes, &b)
	if err != nil {
		return fmt.Errorf("failed to unmarshal block %v", err)
	}

	b.Transactions = append(b.Transactions, tx)

	return nil
}

// Iter
func (c *SimpleChain) Iter() domain.Iterator {
	return &SimpleIterator{
		lastHash:     c.latestHash,
		stateMachine: c.stateMachine,
	}
}
