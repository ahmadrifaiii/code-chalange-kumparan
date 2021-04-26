package nats

type Store interface {
	Close()
	Publish(meow interface{}) error
	Subscription() (<-chan NatsMessage, error)
	On(f func(NatsMessage)) error
}

var impl Store

func SetEventStore(s Store) {
	impl = s
}

func Close() {
	impl.Close()
}

func Publish(p interface{}) error {
	return impl.Publish(p)
}

func Subscription() (<-chan NatsMessage, error) {
	return impl.Subscription()
}

func On(f func(NatsMessage)) error {
	return impl.On(f)
}
