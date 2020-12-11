package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"jkapuscik2/go-worker/src/Consumer/Handlers/MakeHash"
	"jkapuscik2/go-worker/src/Messages"
	"jkapuscik2/go-worker/src/Services/Logger"
	"jkapuscik2/go-worker/src/Services/Rabbitmq"
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
	logger := logger.Logger{uuid.New().String()}
	logger.Log(fmt.Sprintf("Received a message: %s", m.Body))

	defer logger.Log(fmt.Sprintf("Time from start %s", time.Now().Sub(appStart)))
	defer logger.Log(fmt.Sprintf("Msg handled in %s", time.Now().Sub(start)))
	defer m.Ack(false)

	msg := Messages.Msg{}
	err := json.Unmarshal(m.Body, &msg)

	if err != nil {
		logger.Log(fmt.Sprintf(fmt.Sprintf("Invalid msg received %s: %s", m.Body, err.Error())))
		return
	}

	switch msg.Type {
	case "MakeHash":
		logger.Log(fmt.Sprintf("Received make hash msg"))
		makeHashMsg := Messages.MakeHashMsq{}
		err := json.Unmarshal(m.Body, &makeHashMsg)
		if err != nil {
			logger.Log(fmt.Sprintf("Invalid make hash msg received %s: %s", m.Body, err.Error))
			return
		}

		MakeHash.Handle(makeHashMsg, logger)
	}
}
