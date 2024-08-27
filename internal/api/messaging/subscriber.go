package messaging

import (
	"bei-go-boilerplate/internal/api/messaging/dispatcher"
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func ConsumeSubscription(client *pubsub.Client, subID string, dispatcher *dispatcher.Dispatcher) error {
	ctx := context.Background()

	sub := client.Subscription(subID)

	err := sub.Receive(ctx, func(_ context.Context, m *pubsub.Message) {
		fmt.Printf("Got message: %q\n", string(m.Data))
		fmt.Println("Attributes:")
		for key, value := range m.Attributes {
			fmt.Printf("%s = %s\n", key, value)
			err := dispatcher.Dispatch(ctx, m)
			if err != nil {
				fmt.Printf("error dispatching message: %v\n", err)
				return
			}
		}
		m.Ack()
	})

	return err
}
