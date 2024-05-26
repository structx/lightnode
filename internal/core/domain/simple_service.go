package domain

import "context"

// SimpleService chain service interface
//
//go:generate mockery --name SimpleService
type SimpleService interface {
	// Query unmarshal msg and query block
	QueryBlockByHash(ctx context.Context, hash []byte) (*Block, error)
	// PaginateBlocks
	PaginateBlocks(ctx context.Context, limit, offset int64) ([]*PartialBlock, error)
	// ReadTxByHash
	ReadTxByHash(ctx context.Context, blockHash, txHash []byte) (*Transaction, error)
	// PaginateTransactions
	PaginateTransactions(ctx context.Context, blockHash []byte, limit, offset int64) ([]*PartialTransaction, error)
}
