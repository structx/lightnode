// Package routerfx chi router provider
package routerfx

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"go.uber.org/zap"

	pkgcontroller "github.com/structx/go-dpkg/adapter/port/http/controller"
	"github.com/structx/lightnode/internal/adapter/port/http/controller"
	"github.com/structx/lightnode/internal/core/domain"
)

// New constructor
func New(logger *zap.Logger, simpleService domain.SimpleService) chi.Router {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Use(render.SetContentType(render.ContentTypeJSON))

	cc := []interface{}{
		pkgcontroller.NewBundle(logger),
		controller.NewBlocks(logger, simpleService),
		controller.NewTransactions(logger, simpleService),
	}

	v1 := chi.NewRouter()
	v1p := chi.NewRouter()

	for _, c := range cc {

		if v0, ok := c.(pkgcontroller.V0); ok {
			v0.RegisterRoutesV0(r)
		}

		if c1, ok := c.(pkgcontroller.V1); ok {
			c1.RegisterRoutesV1(v1)
		}

		if c1p, ok := c.(pkgcontroller.V1P); ok {
			c1p.RegisterRoutesV1P(v1p)
		}
	}

	r.Mount("/1.0", v1)

	return r
}
