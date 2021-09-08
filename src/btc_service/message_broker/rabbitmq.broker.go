package message_broker

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQBroker struct {
	connection   *amqp.Connection
	channel      *amqp.Channel
	exchangeName string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func CreateMessageBroker(connection string, exchangeName string) *RabbitMQBroker {
	broker := &RabbitMQBroker{}
	var err error

	broker.connection, err = amqp.Dial(connection)
	failOnError(err, "Failed to connect to RabbitMQ")

	broker.channel, err = broker.connection.Channel()
	failOnError(err, "Failed to open a channel")

	broker.exchangeName = exchangeName

	err = broker.channel.ExchangeDeclare(
		exchangeName, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	return broker
}

func (broker RabbitMQBroker) SendInfo(msg string) {
	err := broker.channel.Publish(
		broker.exchangeName, // exchange
		"info",              // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	failOnError(err, "Failed to publish info message")
}

func (broker RabbitMQBroker) SendWarning(msg string) {
	err := broker.channel.Publish(
		broker.exchangeName, // exchange
		"warning",           // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	failOnError(err, "Failed to publish warning message")
}

func (broker RabbitMQBroker) SendError(msg string) {
	err := broker.channel.Publish(
		broker.exchangeName, // exchange
		"error",             // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	failOnError(err, "Failed to publish error message")
}

func (broker RabbitMQBroker) Close() {
	_ = broker.channel.Close()
	_ = broker.connection.Close()
}
