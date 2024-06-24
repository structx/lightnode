package chain

import (
	"encoding/json"
	"fmt"

	"github.com/structx/lightnode/internal/core/domain"
)

// SimpleIterator
type SimpleIterator struct {
	lastHash     string
	stateMachine domain.StateMachine
}

// NewSimpleIterator
func (sc *SimpleChain) NewSimpleIterator() domain.Iterator {
	return &SimpleIterator{
		stateMachine: sc.stateMachine,
		lastHash:     sc.latestHash,
	}
}

// Next
func (si *SimpleIterator) Next() (*domain.Block, error) {

	if si.lastHash == "" {
		return nil, nil
	}

	blockbytes, err := si.stateMachine.Get(si.lastHash)
	if err != nil {
		return nil, fmt.Errorf("unable to get block from state machine %v", err)
	}

	fmt.Println(string(blockbytes))

	var b domain.Block
	err = json.Unmarshal(blockbytes, &b)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal block from state machine %v", err)
	}

	si.lastHash = b.PrevHash

	return &b, nil
}
