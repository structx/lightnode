// Package main entrypoint of application
package main

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/structx/lightnode/internal/adapter/logging"
	"github.com/structx/lightnode/internal/adapter/port/http/routerfx"
	"github.com/structx/lightnode/internal/adapter/port/http/serverfx"
	"github.com/structx/lightnode/internal/adapter/storage/store"
	"github.com/structx/lightnode/internal/core/chain"
	"github.com/structx/lightnode/internal/core/domain"
	"github.com/structx/lightnode/internal/core/service"
	"github.com/structx/lightnode/internal/core/setup"
)

func main() {
	fx.New(
		fx.Provide(context.TODO),
		fx.Provide(setup.NewConfig),
		fx.Invoke(setup.ParseConfigFromEnv),
		fx.Provide(logging.New),
		fx.Provide(fx.Annotate(store.NewStateMachine, fx.As(new(domain.StateMachine)))),
		fx.Provide(fx.Annotate(chain.New, fx.As(new(domain.Chain)))),
		fx.Provide(fx.Annotate(service.NewSimpleService, fx.As(new(domain.SimpleService)))),
		fx.Provide(fx.Annotate(routerfx.New, fx.As(new(http.Handler)))),
		fx.Provide(serverfx.NewHttp1Server),
		fx.Invoke(registerHooks),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	).Run()
}

func registerHooks(lc fx.Lifecycle, server *http.Server) error {
	lc.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {

				var result error

				go func() {
					if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						result = multierr.Append(result, fmt.Errorf("failed to start http server %v", err))
					}
				}()

				return result
			},
			OnStop: func(ctx context.Context) error {

				var result error

				err := server.Shutdown(ctx)
				if err != nil {
					result = multierr.Append(result, fmt.Errorf("failed to shutdown http server %v", err))
				}

				return result
			},
		},
	)
	return nil
}
