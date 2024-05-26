package raft

import "github.com/structx/lightnode/internal/core/domain"

// SimpleRaft
type SimpleRaft struct {
	// persistent state
	currentTerm int64        // latest term server has seen
	votedFor    string       // candidateId that received vote in current term
	logs        []domain.Log // log entries; each entry contains command for state machine
	// volatile state
	commitIndex int64 // index of highest log entry
	lastApplied int64 // index of highest log entry applied to state machine
	// volatile state on leaders
	nextIndex  []int64
	matchIndex []int64

	state domain.RaftStateEnum
}

// interface compliance
var _ domain.Raft = (*SimpleRaft)(nil)

// NewSimpleRaft
func NewSimpleRaft() *SimpleRaft {
	return &SimpleRaft{
		// persistent state
		currentTerm: 0,
		votedFor:    "",
		// volatile state
		commitIndex: 0,
		lastApplied: 0,
		// volatile state on leaders
		matchIndex: []int64{0},

		state: domain.Follower,
	}
}

// GetCurrentTerm
func (sr *SimpleRaft) GetCurrentTerm() int64 {
	return sr.currentTerm
}

// GetVotedFor
func (sr *SimpleRaft) GetVotedFor() string {
	return sr.votedFor
}

// GetLogs
func (sr *SimpleRaft) GetLogs() []domain.Log {
	return sr.logs
}

// AppendEntries
func (sr *SimpleRaft) AppendEntries() {}

// RequestVote
func (sr *SimpleRaft) RequestVote() {}
