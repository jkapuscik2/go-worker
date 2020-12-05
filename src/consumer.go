package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"jkapuscik2/go-worker/src/Consumer/Handlers/MakeHash"
	"jkapuscik2/go-worker/src/Services/Messages"
	"jkapuscik2/go-worker/src/Services/Rabbitmq"
	"log"
	"os"
	"time"
)

func init() {
	godotenv.Load()
}

func main() {
	r := Rabbitmq.Rabbit{}
	r.Connect()
	defer r.Close()

	msgs, _ := r.Consume(os.Getenv("RABBITMQ_QUEUE"))

	start := time.Now()

	for d := range msgs {
		go handle(d, start)
	}
}

func handle(m amqp.Delivery, appStart time.Time) {
	start := time.Now()
	log.Printf("Received a message: %s", m.Body)

	defer log.Printf("Time from start %s", time.Now().Sub(appStart))
	defer log.Printf("Msg handled in %s", time.Now().Sub(start))
	defer m.Ack(false)

	msg := Messages.Msg{}
	err := json.Unmarshal(m.Body, &msg)

	if err != nil {
		log.Printf("Invalid msg received %s: %s", m.Body, err.Error())
		return
	}

	switch msg.Type {
	case "MakeHash":
		log.Println("Received make hash msg")
		makeHashMsg := Messages.MakeHashMsq{}
		err := json.Unmarshal(m.Body, &makeHashMsg)
		if err != nil {
			log.Printf("Invalid make hash msg received %s: %s", m.Body, err.Error)
			return
		}

		MakeHash.Handle(makeHashMsg)
	}
}
