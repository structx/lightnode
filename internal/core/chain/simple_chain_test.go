package chain_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/structx/lightnode/internal/adapter/storage/store"
	"github.com/structx/lightnode/internal/core/chain"
	"github.com/structx/lightnode/internal/core/domain"
	"github.com/structx/lightnode/internal/core/setup"
)

type SimpleChainSuite struct {
	suite.Suite
	chain domain.Chain
}

func (suite *SimpleChainSuite) SetupTest() {

	assert := suite.Assert()

	cfg := &setup.Config{Chain: &setup.Chain{BaseDir: "./testfiles/simple_chain"}, Logger: &setup.Logger{Level: "DEBUG", Path: "xyz"}}
	assert.NoError(setup.ParseConfigFromEnv(context.TODO(), cfg))

	stateMachine, err := store.NewLocalStore(cfg)
	assert.NoError(err)

	suite.chain, err = chain.New(stateMachine)
	assert.NoError(err)
}

func (suite *SimpleChainSuite) TestAddBlock() {

	assert := suite.Assert()

	err := suite.chain.AddBlock(domain.Block{
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		Height:    1,
		Data:      []byte("hello world"),
	})
	assert.NoError(err)
}

func TestSimpleChainSuite(t *testing.T) {
	suite.Run(t, new(SimpleChainSuite))
}
