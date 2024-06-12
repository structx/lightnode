package domain

// StateMachine
//
//go:generate mockery --name StateMachine
type StateMachine interface {
	Get(key []byte) ([]byte, error)
	Put(key, value []byte) error
	Close() error
}
