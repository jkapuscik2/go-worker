package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"jkapuscik2/go-worker/src/services/rabbitmq"
	"os"
)

func init() {
	godotenv.Load()
}

func main() {
	r := rabbitmq.Rabbit{}
	r.Connect()

	for i := 1; i <= 1000; i++ {
		r.Publish(fmt.Sprintf("msq num %d", i), os.Getenv("RABBITMQ_QUEUE"))
	}
}
