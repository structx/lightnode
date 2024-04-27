package chain

import "fmt"

// ErrInvalidBlock validate block error message
type ErrInvalidBlock struct {
	Field string
	Value any
}

// Error stringify error message
func (eib *ErrInvalidBlock) Error() string {
	return fmt.Sprintf("block field %s contains invalid value %v", eib.Field, eib.Value)
}

// ErrBlockMaxHeight block reached max height
type ErrBlockMaxHeight struct {
	CurrentHeight int
	MaxHeight     int
}

// Error stringify error message
func (erm *ErrBlockMaxHeight) Error() string {
	return fmt.Sprintf("block has reached %d and the max height is %d", erm.CurrentHeight, erm.MaxHeight)
}
