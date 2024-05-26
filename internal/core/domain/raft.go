package domain

// RaftStateEnum
type RaftStateEnum int

const (
	// Follower
	Follower RaftStateEnum = iota
	// Candidate
	Candidate
	// Leader
	Leader
)

// Log
type Log struct {
	Index int64  `json:"index"`
	Term  int64  `json:"term"`
	Cmd   []byte `json:"cmd"`
}

// StateMachine
type StateMachine interface {
	Set(key, value []byte) error
	Get(key []byte) ([]byte, error)
}

// Raft service interface
type Raft interface {
	GetState() RaftStateEnum
	GetCurrentTerm() int64
	GetVotedFor() string
	GetLogs() []*Log
	GetCommitIndex() int64
	SetCommitIndex(index int64)
}
