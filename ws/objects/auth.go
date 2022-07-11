package objects

const (
	MSG_TYPE_AUTH MessageType = "auth"
)

type (
	MessageAuth struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
)
