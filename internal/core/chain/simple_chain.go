// Package chain blockchain implementation
package chain

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sync/atomic"

	"sync"
	"time"

	"github.com/dgraph-io/badger/v4"

	"github.com/structx/lightnode/internal/adapter/storage/store"
	"github.com/structx/lightnode/internal/core/domain"
	"github.com/structx/lightnode/internal/core/pow"
)

const (
	lastHash  = "last_hash"
	maxHeight = 100
)

// SimpleChain chain implementation
type SimpleChain struct {
	stateMachine domain.StateMachine

	// volatile state
	latestHash []byte
	candidates sync.Map
	height     atomic.Int64
}

// interface compliance
var _ domain.Chain = (*SimpleChain)(nil)

// New constructor
func New(stateMachine domain.StateMachine) (*SimpleChain, error) {

	hash, err := stateMachine.Get([]byte(lastHash))
	if err != nil {

		var notFound *store.ErrKeyNotFound

		if errors.As(err, &notFound) {

			coinTx := &domain.Transaction{
				Sender:        "",
				Receiver:      "",
				Data:          []byte("genesis block coin transaction"),
				Timestamp:     time.Now().String(),
				Amount:        0,
				Signatures:    []string{},
				AccessCtrlRef: "*",
			}
			err = coinTx.SetID()
			if err != nil {
				return nil, fmt.Errorf("unable to set transaction id %v", err)
			}

			gb := &domain.Block{
				Timestamp:     time.Now().Format(time.RFC3339Nano),
				Height:        1,
				Data:          []byte("genesis block"),
				PrevHash:      []byte{},
				Transactions:  []*domain.Transaction{coinTx},
				AccessCtrlRef: "",
				AccessHash:    "",
			}
			pow.GenerateHash(gb)

			genesisbytes, err := json.Marshal(gb)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal genesis block %v", err)
			}

			err = stateMachine.Put(gb.Hash, genesisbytes)
			if err != nil {
				return nil, fmt.Errorf("failed to put genesis block %v", err)
			}

			err = stateMachine.Put([]byte(lastHash), gb.Hash)
			if err != nil {
				return nil, fmt.Errorf("unable to put last hash %v", err)
			}

			chain := &SimpleChain{}
			chain.height = atomic.Int64{}
			chain.height.Store(int64(gb.Height))

			chain.latestHash = gb.Hash
			chain.stateMachine = stateMachine

			chain.candidates = sync.Map{}

			return chain, nil
		}
		return nil, fmt.Errorf("failed to get last hash %v", err)
	}

	blockbytes, err := stateMachine.Get(hash)
	if err != nil {
		return nil, fmt.Errorf("unable to get last hash %v", err)
	}

	var b domain.Block
	err = json.Unmarshal(blockbytes, &b)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal block %v", err)
	}

	chain := &SimpleChain{}
	chain.height = atomic.Int64{}
	chain.height.Store(int64(b.Height))

	chain.latestHash = b.Hash
	chain.stateMachine = stateMachine

	chain.candidates = sync.Map{}

	return chain, nil
}

// AddBlock to chain
func (c *SimpleChain) AddBlock(block domain.Block) error {

	coinTx := &domain.Transaction{
		Data:          []byte("coin tx"),
		AccessCtrlRef: "*",
		Type:          "coin",
		Sender:        "",
		Receiver:      "",
		Amount:        0,
		Timestamp:     time.Now().UTC().Format(time.RFC3339Nano),
		Signatures:    []string{},
	}
	block.Transactions = append(block.Transactions, coinTx)
	block.PrevHash = c.latestHash

	pow.GenerateHash(&block)

	blockbytes, err := json.Marshal(&block)
	if err != nil {
		return fmt.Errorf("unable to marshal block %v", err)
	}

	err = c.stateMachine.Put(block.Hash, blockbytes)
	if err != nil {
		return fmt.Errorf("unable to put block %v", err)
	}
	c.candidates.Store(block.Hash, &block)

	return nil
}

// GetBlockByHash ...
func (c *SimpleChain) GetBlockByHash(hash []byte) (*domain.Block, error) {

	blockbytes, err := c.stateMachine.Get(hash)
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil, &ErrResourceNotFound{Hash: hex.EncodeToString(hash)}
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
func (c *SimpleChain) AddTransaction(tx *domain.Transaction) error {

	if bytes.Equal(tx.ID, []byte{}) {
		err := tx.SetID()
		if err != nil {
			return fmt.Errorf("failed to set tx ID %v", err)
		}
	}

	if c.height.Load() >= maxHeight {
		// block is full transaction must append to next candidate block
		c.candidates.Range(func(key, value any) bool {

			block, ok := value.(*domain.Block)
			if ok {
				// check if block from iter is correct
				if !bytes.Equal(c.latestHash, block.PrevHash) {
					// next
					return true
				}

				// update last hash
				err := c.stateMachine.Put([]byte(lastHash), block.Hash)
				if err != nil {
					return false
				}
				c.latestHash = block.Hash
			}

			// remove current block from candidates
			c.candidates.Delete(key)

			// stop iteration
			return true
		})

		// begin recursive call
		return c.AddTransaction(tx)
	}

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

	blockbytes, err = json.Marshal(b)
	if err != nil {
		return fmt.Errorf("failed to marshal block %v", err)
	}

	err = c.stateMachine.Put(b.Hash, blockbytes)
	if err != nil {
		return fmt.Errorf("unable to put block %v", err)
	}

	return nil
}

// Iter
func (c *SimpleChain) Iter() domain.Iterator {
	return &SimpleIterator{
		lastHash:     c.latestHash,
		stateMachine: c.stateMachine,
	}
}
