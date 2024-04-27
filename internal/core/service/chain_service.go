package service

import (
	"encoding/json"
	"fmt"

	"github.com/trevatk/olivia/internal/core/domain"
	"github.com/trevatk/olivia/internal/core/topic"
)

// SimpleService simple chain service
type SimpleService struct {
	simpleChain domain.Chain
}

// Query operation against blockchain
func (ss *SimpleService) Query(payload []byte) (interface{}, error) {

	var msg topic.SimpleChainQueryMsg
	err := json.Unmarshal(payload, &msg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal message %v", err)
	}

	bl, err := ss.simpleChain.GetBlockByHash(msg.Hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get block by hash %v", err)
	}

	return bl, nil
}
