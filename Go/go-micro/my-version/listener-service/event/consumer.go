package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

// the events pushed to MQ
type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

// function1: new consumer in main
func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}
	err := consumer.setup() // use new consumer to build channel
	if err != nil {
		return Consumer{}, err
	}
	return consumer, nil
}

// build channel
func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	return declareExchange(channel)
}

// function2: listnen messages
func (consumer *Consumer) Listen(keys []string) error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	q, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	// now we have channel and keys
	for _, s := range keys {
		ch.QueueBind(
			q.Name,
			s, // key
			"logs_topic",
			false,
			nil,
		)
		if err != nil {
			return err
		}
	}
	fmt.Printf("[queue name %s]\n", q.Name)
	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	// do this forever, consumer and handle messages until i exit application => channel
	forever := make(chan bool)
	go func() {
		for d := range messages {
			var payload Payload
			_ = json.Unmarshal(d.Body, &payload)

			go handlePayload(payload)
		}
	}()

	fmt.Printf("Waiting for message [Exchange, Queue] [logs_topic, %s]\n", q.Name)
	<-forever

	return nil
}

// this consumer will implement this handlePayload function
func handlePayload(payload Payload) {
	switch payload.Name {
	case "log", "event":
		// log whatever we get
		err := logEvent(payload)
		if err != nil {
			log.Println(err)
		}

	case "auth":
		// authenticate

	// you can have as many cases as you want, as long as you write the logic

	default: // safety measure
		err := logEvent(payload)
		if err != nil {
			log.Println(err)
		}
	}
}

// moved from broker-service to here
func logEvent(entry Payload) error {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return err
	}

	return nil
}
