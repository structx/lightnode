// Package main entrypoint of application
package main

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
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
		fx.Provide(fx.Annotate(store.NewLocalStore, fx.As(new(domain.StateMachine)))),
		fx.Provide(fx.Annotate(chain.New, fx.As(new(domain.Chain)))),
		fx.Provide(fx.Annotate(service.NewSimpleService, fx.As(new(domain.SimpleService)))),
		fx.Provide(routerfx.New),
		fx.Invoke(serverfx.InvokeHTTPServer, serverfx.InvokePprof),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	).Run()
}
