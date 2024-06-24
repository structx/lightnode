// Package routerfx chi router provider
package routerfx

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/structx/lightnode/internal/adapter/port/http/controller"
	"github.com/structx/lightnode/internal/core/domain"
)

// New constructor
func New(logger *zap.Logger, simpleService domain.SimpleService) *http.ServeMux {

	mux := http.NewServeMux()

	// service bundle endpoints
	controller.NewBundle(logger).RegisterRootRoutes(mux)

	// service endpoints
	controller.NewBlocks(logger, simpleService).RegisterRoutesV1(mux)
	controller.NewTransactions(logger, simpleService).RegisterRoutesV1(mux)

	return mux
}
