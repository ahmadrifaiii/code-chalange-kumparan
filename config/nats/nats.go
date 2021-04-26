package nats

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type NatsEvent struct {
	nc              *nats.Conn
	Subs            *nats.Subscription
	natsMessageChan chan NatsMessage
}

func NewNats(url string) (*NatsEvent, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsEvent{nc: nc}, nil
}

func (es *NatsEvent) EventSubscribe() (<-chan NatsMessage, error) {
	m := NatsMessage{}
	es.natsMessageChan = make(chan NatsMessage, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	es.Subs, err = es.nc.ChanSubscribe(m.Key(), ch)
	if err != nil {
		return nil, err
	}
	// Decode message
	go func() {
		for {
			select {
			case msg := <-ch:
				if err := es.readMessage(msg.Data, &m); err != nil {
					log.Fatal(err)
				}
				es.natsMessageChan <- m
			}
		}
	}()
	return (<-chan NatsMessage)(es.natsMessageChan), nil
}

func (es *NatsEvent) EventOn(f func(NatsMessage)) (err error) {
	m := NatsMessage{}
	es.Subs, err = es.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		if err := es.readMessage(msg.Data, &m); err != nil {
			log.Fatal(err)
		}
		f(m)
	})
	return
}

func (es *NatsEvent) EventClose() {
	if es.nc != nil {
		es.nc.Close()
	}
	if es.Subs != nil {
		if err := es.Subs.Unsubscribe(); err != nil {
			log.Fatal(err)
		}
	}
	close(es.natsMessageChan)
}

func (es *NatsEvent) EventPublish(p *interface{}) error {
	m := NatsMessage{
		ID:      uuid.New().String(),
		Subject: "",
		Data:    p,
	}
	data, err := es.writeMessage(&m)
	if err != nil {
		return err
	}
	return es.nc.Publish(m.Key(), data)
}

func (es *NatsEvent) writeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (es *NatsEvent) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}
