package rpc

import (
	"context"

	"github.com/hashicorp/raft"
	"github.com/trevatk/k2/internal/core/domain"
	pb "github.com/trevatk/k2/proto/k2/v1"
)

// GrpcServer
type GrpcServer struct {
	pb.UnimplementedK2ServiceServer
	chain domain.Chain
	r     *raft.Raft
}

func New(r *raft.Raft, chain domain.Chain) *GrpcServer {
	return &GrpcServer{
		chain: chain,
		r:     r,
	}
}

// SubmitBlock
func (g *GrpcServer) SubmitBlock(ctx context.Context, in *pb.NewBlock) (*pb.Block, error) {
	return nil, nil
}
