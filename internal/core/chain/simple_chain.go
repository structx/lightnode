package chain

import (
	"fmt"
	"io"
	"sync"

	"github.com/hashicorp/raft"

	pkgdomain "github.com/trevatk/go-pkg/domain"
	"github.com/trevatk/olivia/internal/core/domain"
)

// SimpleChain
type SimpleChain struct {
	mtx sync.RWMutex

	kv pkgdomain.KV
}

// interface compliance
var _ domain.Chain = (*SimpleChain)(nil)
var _ raft.FSM = (*SimpleChain)(nil)

// New
func New(kv pkgdomain.KV) *SimpleChain {
	return &SimpleChain{
		kv:  kv,
		mtx: sync.RWMutex{},
	}
}

// AddBlock
func (c *SimpleChain) AddBlock(data []byte) error {
	return nil
}

// Apply
func (c *SimpleChain) Apply(log *raft.Log) interface{} {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	err := c.AddBlock(log.Data)
	if err != nil {
		return fmt.Errorf("failed to add block %v", err)
	}

	return nil
}

// Snapshot
func (c *SimpleChain) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}

// Restore
func (c *SimpleChain) Restore(snapshot io.ReadCloser) error {
	return nil
}
