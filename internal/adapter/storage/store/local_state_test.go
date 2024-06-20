package store_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/structx/lightnode/internal/adapter/storage/store"
	"github.com/structx/lightnode/internal/core/setup"
)

func TestPut(t *testing.T) {
	t.Run("put", func(t *testing.T) {
		assert := assert.New(t)

		cfg := &setup.Config{Logger: &setup.Logger{Level: "DEBUG", Path: "xyz"}, Chain: &setup.Chain{BaseDir: "./testfiles/store"}}
		assert.NoError(setup.ParseConfigFromEnv(context.TODO(), cfg))

		stateMachine, err := store.NewLocalStore(cfg)
		assert.NoError(err)

		for i := 0; i < 100; i++ {
			err = stateMachine.Put([]byte(fmt.Sprintf("%d", i)), []byte(fmt.Sprintf("%d", i)))
			assert.NoError(err)
		}

		assert.NoError(stateMachine.Close())
	})
}
