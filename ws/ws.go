package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"chat-server/session"
	"chat-server/ws/objects"

	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[ERROR] wsupgrader.Upgrade: %+v", err)
		return
	}
	defer conn.Close()

	t, msg, err := conn.ReadMessage()
	if err != nil {
		log.Printf("[ERROR] conn.ReadMessage: %+v", err)
		return
	}

	if t != 1 {
		log.Printf("[ERROR] Unknown Message: %d", t)
		return
	}

	var message objects.Message
	err = json.Unmarshal(msg, &message)
	if err != nil {
		log.Printf("[ERROR] conn.ReadMessage: %+v", err)
		return
	}

	if message.Type != objects.MSG_TYPE_AUTH {
		log.Printf("[ERROR] Unknown message type: %s", message.Type)
		return
	}

	if message.Auth == nil {
		log.Printf("[ERROR] field \"auth\" required: %s", message.Type)
		return
	}

	user_id, ok := session.CheckCred(message.Auth.Login, message.Auth.Password)
	if !ok {
		log.Printf("[ERROR] conn.ReadMessage: %+v", err)
		return
	}

	c := objects.NewConn(user_id, conn)

	// Отправим что авторизация удалась
	message.Auth = nil
	message.Result = new(objects.MessageResult)
	message.Result.Code = objects.MSG_RESULT_CODE_OK
	c.SendMessage(message)

	reader(c)
}
