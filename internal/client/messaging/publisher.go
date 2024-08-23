package messaging

import (
	"bei-go-boilerplate/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
	"github.com/sirupsen/logrus"
)

type Publisher[T model.Event] struct {
	Publisher *pubsub.Client
	Topic     string
	Log       *logrus.Logger
}

func (p *Publisher[T]) GetTopic() string {
	return p.Topic
}

func (p *Publisher[T]) Publish(event []T, attributes map[string]string) error {
	ctx := context.Background()

	var wg sync.WaitGroup
	var totalErrors uint64
	t := p.Publisher.Topic(p.GetTopic())

	for _, msg := range event {
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			p.Log.WithError(err).Error("Failed to marshal message")
			atomic.AddUint64(&totalErrors, 1)
			continue
		}
		result := t.Publish(ctx, &pubsub.Message{
			Data: msgBytes,
			// TODO check if this parameter can accept empty map
			Attributes: attributes,
		})

		wg.Add(1)
		go func(msg T, res *pubsub.PublishResult) {
			defer wg.Done()

			id, err := res.Get(ctx)
			if err != nil {
				p.Log.WithError(err).Error("Failed to publish")
				atomic.AddUint64(&totalErrors, 1)
				return
			}

			p.Log.Info("Published message with msg ID: ", id, msg)
		}(msg, result)
	}

	wg.Wait()

	if totalErrors > 0 {
		err := fmt.Errorf("%d of %d messages did not publish successfully", totalErrors, len(event))
		p.Log.WithError(err)
		return err
	}
	return nil
}
