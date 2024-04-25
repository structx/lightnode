package main

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/multierr"

	"github.com/hashicorp/raft"

	"github.com/trevatk/olivia/internal/adapter/logging"
	"github.com/trevatk/olivia/internal/adapter/port/http/controller"
	"github.com/trevatk/olivia/internal/adapter/port/http/server"
	r "github.com/trevatk/olivia/internal/adapter/port/raft"
	"github.com/trevatk/olivia/internal/adapter/setup"
)

func main() {
	fx.New(
		fx.Provide(setup.New),
		fx.Invoke(setup.DecodeConfigFromEnv),
		fx.Provide(logging.New),
		fx.Provide(r.New),
		fx.Provide(server.New),
		fx.Invoke(controller.InvokeMetricsController),
		fx.Invoke(registerHooks),
	).Run()
}

func registerHooks(lc fx.Lifecycle, server *http.Server, raft *raft.Raft) error {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {

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

				fut := raft.Shutdown()
				err = fut.Error()
				if err != nil {
					result = multierr.Append(result, fmt.Errorf("failed to shutdown raft %v", err))
				}

				return result
			},
		},
	)
	return nil
}
