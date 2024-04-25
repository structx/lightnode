package storage

import (
	"fmt"
	"path/filepath"

	"github.com/cockroachdb/pebble"
	"github.com/trevatk/olivia/internal/adapter/setup"
)

// PebbleDB
type PebbleDB struct {
	db *pebble.DB
}

// New
func New(cfg *setup.Config) (*PebbleDB, error) {

	path := filepath.Clean(cfg.Chain.BaseDir)
	db, err := pebble.Open(path, &pebble.Options{})
	if err != nil {
		return nil, fmt.Errorf("failed to open pebble database %v", err)
	}

	return &PebbleDB{
		db: db,
	}, nil
}
