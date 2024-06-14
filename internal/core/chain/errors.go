package chain

import "fmt"

type ErrResourceNotFound struct {
	Hash string
}

func (e *ErrResourceNotFound) Error() string {
	return fmt.Sprintf("%s not found", e.Hash)
}
