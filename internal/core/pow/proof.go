package pow

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/sha3"

	"github.com/structx/lightnode/internal/core/domain"
)

// GenerateHash
func GenerateHash(b *domain.Block) {

	in := fmt.Sprintf("%x-%s-%d", b.PrevHash, b.Timestamp, b.Difficulty)

	for i := 0; ; i++ {
		nounce := fmt.Sprintf("%x", i)
		hash := computeHash(nounce, in)
		if isValidHash(b.Difficulty, hash) {
			b.Hash = hash
			return
		}
	}
}

func isValidHash(difficulty int, hash []byte) bool {
	prefix := bytes.Repeat([]byte{0}, difficulty)
	return bytes.HasPrefix(hash, prefix)
}

func computeHash(nounce, data string) []byte {

	h := sha3.New224()
	h.Write([]byte(data + nounce))

	return h.Sum(nil)
}
