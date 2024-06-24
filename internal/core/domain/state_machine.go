package domain

// StateMachine
//
//go:generate mockery --name StateMachine
type StateMachine interface {
	Get(key string) ([]byte, error)
	Put(key string, value []byte) error
	Close() error
}
