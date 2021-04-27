package nats

import (
	"context"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

func NewConnection() (nc *nats.Conn, err error) {
	// Connect Options.
	opts := []nats.Option{nats.Name("NATS")}

	// Connect to NATS
	nc, err = nats.Connect(nats.DefaultURL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	return
}

// nats publish
func Publish(nc *nats.Conn, topic string, message interface{}) (err error) {

	js, _ := json.Marshal(message)
	nc.Publish(topic, js)

	nc.FlushWithContext(context.Background())

	if err = nc.LastError(); err != nil {
		return
	}

	return nil
}

// nats subscription
func Subscription(nc *nats.Conn, topic string) (response interface{}, err error) {
	nc.Subscribe(topic, func(msg *nats.Msg) {
		err = json.Unmarshal(msg.Data, &response)
	})

	nc.FlushWithContext(context.Background())

	if err = nc.LastError(); err != nil {
		return
	}

	return
}
