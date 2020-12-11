package logger

import (
	"fmt"
	"log"
)

type Logger struct {
	Uid string
}

func (l Logger) Log(msg string) {
	log.Println(fmt.Sprintf("%s: %s", l.Uid, msg))
}
