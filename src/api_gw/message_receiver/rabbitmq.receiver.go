package message_receiver

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQReceiver struct {
	connection   *amqp.Connection
	channel      *amqp.Channel
	exchangeName string

	InfoChannel    <-chan amqp.Delivery
	WarningChannel <-chan amqp.Delivery
	ErrorChannel   <-chan amqp.Delivery
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func CreateMessageReceiver(connection string, exchangeName string) *RabbitMQReceiver {
	receiver := &RabbitMQReceiver{}
	var err error

	receiver.connection, err = amqp.Dial(connection)
	failOnError(err, "Failed to connect to RabbitMQ")

	receiver.channel, err = receiver.connection.Channel()
	failOnError(err, "Failed to open a channel")

	err = receiver.channel.ExchangeDeclare(
		exchangeName, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	infoQueue, err := receiver.channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = receiver.channel.QueueBind(
		infoQueue.Name, // queue name
		"info",         // routing key
		exchangeName,   // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	warningQueue, err := receiver.channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = receiver.channel.QueueBind(
		warningQueue.Name, // queue name
		"warning",         // routing key
		exchangeName,      // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	errorQueue, err := receiver.channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = receiver.channel.QueueBind(
		errorQueue.Name, // queue name
		"error",         // routing key
		exchangeName,    // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	receiver.InfoChannel, err = receiver.channel.Consume(
		infoQueue.Name, // queue
		"",             // consumer
		true,           // auto ack
		false,          // exclusive
		false,          // no local
		false,          // no wait
		nil,            // args
	)
	failOnError(err, "Failed to register a consumer")

	receiver.WarningChannel, err = receiver.channel.Consume(
		warningQueue.Name, // queue
		"",                // consumer
		true,              // auto ack
		false,             // exclusive
		false,             // no local
		false,             // no wait
		nil,               // args
	)
	failOnError(err, "Failed to register a consumer")

	receiver.ErrorChannel, err = receiver.channel.Consume(
		errorQueue.Name, // queue
		"",              // consumer
		true,            // auto ack
		false,           // exclusive
		false,           // no local
		false,           // no wait
		nil,             // args
	)
	failOnError(err, "Failed to register a consumer")

	return receiver
}

func (receiver RabbitMQReceiver) Close() {
	_ = receiver.channel.Close()
	_ = receiver.connection.Close()
}
