package domain

import "context"

// Msg
type Msg struct {
	Topic   string
	Payload []byte
}

// MessageBroker
type MessageBroker interface {
	// Publish
	Publish(context.Context, string, []byte) error
	// Subscribe
	Subscribe(context.Context, string) (<-chan Msg, error)
	// Close
	Close() error
}

// Topics
type Topics []string
