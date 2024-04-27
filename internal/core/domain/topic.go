package domain

// Topic lightnode topic list
type Topic string

const (
	// SimpleChainQuery query by block hash
	SimpleChainQuery Topic = "simple_chain_query"
)

// String cast enum to string
func (t Topic) String() string {
	return string(t)
}

// List all topics
func (t Topic) List() []Topic {
	return []Topic{SimpleChainQuery}
}
