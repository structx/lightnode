package rpcfx

import "google.golang.org/grpc"

// RaftClient
type RaftClient struct {
	conn *grpc.ClientConn
}

// NewRaftClient
func NewRaftClient(addr string) (*RaftClient, error) {
	return &RaftClient{}, nil
}
