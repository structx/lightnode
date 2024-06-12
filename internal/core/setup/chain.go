package setup

// Chain
type Chain struct {
	BaseDir string `env:"CHAIN_BASE_DIR, required"`
}
