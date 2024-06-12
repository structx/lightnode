package chain

import (
	"bytes"
	"encoding/json"
	"fmt"

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
		lastHash:     []byte(sc.latestHash),
	}
}

// Next
func (si *SimpleIterator) Next() (*domain.Block, error) {

	if bytes.Equal(si.lastHash, []byte{}) {
		si.lastHash = []byte{}
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

	if bytes.Equal(b.PrevHash, []byte{}) {
		si.lastHash = []byte{}
	}

	return &b, nil
}
