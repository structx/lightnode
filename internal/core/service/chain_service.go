package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/structx/lightnode/internal/core/domain"
	"github.com/structx/lightnode/internal/core/topic"
)

// SimpleService simple chain service
type SimpleService struct {
	simpleChain domain.Chain
}

// Query operation against blockchain
func (ss *SimpleService) ReadBlockByHash(ctx context.Context, hash []byte) (interface{}, error) {
	select {
	case <-ctx.Done():
		return nil, nil
	default:
		return nil, nil
	}
	var msg topic.SimpleChainQueryMsg
	err := json.Unmarshal(hash, &msg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal message %v", err)
	}

	bl, err := ss.simpleChain.GetBlockByHash(msg.Hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get block by hash %v", err)
	}

	return bl, nil
}

// PaginateBlocks
func (ss *SimpleService) PaginateBlocks(ctx context.Context, limit, offset int) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		return nil
	}
}
