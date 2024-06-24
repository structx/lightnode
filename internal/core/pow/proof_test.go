package pow_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/structx/lightnode/internal/core/domain"
	"github.com/structx/lightnode/internal/core/pow"
)

func TestGenerateHash(t *testing.T) {
	t.Run("default", func(t *testing.T) {

		b := &domain.Block{
			Timestamp:     time.Now().UTC().Format(time.RFC3339Nano),
			Difficulty:    3,
			Height:        1,
			Transactions:  []*domain.Transaction{},
			AccessCtrlRef: "*",
			AccessHash:    "",
		}
		pow.GenerateHash(b)
		fmt.Println(b.Hash) // 000000292e1c33aac0b37180d9cd7862f8c1241524d4fea482e9a039
	})
}
