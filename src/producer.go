package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"jkapuscik2/go-worker/src/Services/Messages"
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

	for i := 1; i <= 100; i++ {

		msg := Messages.MakeHashMsq{
			Msg: Messages.Msg{
				Type: "MakeHash",
				Date: time.Now(),
			},
			Data: Messages.MakeHashData{Password: fmt.Sprintf("msq num %d", i)},
		}
		jsonMsg, _ := json.Marshal(msg)

		r.Publish(string(jsonMsg), os.Getenv("RABBITMQ_QUEUE"))
	}
}
