// Package service application logic and implementation
package service

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"go.uber.org/multierr"

	pkgdomain "github.com/structx/go-dpkg/domain"
	"github.com/structx/lightnode/internal/core/domain"
)

var (
	topics = []string{"test"}
)

// Bundle application service bundle
type Bundle struct {
	c domain.Chain
	m pkgdomain.MessageBroker
}

// New constructor
func New(chain domain.Chain, messenger pkgdomain.MessageBroker) *Bundle {
	return &Bundle{
		c: chain,
		m: messenger,
	}
}

// Subscribe to all topics
func (b *Bundle) Subscribe(ctx context.Context) error {

	var result error
	cs := make([]<-chan pkgdomain.Envelope, 0)

	for _, t := range topics {
		ch, err := b.m.Subscribe(ctx, t)
		if err != nil {
			result = multierr.Append(result, fmt.Errorf("unable to subscribe to %s %v", t, err))
		}
		cs = append(cs, ch)
	}

	var wg sync.WaitGroup
	defer wg.Wait()

	ch := merge(ctx, cs...)

	for i := 0; i < len(topics); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = b.subscriber(ctx, ch)
		}()
	}

	return result
}

func (b *Bundle) subscriber(ctx context.Context, ch <-chan pkgdomain.Envelope) error {

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg, ok := <-ch:

			if !ok {
				return nil
			}

			switch msg.GetTopic() {
			case "":
			default:
				return errors.New("invalid topic")
			}
		}
	}
}

func merge(ctx context.Context, cs ...<-chan pkgdomain.Envelope) <-chan pkgdomain.Envelope {

	out := make(chan pkgdomain.Envelope)

	output := func(c <-chan pkgdomain.Envelope) {
		for n := range c {
			select {
			case <-ctx.Done():
				return
			case out <- n:

			}
		}
	}

	for _, c := range cs {
		go output(c)
	}

	go func() {
		<-ctx.Done()
		close(out)
	}()

	return out
}
