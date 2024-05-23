package chain

import (
	"encoding/json"
	"fmt"

	dpkg "github.com/structx/go-dpkg/domain"
	"github.com/structx/lightnode/internal/core/domain"
)

// SimpleIterator
type SimpleIterator struct {
	lastHash string
	kv       dpkg.KV
}

// NewSimpleIterator
func (sc *SimpleChain) NewSimpleIterator() domain.Iterator {
	return &SimpleIterator{
		kv:       sc.kv,
		lastHash: sc.latestBlock.Hash,
	}
}

// Next
func (si *SimpleIterator) Next() (*domain.Block, error) {

	// if empty return empty
	if si.lastHash == "" {
		return nil, nil
	}

	blockbytes, err := si.kv.Get([]byte(si.lastHash))
	if err != nil {
		return nil, fmt.Errorf("failed to get hash %v", err)
	}

	var b domain.Block
	err = json.Unmarshal(blockbytes, &b)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal block from store %v", err)
	}

	si.lastHash = b.PrevHash

	return &b, nil
}
