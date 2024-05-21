package domain

// SimpleService chain service interface
type SimpleService interface {
	// Query unmarshal msg and query block
	Query([]byte) (*Block, error)
}
