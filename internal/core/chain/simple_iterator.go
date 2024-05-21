package chain

import "github.com/structx/go-pkg/domain"

// SimpleIterator
type SimpleIterator struct {
	store domain.KvIterator
}

// NewSimpleIterator
func NewSimpleIterator() SimpleIterator {
	return SimpleIterator{}
}
