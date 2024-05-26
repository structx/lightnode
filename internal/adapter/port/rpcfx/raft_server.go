package rpcfx

import (
	"context"

	"go.uber.org/zap"

	"github.com/structx/lightnode/internal/core/domain"
	pbv1 "github.com/structx/lightnode/proto/raft/v1"
)

// RaftServer
type RaftServer struct {
	pbv1.UnimplementedRaftServiceV1Server

	log  *zap.SugaredLogger
	raft domain.Raft
}

// NewRaftServer
func NewRaftServer(logger *zap.Logger) *RaftServer {
	return &RaftServer{
		log: logger.Sugar().Named("RaftServer"),
	}
}

// AppendEntries
func (rs *RaftServer) AppendEntries(ctx context.Context, in *pbv1.AppendEntryRequest) (*pbv1.AppendEntryResponse, error) {

	rs.log.Debugw("AppendEntries", "request", in)

	if in.Term < rs.raft.GetCurrentTerm() {
		return &pbv1.AppendEntryResponse{
			Term:    rs.raft.GetCurrentTerm(),
			Success: false,
		}, nil
	}

	ls := rs.raft.GetLogs()
	if l := ls[in.PrevLogIndex]; l == nil {
		return &pbv1.AppendEntryResponse{
			Term:    rs.raft.GetCurrentTerm(),
			Success: false,
		}, nil
	}

	if in.LeaderCommit > rs.raft.GetCommitIndex() {
		commitIndex := in.LeaderCommit
		if l := in.Entries[len(in.Entries)].Index; l < commitIndex {
			commitIndex = l
		}
		rs.raft.SetCommitIndex(commitIndex)
	}

	return &pbv1.AppendEntryResponse{
		Success: true,
		Term:    rs.raft.GetCurrentTerm(),
	}, nil
}

// InstallSnapshot
func (rs *RaftServer) InstallSnapshot(ctx context.Context, in *pbv1.InstallSnapshotRequest) (*pbv1.InstallSnapshotResponse, error) {
	rs.log.Debugw("InstallSnapshot", "request", in)
	return &pbv1.InstallSnapshotResponse{}, nil
}

// RequestVote
func (rs *RaftServer) RequestVote(ctx context.Context, in *pbv1.RequestVoteRequest) (*pbv1.RequestVoteResponse, error) {

	rs.log.Debugw("RequestVote", "request", in)

	if in.Term < rs.raft.GetCurrentTerm() {
		return &pbv1.RequestVoteResponse{
			Term:        rs.raft.GetCurrentTerm(),
			VoteGranted: false,
		}, nil
	}

	// TODO
	// implement check if candidates log is at least up to data as receivers log

	return &pbv1.RequestVoteResponse{
		Term:        rs.raft.GetCurrentTerm(),
		VoteGranted: true,
	}, nil
}
