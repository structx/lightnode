package store_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/structx/lightnode/internal/adapter/storage/store"
	"github.com/structx/lightnode/internal/core/domain"
	"github.com/structx/lightnode/internal/core/setup"
)

type LocalStateSuite struct {
	suite.Suite
	stateMachine domain.StateMachine
}

func (suite *LocalStateSuite) SetupTest() {

	assert := suite.Assert()
	ctx := context.TODO()

	cfg := &setup.Config{Chain: &setup.Chain{BaseDir: "./testfiles/store"}, Logger: &setup.Logger{Level: "DEBUG", Path: "xyz"}}
	assert.NoError(setup.ParseConfigFromEnv(ctx, cfg))

	var err error
	suite.stateMachine, err = store.NewLocalStore(cfg)
	assert.NoError(err)
}

func (suite *LocalStateSuite) TestPut() {

	assert := suite.Assert()

	err := suite.stateMachine.Put([]byte("1"), []byte("1"))
	assert.NoError(err)

	err = suite.stateMachine.Put([]byte("2"), []byte("2"))
	assert.NoError(err)
}

func (suite *LocalStateSuite) TestGet() {

	assert := suite.Assert()

	// err := suite.stateMachine.Put([]byte("hello"), []byte("world"))
	// assert.NoError(err)

	_, err := suite.stateMachine.Get([]byte("hello"))
	assert.NoError(err)
	defer suite.stateMachine.Close()

}

func TestLocalStateSuite(t *testing.T) {
	suite.Run(t, new(LocalStateSuite))
}
