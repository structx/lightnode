package store

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/golang/snappy"

	"google.golang.org/protobuf/proto"

	"github.com/structx/lightnode/internal/core/setup"
	pbv1 "github.com/structx/lightnode/proto/store/v1"
)

// LocalStore
type LocalStore struct {
	index atomic.Int64
	file  *os.File
	data  sync.Map
}

// NewLocalStore
func NewLocalStore(cfg *setup.Config) (*LocalStore, error) {

	path := filepath.Join(cfg.Chain.BaseDir, "chain_data")
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("unable to open %s %v", path, err)
	}

	s := bufio.NewScanner(f)

	var index int64 = 0
	ls := &LocalStore{}
	ls.data = sync.Map{}

	for s.Scan() {
		line := s.Text()
		buf, err := hex.DecodeString(line)
		if err != nil {
			return nil, fmt.Errorf("unable to decode line %v", err)
		}

		decompressed, err := snappy.Decode(nil, buf)
		if err != nil {
			return nil, fmt.Errorf("unable to decompress line %v", err)
		}

		var record pbv1.Record
		err = proto.Unmarshal(decompressed, &record)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal line %v", err)
		}

		ls.data.Store(record.Key, line)

		index += 1
	}

	ls.file = f
	ls.index = atomic.Int64{}
	ls.index.Store(index)

	return ls, nil
}

// Get
func (ls *LocalStore) Get(key []byte) ([]byte, error) {

	value, ok := ls.data.Load(hex.EncodeToString(key))
	if !ok {
		return []byte{}, &ErrKeyNotFound{Hash: hex.EncodeToString(key)}
	}

	decoded, err := hex.DecodeString(value.(string))
	if err != nil {
		return []byte{}, fmt.Errorf("unable to decode string %v", err)
	}

	decompressed, err := snappy.Decode(nil, decoded)
	if err != nil {
		return nil, fmt.Errorf("failed to decode message %v", err)
	}

	var record pbv1.Record
	err = proto.Unmarshal(decompressed, &record)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to unmarshal decode bytes %v", err)
	}

	return record.Value, nil
}

// Put
func (ls *LocalStore) Put(key, value []byte) error {

	if bytes.Equal([]byte{}, key) || bytes.Equal([]byte{}, value) {
		return errors.New("empty values provided")
	}

	_, ok := ls.data.Load(hex.EncodeToString(key))
	if ok {
		return &ErrKeyExists{Hash: hex.EncodeToString(key)}
	}

	recordbytes, err := proto.Marshal(&pbv1.Record{
		Key:   hex.EncodeToString(key),
		Value: value,
	})
	if err != nil {
		return fmt.Errorf("unable to marshal record %v", err)
	}

	compressed := snappy.Encode(nil, recordbytes)
	compressedStr := hex.EncodeToString(compressed)

	ls.file.Seek(0, io.SeekStart)
	_, err = ls.file.WriteString(compressedStr + "\n")
	if err != nil {
		return fmt.Errorf("unable to write to file %v", err)
	}

	ls.data.Store(hex.EncodeToString(key), compressedStr)

	ls.index.Add(1)

	return nil
}

// Close
func (ls *LocalStore) Close() error {
	return ls.file.Close()
}
