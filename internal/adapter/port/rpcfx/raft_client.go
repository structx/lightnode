package rpcfx

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/structx/lightnode/proto/raft/v1"
)

// RaftClientConfig
type RaftClientConfig struct {
	//
	ClientAddr string
	//
	Term         int64
	LeaderId     string
	PrevLogIndex int64
	PrevLogTerm  int64
	LeaderCommit int64
}

// RaftClient
type RaftClient struct {
	cfg  *RaftClientConfig
	conn *grpc.ClientConn
}

// NewRaftClient
func NewRaftClientWithConfig(ctx context.Context, cfg *RaftClientConfig) (*RaftClient, error) {

	timeout, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	conn, err := grpc.DialContext(
		timeout,
		cfg.ClientAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.DialContext: %v", err)
	}

	return &RaftClient{
		conn: conn,
		cfg:  cfg,
	}, nil
}

// SendHeartbeat
func (rc *RaftClient) SendHeartbeat(ctx context.Context) (int64, error) {
	rc.conn.Connect()

	cli := pb.NewRaftServiceV1Client(rc.conn)

	timeout, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	response, err := cli.AppendEntries(timeout, &pb.AppendEntryRequest{
		Term:         rc.cfg.Term,
		LeaderId:     rc.cfg.LeaderId,
		PrevLogIndex: rc.cfg.PrevLogIndex,
		PrevLogTerm:  rc.cfg.PrevLogTerm,
		LeaderCommit: rc.cfg.LeaderCommit,
		Entries:      make([]*pb.Log, 0),
	})
	if err != nil {
		return 0, fmt.Errorf("failed to send append entries request %v", err)
	}

	return response.Term, nil
}
