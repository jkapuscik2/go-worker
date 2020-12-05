package Messages

import "time"

type Msg struct {
	Type string    `json:"type"`
	Date time.Time `json:"date"`
}
