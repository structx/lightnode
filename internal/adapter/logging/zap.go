package logging

import (
	"context"
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// New
func New(lc fx.Lifecycle) (*zap.Logger, error) {

	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("unable to create development logger %v", err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return logger.Sync()
		},
	})

	return logger, nil
}
