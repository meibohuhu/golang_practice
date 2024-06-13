package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// declare channel
func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable?
		false,        // auto-deleted?  // auto-deleted: the exchange will not be automatically deleted when no queues are bound to it
		false,        // internal?
		false,        // no-wait?
		nil,          // arguements?
	)
}

// take a channel return the queue
// In RabbitMQ, a queue is a basic data structure used to store and forward messages. It's a buffer that stores messages sent by producers and delivers them to consumers.
// Queues decouple producers from consumers – producers send messages to a queue, and consumers retrieve messages from the queue.
func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name?
		false, // durable?
		false, // delete when unused?
		true,  // exclusive?
		false, // no-wait?
		nil,   // arguments?
	)
}
