package Rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"runtime"
)

type Rabbit struct {
	channel *amqp.Channel
}

func (r Rabbit) Publish(msg string, queue string) {
	err := r.channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         []byte(msg),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", msg)
}

func (r Rabbit) Consume(queue string) (<-chan amqp.Delivery, error) {
	return r.channel.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}

func (r *Rabbit) Connect() {
	conn, err := amqp.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%s",
			os.Getenv("RABBITMQ_USER"),
			os.Getenv("RABBITMQ_PASSWORD"),
			os.Getenv("RABBITMQ_HOST"),
			os.Getenv("RABBITMQ_PORT"),
		),
	)
	ch, err := conn.Channel()
	failOnError(err, "Failed to connect to RabbitMQ")

	err = ch.Qos(
		runtime.NumCPU(),
		0,
		false,
	)

	r.channel = ch
}

func (r Rabbit) Close() {
	r.channel.Close()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
