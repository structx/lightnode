package chain

import (
	"fmt"
	"io"
	"sync"

	"github.com/hashicorp/raft"
	"github.com/trevatk/olivia/internal/core/domain"
)

type simpleChain struct {
	mtx sync.RWMutex
	kv  domain.KV
}

var _ raft.FSM = (*simpleChain)(nil)

// New
func New(kv domain.KV) *simpleChain {
	return &simpleChain{
		kv:  kv,
		mtx: sync.RWMutex{},
	}
}

// AddBlock
func (c *simpleChain) AddBlock(data []byte) error {
	return nil
}

// Apply
func (c *simpleChain) Apply(log *raft.Log) interface{} {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	err := c.AddBlock(log.Data)
	if err != nil {
		return fmt.Errorf("failed to add block %v", err)
	}

	return nil
}

// Snapshot
func (c *simpleChain) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}

// Restore
func (c *simpleChain) Restore(snapshot io.ReadCloser) error {
	return nil
}
