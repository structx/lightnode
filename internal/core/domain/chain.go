// Package domain application layer models and interfaces
package domain

// ChainState current state of chain
type ChainState int

const (
	// Initializing chain is starting up
	Initializing ChainState = iota
)

// SigWallet ...
type SigWallet struct{}

// Wallet ...
type Wallet struct{}

// Chain ...
type Chain interface {
	// AddBlock to chain
	AddBlock(Block) error
	// GetLatestBlock getter latest block
	GetLatestBlock() Block
	// GetBlockByHash retrieve block by hash
	GetBlockByHash(string) (*Block, error)
	// GetBlockHeight getter latest block height
	GetBlockHeight() int
	// ValidateBlock verify block is valid
	ValidateBlock(Block) error
	// ExecuteTransaction ...
	ExecuteTransaction(tx Transaction) error
	// GetPendingTransactions getter list of pending transactions
	GetPendingTransactions() []Transaction
	// AddTransaction add transaction to block
	AddTransaction(tx Transaction) error
	// GetState getter current chain state
	GetState() ChainState
	// NewSimpleIterator
	NewSimpleIterator() Iterator
}

// Iterator
type Iterator interface {
	Next() (*Block, error)
}

// PartialBlock
type PartialBlock struct {
	Hash      []byte
	PrevHash  []byte
	Timestamp string
	Height    int
}

// Block model
type Block struct {
	Hash          []byte        `json:"hash"`
	PrevHash      []byte        `json:"prev_hash"`
	Timestamp     string        `json:"timestamp"`
	Data          []byte        `json:"data"`
	Height        int           `json:"height"`
	Transactions  []Transaction `json:"transactions"`
	AccessCtrlRef string        `json:"access_ctrl_ref"`
	AccessHash    string        `json:"access_hash"`
}

// PartialTransaction
type PartialTransaction struct {
	ID        []byte
	Type      string
	Sender    string
	Receiver  string
	Timestamp string
}

// Transaction model
type Transaction struct {
	ID            []byte   `json:"id"`
	Type          string   `json:"type"`
	Sender        string   `json:"sender"`
	Receiver      string   `json:"receiver"`
	Data          []byte   `json:"data"`
	Timestamp     string   `json:"timestamp"`
	Signatures    []string `json:"signatures"`
	AccessCtrlRef string   `json:"access_ctrl_ref"`
}
