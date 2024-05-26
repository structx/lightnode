// Package chain blockchain implementation
package chain

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/cockroachdb/pebble"

	"github.com/structx/lightnode/internal/core/domain"
)

// SimpleChain chain implementation
type SimpleChain struct {
	mtx          sync.RWMutex
	latestBlock  domain.Block
	state        domain.ChainState
	stateMachine domain.StateMachine
}

// interface compliance
var _ domain.Chain = (*SimpleChain)(nil)

// New constructor
func New(stateMachine domain.StateMachine) *SimpleChain {
	return &SimpleChain{
		stateMachine: stateMachine,
		state:        domain.Initializing,
		latestBlock:  domain.Block{},
		mtx:          sync.RWMutex{},
	}
}

// AddBlock to chain
func (c *SimpleChain) AddBlock(block domain.Block) error {
	// TODO:
	// implement handler
	return nil
}

// GetLatestBlock getter latest block
func (c *SimpleChain) GetLatestBlock() domain.Block {
	return c.latestBlock
}

// GetBlockByHash ...
func (c *SimpleChain) GetBlockByHash(hash string) (*domain.Block, error) {

	blockbytes, err := c.stateMachine.Get([]byte(hash))
	if err != nil {
		if errors.Is(err, pebble.ErrNotFound) {
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

// GetBlockHeight getter latest block height
func (c *SimpleChain) GetBlockHeight() int {
	return c.latestBlock.Height
}

// ValidateBlock verify block is valid
func (c *SimpleChain) ValidateBlock(block domain.Block) error {
	// TODO:
	// implement function once block is defined
	return nil
}

// ExecuteTransaction add transaction to latest block
func (c *SimpleChain) ExecuteTransaction(tx domain.Transaction) error {
	// TODO:
	// implement max block height check
	c.latestBlock.Transactions = append(c.latestBlock.Transactions, tx)
	return nil
}

// GetPendingTransactions ...
func (c *SimpleChain) GetPendingTransactions() []domain.Transaction {
	// TODO:
	// implement function
	return []domain.Transaction{}
}

// AddTransaction ...
func (c *SimpleChain) AddTransaction(_ domain.Transaction) error {
	// TODO:
	// implement function
	return nil
}

// GetState getter chain state
func (c *SimpleChain) GetState() domain.ChainState {
	return c.state
}
