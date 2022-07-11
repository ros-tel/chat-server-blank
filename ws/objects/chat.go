package objects

const (
	MSG_TYPE_CHAT      MessageType = "chat"
	MSG_TYPE_GROUPCHAT MessageType = "groupchat"

	CHAT_TYPE_TEXT ChatType = "text"
)

type (
	ChatType string

	MessageChat struct {
		Type ChatType `json:"type"`
		Text *string  `json:"text,omitempty"`
	}
)
