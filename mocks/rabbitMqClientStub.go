package mocks

import "github.com/streadway/amqp"

// RabbitClientStub -
type RabbitClientStub struct {
	PublishCalled func(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
	DialCalled    func(url string) (*amqp.Connection, error)
}

// Publish -
func (rc *RabbitClientStub) Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	if rc.PublishCalled != nil {
		return rc.PublishCalled(exchange, key, mandatory, immediate, msg)
	}
	return nil
}

// Dial -
func (rc *RabbitClientStub) Dial(url string) (*amqp.Connection, error) {
	if rc.DialCalled != nil {
		return rc.DialCalled(url)
	}

	return nil, nil
}

// IsInterfaceNil -
func (rc *RabbitClientStub) IsInterfaceNil() bool {
	return false
}
