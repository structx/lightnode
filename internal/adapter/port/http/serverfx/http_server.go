package serverfx

import (
	"context"
	"net/http"

	"github.com/gorilla/handlers"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// InvokeHTTPServer
func InvokeHTTPServer(lc fx.Lifecycle, logger *zap.Logger, mux *http.ServeMux) error {
	s := &http.Server{
		Addr:    ":8080",
		Handler: handlers.CompressHandler(handlers.CORS()(mux)),
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("start http server")
				if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Fatal("unable to start http server", zap.Any("error", err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("shutdown http server")
			return s.Shutdown(ctx)
		},
	})
	return nil
}
