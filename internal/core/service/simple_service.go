package service

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/structx/lightnode/internal/core/chain"
	"github.com/structx/lightnode/internal/core/domain"
)

// SimpleService simple chain service
type SimpleService struct {
	ch domain.Chain
}

// interface compliance
var _ domain.SimpleService = (*SimpleService)(nil)

// NewSimpleService
func NewSimpleService(chain domain.Chain) *SimpleService {
	return &SimpleService{
		ch: chain,
	}
}

// Query operation against blockchain
func (ss *SimpleService) ReadBlockByHash(ctx context.Context, hash string) (*domain.Block, error) {

	select {
	case <-ctx.Done():
		return nil, nil
	default:

		decodedHash, err := hex.DecodeString(hash)
		if err != nil {
			return nil, fmt.Errorf("failed to decode hash %v", err)
		}

		block, err := ss.ch.GetBlockByHash(decodedHash)
		if err != nil {

			var notFound *chain.ErrResourceNotFound
			if errors.As(err, &notFound) {
				return nil, ErrNotFound
			}
			return nil, fmt.Errorf("failed to get block by hash %v", err)
		}

		return block, nil
	}
}

// PaginateBlocks
func (ss *SimpleService) PaginateBlocks(ctx context.Context, limit, offset int64) ([]*domain.PartialBlock, error) {

	select {
	case <-ctx.Done():
		return nil, nil
	default:

		it := ss.ch.Iter()
		blockSlice := make([]*domain.PartialBlock, 0, limit)

		var count int64 = 0

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

				blockSlice = append(blockSlice, &domain.PartialBlock{
					Hash:      block.Hash,
					PrevHash:  block.PrevHash,
					Timestamp: block.Timestamp,
					Height:    block.Height,
				})
				count++
			}
		}

		return blockSlice, nil
	}
}

// ReadTxByHash
func (ss *SimpleService) ReadTxByHash(ctx context.Context, blockHash, txHash []byte) (*domain.Transaction, error) {
	return nil, nil
}

// PaginateTransactions
func (ss *SimpleService) PaginateTransactions(ctx context.Context, hash string, limit, offset int64) ([]*domain.PartialTransaction, error) {
	return nil, nil
}
