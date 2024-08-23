package messaging

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

type EventHandler func(ctx context.Context, msg *pubsub.Message) error

type Dispatcher struct {
	Handlers map[string]EventHandler
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		Handlers: make(map[string]EventHandler),
	}
}

func (d *Dispatcher) RegisterHandler(pattern string, handler EventHandler) {
	d.Handlers[pattern] = handler
}

func (d *Dispatcher) Dispatch(ctx context.Context, msg *pubsub.Message) error {
	pattern, ok := msg.Attributes["pattern"]
	if !ok {
		return fmt.Errorf("pattern attribute not found in message")
	}

	handler, ok := d.Handlers[pattern]
	if !ok {
		return fmt.Errorf("no handler found for pattern: %s", pattern)
	}

	return handler(ctx, msg)
}
