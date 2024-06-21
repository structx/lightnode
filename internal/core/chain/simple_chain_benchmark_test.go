package chain_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/structx/lightnode/internal/adapter/storage/store"
	"github.com/structx/lightnode/internal/core/chain"
	"github.com/structx/lightnode/internal/core/domain"
	"github.com/structx/lightnode/internal/core/setup"
)

var (
	simpleChain domain.Chain
)

func init() {

	cfg := &setup.Config{Chain: &setup.Chain{BaseDir: "./testfiles/simple_chain"}}

	stateMachine, err := store.NewLocalStore(cfg)
	if err != nil {
		fmt.Printf("unable to create local store %v", err)
		os.Exit(1)
	}

	sc, err := chain.New(stateMachine)
	if err != nil {
		fmt.Printf("unable to create new chain %v", err)
		os.Exit(1)
	}

	simpleChain = sc
}

func BenchmarkAddTransaction(b *testing.B) {

	for i := 0; i < b.N; i++ {
		simpleChain.AddTransaction(&domain.Transaction{
			ID:            []byte{},
			Type:          "coin",
			Sender:        "",
			Receiver:      "",
			Data:          []byte("coin transaction"),
			Amount:        i,
			Timestamp:     time.Now().UTC().Format(time.RFC3339Nano),
			Signatures:    []string{},
			AccessCtrlRef: "",
		})
	}
}

func BenchmarkAddBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleChain.AddBlock(domain.Block{
			Timestamp:     time.Now().String(),
			Difficulty:    0,
			Height:        1,
			Transactions:  []*domain.Transaction{},
			Data:          []byte(fmt.Sprintf("%x", i)),
			AccessCtrlRef: "*",
			AccessHash:    "",
		})
	}
}
