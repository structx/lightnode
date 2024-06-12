package setup

// Logger
type Logger struct {
	Level string `env:"LOG_LEVEL, required"`
	Path  string `env:"LOG_PATH, required"`
}
