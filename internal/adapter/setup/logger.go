package setup

// Logger
type Logger struct {
	Level   string `hcl:"level"`
	LogPath string `hcl:"log_path"`
}
