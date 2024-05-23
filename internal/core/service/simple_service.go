package service

import (
	"context"
	"fmt"

	"github.com/structx/lightnode/internal/core/domain"
)

// SimpleService simple chain service
type SimpleService struct {
	ch domain.Chain
}

// Query operation against blockchain
func (ss *SimpleService) ReadBlockByHash(ctx context.Context, hash []byte) (*domain.Block, error) {

	select {
	case <-ctx.Done():
		return nil, nil
	default:

		block, err := ss.ch.GetBlockByHash(string(hash))
		if err != nil {
			return nil, fmt.Errorf("failed to get block by hash %v", err)
		}

		return block, nil
	}
}

// PaginateBlocks
func (ss *SimpleService) PaginateBlocks(ctx context.Context, limit, offset int) ([]*domain.Block, error) {

	select {
	case <-ctx.Done():
		return nil, nil
	default:

		it := ss.ch.NewSimpleIterator()
		blockSlice := make([]*domain.Block, 0, limit)

		count := 0

		for {

			block, err := it.Next()
			if err != nil {
				return nil, fmt.Errorf("iterator failed to get next block %v", err)
			}

			if block == nil {
				// end of chain
				break
			} else if offset == 0 {
				// check if limit reached
				if count == limit {
					break
				}

				blockSlice = append(blockSlice, block)
				count++
			}
		}

		return blockSlice, nil
	}
}
