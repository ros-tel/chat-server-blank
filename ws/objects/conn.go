package objects

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/gorilla/websocket"
)

type (
	Conn struct {
		ID   int64 `json:"id"`
		conn *websocket.Conn
	}
)

func NewConn(id int64, conn *websocket.Conn) *Conn {
	return &Conn{
		ID:   id,
		conn: conn,
	}
}

func (c *Conn) SendMessage(message Message) {
	msg, err := json.Marshal(message)
	if err != nil {
		log.Printf("[ERROR] json.Marshal %+v", err)
	}

	c.conn.WriteMessage(1, msg)
}

func (c *Conn) Read() (Message, error) {
	var (
		t       int
		msg     []byte
		message Message
		err     error
	)
	t, msg, err = c.conn.ReadMessage()
	if err != nil {
		return message, err
	}
	if t != 1 {
		log.Printf("[ERROR] Unknown Message: %d", t)
		return message, errors.New("unknown Message")
	}

	err = json.Unmarshal(msg, &message)
	if err != nil {
		log.Printf("[ERROR] conn.ReadMessage: %+v", err)
		return message, err
	}

	return message, err
}
