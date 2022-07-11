package objects

type (
	MessageType string

	Message struct {
		ID     string         `json:"id,omitempty"`
		Type   MessageType    `json:"type"`
		From   int64          `json:"from,omitempty"`
		To     int64          `json:"to,omitempty"`
		Auth   *MessageAuth   `json:"auth,omitempty"`
		Chat   *MessageChat   `json:"chat,omitempty"`
		Result *MessageResult `json:"result,omitempty"`
	}
)
