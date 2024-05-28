package raft

import (
	"context"
	"fmt"
	"sync/atomic"

	"golang.org/x/sync/errgroup"

	"github.com/structx/lightnode/internal/adapter/port/rpcfx"
	"github.com/structx/lightnode/internal/core/domain"
)

// SimpleRaft
type SimpleRaft struct {
	// persistent state
	currentTerm int64         // latest term server has seen
	votedFor    string        // candidateId that received vote in current term
	logs        []*domain.Log // log entries; each entry contains command for state machine
	// volatile state
	commitIndex int64 // index of highest log entry
	lastApplied int64 // index of highest log entry applied to state machine
	// volatile state on leaders
	nextIndex  []int64
	matchIndex []int64

	state   domain.RaftStateEnum
	running atomic.Bool
}

// interface compliance
var _ domain.Raft = (*SimpleRaft)(nil)

// NewSimpleRaft
func NewSimpleRaft() *SimpleRaft {

	var raft SimpleRaft

	raft.running = atomic.Bool{}
	raft.running.Store(false)

	raft.currentTerm = 0
	raft.votedFor = ""

	raft.commitIndex = 0
	raft.lastApplied = 0

	raft.matchIndex = []int64{0}

	raft.state = domain.Follower

	return &raft
}

func (sr *SimpleRaft) start(ctx context.Context) error {

	sr.running.Store(true)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return sr.worker(ctx)
	})

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("failed to execute worker %v", err)
	}

	return nil
}

func (sr *SimpleRaft) worker(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:

			term, err := sr.sendHeartbeat(ctx, "")
			if err != nil {
				// handle gracefully
			}

			if term < sr.currentTerm {
				sr.state = domain.Follower
			}
		}
	}
}

func (sr *SimpleRaft) stop(ctx context.Context) {
	sr.running.Store(false)
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
func (sr *SimpleRaft) GetLogs() []*domain.Log {
	return sr.logs
}

// GetCommitIndex
func (sr *SimpleRaft) GetCommitIndex() int64 {
	return sr.commitIndex
}

// SetCommitIndex
func (sr *SimpleRaft) SetCommitIndex(index int64) {
	sr.commitIndex = index
}

// GetState
func (sr *SimpleRaft) GetState() domain.RaftStateEnum {
	return sr.state
}

func (sr *SimpleRaft) sendHeartbeat(ctx context.Context, addr string) (int64, error) {

	cfg := &rpcfx.RaftClientConfig{
		ClientAddr: addr,
		Term:       sr.currentTerm,
		LeaderId:   sr.votedFor,
	}
	client, err := rpcfx.NewRaftClientWithConfig(ctx, cfg)
	if err != nil {
		return 0, fmt.Errorf("unable to create raft client %v", err)
	}

	term, err := client.SendHeartbeat(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to send heartbeat %v", err)
	}

	return term, nil
}

func (sr *SimpleRaft) appendEntries() {}

func (sr *SimpleRaft) requestVote() {}
