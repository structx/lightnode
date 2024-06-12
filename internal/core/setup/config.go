package setup

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

// Config
type Config struct {
	Chain  *Chain
	Logger *Logger
}

// NewConfig
func NewConfig() *Config {
	return &Config{
		Chain:  &Chain{},
		Logger: &Logger{},
	}
}

// ParseConfigFromEnv
func ParseConfigFromEnv(ctx context.Context, cfg *Config) error {
	return envconfig.Process(ctx, cfg)
}
