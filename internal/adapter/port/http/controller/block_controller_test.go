package controller_test

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/structx/go-dpkg/adapter/logging"
	"github.com/structx/go-dpkg/adapter/setup"
	"github.com/structx/go-dpkg/util/decode"

	"github.com/structx/lightnode/internal/adapter/port/http/controller"
	"github.com/structx/lightnode/internal/core/domain"
	"github.com/structx/lightnode/internal/core/domain/mocks"
)

func init() {
	_ = os.Setenv("DSERVICE_CONFIG", "testfiles/controller.test.hcl")
}

type BlockControllerSuite struct {
	suite.Suite
	blocks *controller.Blocks
}

func (suite *BlockControllerSuite) SetupTest() {

	assert := suite.Assert()

	cfg := setup.New()
	assert.NoError(decode.ConfigFromEnv(cfg))

	logger, err := logging.New(cfg)
	assert.NoError(err)

	mockService := mocks.NewSimpleService(suite.T())
	mockService.EXPECT().ReadBlockByHash(
		mock.Anything,
		mock.Anything,
	).Return(
		&domain.Block{},
		nil,
	).Maybe()

	mockService.EXPECT().PaginateBlocks(
		mock.Anything,
		mock.AnythingOfType("int64"),
		mock.AnythingOfType("int64"),
	).Return(
		[]*domain.PartialBlock{},
		nil,
	).Maybe()

	suite.blocks = controller.NewBlocks(logger, mockService)
}

func (suite *BlockControllerSuite) TestFetchByHash() {

	assert := suite.Assert()

	tt := []struct {
		expected int
		hash     []byte
	}{
		{
			// success
			expected: http.StatusAccepted,
			hash:     []byte("hello"),
		},
	}

	for _, t := range tt {

		rr := httptest.NewRecorder()
		r, e1 := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/blocks/%s", hex.EncodeToString(t.hash)), nil)
		assert.NoError(e1)

		suite.blocks.FetchByHash(rr, r)

		assert.Equal(t.expected, rr.Code)
	}
}

func (suite *BlockControllerSuite) TestNewPaginatePartialsResponse() {

	assert := suite.Assert()

	tt := []struct {
		expected      int
		limit, offset int
	}{
		{
			// success
			expected: http.StatusAccepted,
			limit:    10,
			offset:   0,
		},
	}

	for _, t := range tt {

		rr := httptest.NewRecorder()
		r, e1 := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/blocks?limit=%d&offset=%d", t.limit, t.offset), nil)
		assert.NoError(e1)

		suite.blocks.FetchByHash(rr, r)

		assert.Equal(t.expected, rr.Code)
	}
}

func TestBlockControllerSuite(t *testing.T) {
	suite.Run(t, new(BlockControllerSuite))
}
