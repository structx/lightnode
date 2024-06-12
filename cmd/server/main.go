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

	"github.com/structx/go-dpkg/adapter/port/http/serverfx"
	"github.com/structx/go-dpkg/adapter/port/raftfx"
	"github.com/structx/go-dpkg/adapter/setup"
	"github.com/structx/go-dpkg/adapter/storage/kv"
	pkgdomain "github.com/structx/go-dpkg/domain"
	"github.com/structx/go-dpkg/util/decode"

	"github.com/structx/lightnode/internal/adapter/logging"
	"github.com/structx/lightnode/internal/adapter/port/http/routerfx"
	"github.com/structx/lightnode/internal/core/chain"
	"github.com/structx/lightnode/internal/core/domain"
)

func main() {
	fx.New(
		fx.Provide(fx.Annotate(setup.New, fx.As(new(pkgdomain.Config)))),
		fx.Invoke(decode.ConfigFromEnv),
		fx.Provide(logging.New),
		fx.Provide(fx.Annotate(kv.NewPebble, fx.As(new(pkgdomain.KV)))),
		fx.Provide(fx.Annotate(chain.New, fx.As(new(domain.Chain)))),
		fx.Provide(fx.Annotate(routerfx.New, fx.As(new(http.Handler)))),
		fx.Provide(serverfx.New),
		fx.Provide(raftfx.New),
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
