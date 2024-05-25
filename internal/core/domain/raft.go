package domain

// Log
type Log struct{}

// Raft service interface
type Raft interface {
	GetCurrentTerm() int64
	GetVotedFor() string
	GetLogs() []Log
}
