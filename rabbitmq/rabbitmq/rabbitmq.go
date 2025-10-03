package rabbitmq

import (
	"fmt"
	"rabbitmq/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbitmq struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func Connect(config config.Config) (Rabbitmq, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s",
		config.Rabbitmq.RabbitmqUser,
		config.Rabbitmq.RabbitmqPassword,
		config.Rabbitmq.RabbitmqHost,
		config.Rabbitmq.RabbitmqPort))
	if err != nil {
		return Rabbitmq{}, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return Rabbitmq{}, err
	}

	return Rabbitmq{
		Conn:    conn,
		Channel: ch,
	}, err
}

func (r Rabbitmq) DeclareQueue(qName string) (amqp.Queue, error) {
	q, err := r.Channel.QueueDeclare(qName, false, false, false, false, nil)
	if err != nil {
		return q, err
	}
	return q, err
}

func (r Rabbitmq) PublishMessage(qName string, message string) error {
	err := r.Channel.Publish("", qName, false, false, amqp.Publishing{

		ContentType: "text/plain",
		Body:        []byte(message),
	})
	return err
}

func (r Rabbitmq) Consumer(qName string) (<-chan amqp.Delivery, error) {
	msgs, err := r.Channel.Consume(qName, "", true, false, false, false, nil)
	return msgs, err
}
