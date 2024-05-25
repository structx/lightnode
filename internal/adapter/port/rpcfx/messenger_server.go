// Package rpc gRPC interface implementation
package rpcfx

import (
	"context"
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pkgdomain "github.com/structx/go-dpkg/domain"
	pb "github.com/structx/go-dpkg/proto/messaging/v1"
	"github.com/structx/lightnode/internal/core/domain"
)

// GRPCServer protobuf implementation
type GRPCServer struct {
	pb.UnimplementedMessagingServiceV1Server

	log *zap.SugaredLogger
	ss  domain.SimpleService
}

// New grpc server constructor
func New(config pkgdomain.Config, logger *zap.Logger) *GRPCServer {
	return &GRPCServer{
		log: logger.Sugar().Named("GrpcServer"),
	}
}

// Publish not implemented
func (g *GRPCServer) Publish(context.Context, *pb.Envelope) (*pb.Stub, error) {
	// return empty responses
	return &pb.Stub{}, nil
}

// Subscribe not implemented
func (g *GRPCServer) Subscribe(*pb.Subscription, pb.MessagingServiceV1_SubscribeServer) error {
	// return empty responses
	return nil
}

// RequestResponse handler
func (g *GRPCServer) RequestResponse(ctx context.Context, in *pb.Envelope) (*pb.Envelope, error) {

	var (
		result interface{}
		err    error
	)

	switch domain.Topic(in.Topic) {
	case domain.SimpleChainQuery:
		result, err = g.ss.QueryBlockByHash(ctx, in.GetPayload())
	}

	if err != nil {
		g.log.Errorf("failed to run message operation %v", err)
		return nil, status.Errorf(codes.Internal, "unable to run message operation")
	}

	resultbytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result %v", err)
	}

	return &pb.Envelope{
		Topic:   in.Topic,
		Payload: resultbytes,
	}, nil
}
