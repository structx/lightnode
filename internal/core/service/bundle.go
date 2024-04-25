package service

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/trevatk/olivia/internal/core/domain"
	"go.uber.org/multierr"
)

var (
	topics domain.Topics = []string{
		"",
	}
)

// Bundle
type Bundle struct {
	c domain.Chain
	r domain.Raft
	m domain.MessageBroker
}

// Subscribe
func (b *Bundle) Subscribe(ctx context.Context) error {

	var result error
	cs := make([]<-chan domain.Msg, 0)

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
			b.subscriber(ctx, ch)
		}()
	}

	return result
}

func (b *Bundle) subscriber(ctx context.Context, ch <-chan domain.Msg) error {

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg, ok := <-ch:

			if !ok {
				return nil
			}

			switch msg.Topic {
			case "":
			default:
				return errors.New("invalid topic")
			}
		}
	}
}

func merge(ctx context.Context, cs ...<-chan domain.Msg) <-chan domain.Msg {

	out := make(chan domain.Msg)

	output := func(c <-chan domain.Msg) {
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
