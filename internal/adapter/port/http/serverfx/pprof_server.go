package serverfx

import (
	"context"
	"net/http"
	_ "net/http/pprof"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// InvokePprof
func InvokePprof(lc fx.Lifecycle, logger *zap.Logger) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Sugar().Fatal(http.ListenAndServe("localhost:6060", nil))
			}()
			return nil
		},
	})
	return nil
}
