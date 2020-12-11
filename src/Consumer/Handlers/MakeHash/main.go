package MakeHash

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"jkapuscik2/go-worker/src/Messages"
	logger "jkapuscik2/go-worker/src/Services/Logger"
)

func Handle(msg Messages.MakeHashMsq, logger logger.Logger) {
	hash, err := bcrypt.GenerateFromPassword([]byte(msg.Data.Password), bcrypt.DefaultCost)

	if err != nil {
		logger.Log(fmt.Sprintf(err.Error()))
	} else {
		logger.Log(fmt.Sprintf("Calculated hash: %s for %s", hash, msg.Data.Password))
	}
}
