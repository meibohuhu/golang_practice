package event

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	connection *amqp.Connection
}

func NewEventEmitter(conn *amqp.Connection) (Emitter, error) {
	emitter := Emitter{
		connection: conn,
	}

	err := emitter.setup()
	if err != nil {
		return Emitter{}, err
	}

	return emitter, nil
}

func (e *Emitter) Push(event string, severity string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	log.Println("Pushing to channel")

	err = channel.Publish(
		"logs_topic",
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event), // we cast to a slice of bytes for Payload data sent from UI, like kafka message data
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (e *Emitter) setup() error {
	channel, err := e.connection.Channel() // build channel
	if err != nil {
		return err
	}

	defer channel.Close()
	return declareExchange(channel) // in event.go
}
