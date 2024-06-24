// Package domain application layer models and interfaces
package domain

import (
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/sha3"
)

// ChainState current state of chain
type ChainState int

const (
	// Initializing chain is starting up
	Initializing ChainState = iota
	// Running chain is running
	Running
)

// SigWallet ...
type SigWallet struct{}

// Wallet ...
type Wallet struct{}

// Chain ...
//
//go:generate mockery --name Chain
type Chain interface {
	// AddBlock to chain
	AddBlock(Block) error
	// GetBlockByHash retrieve block by hash
	GetBlockByHash(string) (*Block, error)
	// AddTransaction add transaction to block
	AddTransaction(*Transaction) error
	// Iter
	Iter() Iterator
}

// Iterator
type Iterator interface {
	Next() (*Block, error)
}

// Block model
type Block struct {
	Hash          string         `json:"hash"`
	PrevHash      string         `json:"prev_hash"`
	Timestamp     string         `json:"timestamp"`
	Difficulty    int            `json:"difficulty"`
	Data          []byte         `json:"data"`
	Height        int            `json:"height"`
	Transactions  []*Transaction `json:"transactions"`
	AccessCtrlRef string         `json:"access_ctrl_ref"`
	AccessHash    string         `json:"access_hash"`
}

// PartialBlock
type PartialBlock struct {
	Hash      string
	PrevHash  string
	Timestamp string
	Height    int
}

// NewTransaction
type NewTransaction struct {
	Sender        string
	Receiver      string
	Data          []byte
	Signatures    []string
	AccessCtrlRef string
}

// Transaction model
type Transaction struct {
	ID            []byte   `json:"id"`
	Type          string   `json:"type"`
	Sender        string   `json:"sender"`
	Receiver      string   `json:"receiver"`
	Data          []byte   `json:"data"`
	Amount        int      `json:"amount"`
	Timestamp     string   `json:"timestamp"`
	Signatures    []string `json:"signatures"`
	AccessCtrlRef string   `json:"access_ctrl_ref"`
}

func (tx *Transaction) SetID() error {

	txbytes, err := json.Marshal(tx)
	if err != nil {
		return fmt.Errorf("unable to marshal tranasction %v", err)
	}

	h := sha3.New224()
	h.Write(txbytes)

	tx.ID = h.Sum(nil)

	return nil
}

// PartialTransaction
type PartialTransaction struct {
	ID        []byte
	Type      string
	Sender    string
	Receiver  string
	Timestamp string
}
