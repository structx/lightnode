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

	"github.com/structx/lightnode/internal/core/domain"
	"github.com/structx/lightnode/internal/core/setup"
	pbv1 "github.com/structx/lightnode/proto/store/v1"
)

// LocalStore
type LocalStore struct {
	index atomic.Int64
	file  *os.File
	data  sync.Map
}

// interface compliance
var _ domain.StateMachine = (*LocalStore)(nil)

// NewLocalStore
func NewLocalStore(cfg *setup.Config) (*LocalStore, error) {

	path := filepath.Join(cfg.Chain.BaseDir, "chain_data")
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("unable to open %s %v", path, err)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var index int64 = 0
	ls := &LocalStore{
		data: sync.Map{},
		file: f,
	}

	for s.Scan() {

		err := s.Err()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("unable to scan line %v", err)
		}

		line := s.Text()
		if line == "" {
			break
		}

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

	ls.index = atomic.Int64{}
	ls.index.Store(index)

	return ls, nil
}

// Get
func (ls *LocalStore) Get(key string) ([]byte, error) {

	value, ok := ls.data.Load(key)
	if !ok || value == nil {
		return []byte{}, &ErrKeyNotFound{Hash: key}
	}

	decompressed, err := snappy.Decode(nil, value.([]byte))
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
func (ls *LocalStore) Put(key string, value []byte) error {

	if key == "" || bytes.Equal([]byte{}, value) {
		return errors.New("empty values provided")
	}

	recordbytes, err := proto.Marshal(&pbv1.Record{
		Key:   key,
		Value: value,
	})
	if err != nil {
		return fmt.Errorf("unable to marshal record %v", err)
	}

	compressed := snappy.Encode(nil, recordbytes)

	ls.file.Seek(0, io.SeekStart)
	_, err = ls.file.WriteString(hex.EncodeToString(compressed) + "\n")
	if err != nil {
		return fmt.Errorf("unable to write to file %v", err)
	}

	ls.data.Store(key, compressed)

	ls.index.Add(1)

	return nil
}

// Close
func (ls *LocalStore) Close() error {
	return ls.file.Close()
}
