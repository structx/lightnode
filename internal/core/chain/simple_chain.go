// Package chain blockchain implementation
package chain

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/hashicorp/raft"

	pkgdomain "github.com/trevatk/go-pkg/domain"
	"github.com/trevatk/olivia/internal/core/domain"
)

const (
	maxHeight = 15268
)

// SimpleChain chain implementation
type SimpleChain struct {
	mtx         sync.RWMutex
	latestBlock domain.Block
	state       domain.ChainState
	kv          pkgdomain.KV
}

// interface compliance
var _ domain.Chain = (*SimpleChain)(nil)
var _ raft.FSM = (*SimpleChain)(nil)

// New constructor
func New(kv pkgdomain.KV) *SimpleChain {
	return &SimpleChain{
		kv:          kv,
		state:       domain.Initializing,
		latestBlock: domain.Block{},
		mtx:         sync.RWMutex{},
	}
}

// AddBlock to chain
func (c *SimpleChain) AddBlock(block domain.Block) error {

	bb, err := json.Marshal(c.latestBlock)
	if err != nil {
		return fmt.Errorf("failed to marshal block %v", err)
	}

	err = c.kv.Put([]byte(c.latestBlock.Hash), bb)
	if err != nil {
		return fmt.Errorf("failed to put block into keyvalue store %v", err)
	}
	c.latestBlock = block

	return nil
}

// GetLatestBlock getter latest block
func (c *SimpleChain) GetLatestBlock() domain.Block {
	return c.latestBlock
}

// GetBlockByHash ...
func (c *SimpleChain) GetBlockByHash(_ string) (domain.Block, error) {
	// TODO:
	// implement function
	return c.latestBlock, nil
}

// GetBlockHeight getter latest block height
func (c *SimpleChain) GetBlockHeight() int {
	return c.latestBlock.Height
}

// ValidateBlock verify block is valid
func (c *SimpleChain) ValidateBlock(block domain.Block) error {

	when, err := time.Parse(time.RFC3339Nano, block.Timestamp)
	if err != nil {
		return fmt.Errorf("failed to validate timestamp %v", err)
	}

	if when.After(time.Now()) {
		return &ErrInvalidBlock{Field: "timestamp", Value: block.Timestamp}
	}

	return nil
}

// ExecuteTransaction add transaction to latest block
func (c *SimpleChain) ExecuteTransaction(tx domain.Transaction) error {

	// compare max height with current height
	if c.latestBlock.Height >= maxHeight {
		return &ErrBlockMaxHeight{CurrentHeight: c.latestBlock.Height, MaxHeight: maxHeight}
	}

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

// Apply ...
func (c *SimpleChain) Apply(_ *raft.Log) interface{} {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	// TODO:
	// implement handler
	return nil
}

// Snapshot ...
func (c *SimpleChain) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}

// Restore ...
func (c *SimpleChain) Restore(_ io.ReadCloser) error {
	return nil
}
