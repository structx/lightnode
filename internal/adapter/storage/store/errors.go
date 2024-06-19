package store

import "fmt"

// ErrKeyNotFound
type ErrKeyNotFound struct {
	Hash string
}

// Error
func (enf *ErrKeyNotFound) Error() string {
	return fmt.Sprintf("key %s not found", enf.Hash)
}

// ErrKeyExists
type ErrKeyExists struct {
	Hash string
}

func (eke *ErrKeyExists) Error() string {
	return fmt.Sprintf("key %s already exists", eke.Hash)
}
