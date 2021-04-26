package nats

import "time"

type Message interface {
	Key() string
}

type NatsMessage struct {
	ID        string
	Subject   string
	Data      interface{}
	CreatedAt time.Time
}

func (msg *NatsMessage) Key() string {
	return "event.created"
}
