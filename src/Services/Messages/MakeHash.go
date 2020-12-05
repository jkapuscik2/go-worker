package Messages

type MakeHashData struct {
	Password string `json:"password"`
}

type MakeHashMsq struct {
	Msg
	Data MakeHashData `json:"data"`
}
