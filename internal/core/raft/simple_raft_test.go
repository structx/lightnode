package raft_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/structx/lightnode/internal/core/domain"
	"github.com/structx/lightnode/internal/core/raft"
)

type SimpleRaftSuite struct {
	suite.Suite
	raft domain.Raft
}

func (suite *SimpleRaftSuite) SetupTest() {
	suite.raft = raft.NewSimpleRaft()
}

func TestSimpleRaftSuite(t *testing.T) {
	suite.Run(t, new(SimpleRaftSuite))
}
