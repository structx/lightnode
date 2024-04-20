package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"moul.io/chizap"

	"github.com/trevatk/k2/internal/adapter/port/http/controller"
)

// New
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
		controller.NewBundle(logger),
	}

	v1 := chi.NewRouter()

	for _, c := range cc {

		if r0, ok := c.(controller.V0); ok {
			h := r0.RegisterRoutesV0()
			r.Mount("/", h)
		}

		if r1, ok := c.(controller.V1); ok {
			h := r1.RegisterRoutesV1()
			v1.Mount("/", h)
		}
	}

	r.Mount("/api/v1", v1)

	return r
}
