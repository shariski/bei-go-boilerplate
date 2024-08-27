package messaging

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

func HandleUserCreated(ctx context.Context, msg *pubsub.Message) error {
	// Process the UserCreated event
	log.Printf("Handled UserCreated event: %s", string(msg.Data))
	return nil
}
