package messaging

import "bei-go-boilerplate/internal/api/messaging/dispatcher"

func RegisterHandlers(dispatcher *dispatcher.Dispatcher) {
	dispatcher.RegisterHandler("UserCreated", HandleUserCreated)
}
