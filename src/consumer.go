package main

import (
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"golang.org/x/crypto/bcrypt"
	"jkapuscik2/go-worker/src/services/rabbitmq"
	"log"
	"os"
	"time"
)

func init() {
	godotenv.Load()
}

func main() {
	r := rabbitmq.Rabbit{}
	r.Connect()
	defer r.Close()

	msgs, _ := r.Consume(os.Getenv("RABBITMQ_QUEUE"))

	start := time.Now()

	for d := range msgs {
		go handle(d, start)
	}
}

func handle(msg amqp.Delivery, start time.Time) {
	log.Printf("Received a message: %s", msg.Body)
	hash, _ := bcrypt.GenerateFromPassword(msg.Body, bcrypt.DefaultCost)
	log.Printf("Calculated hash: %s", hash)

	log.Printf("Time from start %s", time.Now().Sub(start))
	msg.Ack(false)
}
