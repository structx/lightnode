// Package router chi router provider
package routerfx

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"go.uber.org/zap"
	"moul.io/chizap"

	pkgcontroller "github.com/structx/go-dpkg/adapter/port/http/controller"
)

// New constructor
func New(logger *zap.Logger) chi.Router {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(chizap.New(
		logger, &chizap.Opts{
			WithReferer:   true,
			WithUserAgent: true,
		},
	))
	r.Use(middleware.Recoverer)

	cc := []interface{}{
		pkgcontroller.NewBundle(logger),
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

	r.Mount("/api/v1", v1)
	r.Mount("/api/v1/protected", v1p)

	return r
}
