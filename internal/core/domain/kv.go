package domain

// KV
type KV interface {
	Get(key []byte) ([]byte, error)
	Put(key, value []byte) error
}

// KvIterator
type KvIterator interface{}
