package logging

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/trevatk/k2/internal/adapter/setup"
)

// New
func New(config *setup.Config) (*zap.Logger, error) {

	level := strings.ToLower(config.Logger.Level)
	var l zapcore.Level

	switch level {
	case "info":
		l = zapcore.InfoLevel
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(l)
	cfg.OutputPaths = []string{}

	return cfg.Build()
}
