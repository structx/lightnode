package main

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/multierr"

	"github.com/hashicorp/raft"

	"github.com/trevatk/go-pkg/adapter/logging"
	"github.com/trevatk/go-pkg/adapter/port/raftfx"
	"github.com/trevatk/go-pkg/adapter/setup"
	"github.com/trevatk/go-pkg/adapter/storage/kv"
	pkgdomain "github.com/trevatk/go-pkg/domain"
	"github.com/trevatk/go-pkg/util/decode"

	"github.com/trevatk/olivia/internal/adapter/port/http/router"
	"github.com/trevatk/olivia/internal/adapter/port/http/server"
	"github.com/trevatk/olivia/internal/core/chain"
	"github.com/trevatk/olivia/internal/core/domain"
)

func main() {
	fx.New(
		fx.Provide(fx.As(setup.New, fx.As(new(pkgdomain.Config)))),
		fx.Invoke(decode.ConfigFromEnv),
		fx.Provide(logging.New),
		fx.Provide(fx.Annotate(kv.NewPebble, fx.As(new(pkgdomain.KV)))),
		fx.Provide(fx.Annotate(chain.New, fx.As(new(domain.Chain), fx.As(new(raft.FSM))))),
		fx.Provide(fx.Annotate(router.New, fx.As(new(http.Handler)))),
		fx.Provide(raftfx.New),
		fx.Provide(server.New),
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
