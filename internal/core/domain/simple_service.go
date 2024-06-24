package domain

import "context"

// SimpleService chain service interface
//
//go:generate mockery --name SimpleService
type SimpleService interface {
	// Query unmarshal msg and query block
	ReadBlockByHash(ctx context.Context, hash string) (*Block, error)
	// PaginateBlocks
	PaginateBlocks(ctx context.Context, limit, offset int64) ([]*PartialBlock, error)
	// ReadTxByHash
	ReadTxByHash(ctx context.Context, blockHash, txHash string) (*Transaction, error)
	// PaginateTransactions
	PaginateTransactions(ctx context.Context, hash string, limit, offset int64) ([]*PartialTransaction, error)
}
