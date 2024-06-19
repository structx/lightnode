package chain

import (
	"errors"
	"fmt"
)

var (
	// ErrNoCandidates no candiate blocks found for mining
	ErrNoCandidates = errors.New("no canidate blocks found")
)

// ErrResourceNotFound
type ErrResourceNotFound struct {
	Hash string
}

// Error
func (e *ErrResourceNotFound) Error() string {
	return fmt.Sprintf("%s not found", e.Hash)
}
