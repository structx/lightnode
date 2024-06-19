package store

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/structx/lightnode/internal/core/setup"
)

type record struct {
	Index int64  `json:"index"`
	Key   []byte `json:"key"`
	Value []byte `json:"value"`
}

// LocalStore
type LocalStore struct {
	index atomic.Int64
	file  *os.File
	data  sync.Map
}

// NewLocalStore
func NewLocalStore(cfg *setup.Config) (*LocalStore, error) {

	path := filepath.Join(cfg.Chain.BaseDir, "chain_data.json")
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {

		if errors.Is(err, os.ErrExist) {
			f, err = os.Open(path)
			if err != nil {
				return nil, fmt.Errorf("unable to open %s %v", path, err)
			}

			contents, err := io.ReadAll(f)
			if err != nil {
				return nil, fmt.Errorf("unable to read file contents %v", err)
			}

			var records []record
			err = json.Unmarshal(contents, &records)
			if err != nil {
				return nil, fmt.Errorf("unable to unmarshal records %v", err)
			}

			ls := &LocalStore{}
			ls.file = f
			ls.data = sync.Map{}
			ls.index = atomic.Int64{}
			ls.index.Store(int64(len(records) - 1))

			for _, r := range records {
				ls.data.Store(r.Key, r.Value)
			}

			return ls, nil
		}

		return nil, fmt.Errorf("unable to create %s %v", path, err)
	}

	ls := &LocalStore{}
	ls.file = f
	ls.data = sync.Map{}
	ls.index = atomic.Int64{}

	return &LocalStore{
		file: f,
		data: sync.Map{},
	}, nil
}

// Get
func (ls *LocalStore) Get(key []byte) ([]byte, error) {

	value, ok := ls.data.Load(hex.EncodeToString(key))
	if !ok {
		return []byte{}, &ErrKeyNotFound{Hash: hex.EncodeToString(key)}
	}

	return value.([]byte), nil
}

// Put
func (ls *LocalStore) Put(key, value []byte) error {

	_, ok := ls.data.Load(hex.EncodeToString(key))
	if ok {
		return &ErrKeyExists{Hash: hex.EncodeToString(key)}
	}

	recordbytes, err := json.Marshal(&record{
		Index: ls.index.Load(),
		Key:   key,
		Value: value,
	})
	if err != nil {
		return fmt.Errorf("unable to marshal record %v", err)
	}

	_, err = ls.file.Write(recordbytes)
	if err != nil {
		return fmt.Errorf("failed to write to file %v", err)
	}

	ls.data.Store(hex.EncodeToString(key), value)

	ls.index.Add(1)

	return nil
}

// Close
func (ls *LocalStore) Close() error {
	return ls.file.Close()
}
