package setup

import (
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

// Config
type Config struct {
	Logger        Logger        `hcl:"log,block"`
	Chain         Chain         `hcl:"chain,block"`
	Server        Server        `hcl:"server,block"`
	Raft          Raft          `hcl:"raft,block"`
	MessageBroker MessageBroker `hcl:"message_broker,block"`
}

// New
func New() *Config {
	return &Config{
		Logger:        Logger{},
		Chain:         Chain{},
		Server:        Server{},
		Raft:          Raft{},
		MessageBroker: MessageBroker{},
	}
}

// DecodeConfigFromEnv
func DecodeConfigFromEnv(cfg *Config) (*Config, error) {

	cfgPath := os.Getenv("DSERVICE_CONFIG")
	if cfgPath == "" {
		return nil, errors.New("$DSERVICE_CONFIG must be set")
	}

	err := hclsimple.DecodeFile(cfgPath, nil, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hcl file %v", err)
	}

	return cfg, nil
}
