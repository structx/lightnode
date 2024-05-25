package logging

import "go.uber.org/zap"

// New
func New() (*zap.Logger, error) {
	return zap.NewDevelopment()
}
