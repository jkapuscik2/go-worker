package MakeHash

import (
	"golang.org/x/crypto/bcrypt"
	"jkapuscik2/go-worker/src/Services/Messages"
	"log"
)

func Handle(msg Messages.MakeHashMsq) {
	hash, err := bcrypt.GenerateFromPassword([]byte(msg.Data.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Printf("Calculated hash: %s for %s", hash, msg.Data.Password)
	}
}
