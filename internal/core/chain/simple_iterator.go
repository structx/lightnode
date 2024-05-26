package chain

import (
	"github.com/structx/lightnode/internal/core/domain"
)

// SimpleIterator
type SimpleIterator struct {
	lastHash     []byte
	stateMachine domain.StateMachine
}

// NewSimpleIterator
func (sc *SimpleChain) NewSimpleIterator() domain.Iterator {
	return &SimpleIterator{
		stateMachine: sc.stateMachine,
		lastHash:     sc.latestBlock.Hash,
	}
}

// Next
func (si *SimpleIterator) Next() (*domain.Block, error) {
	// TODO:
	// implement iterator from badger
	return nil, nil
}
