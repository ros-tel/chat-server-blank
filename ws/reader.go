package ws

import (
	"chat-server/ws/objects"
	"log"
)

func reader(c *objects.Conn) {
	for {
		msg, err := c.Read()
		if err != nil {
			break
		}

		log.Printf("msg: %+v", msg)

		switch msg.Type {
		case objects.MSG_TYPE_CHAT, objects.MSG_TYPE_GROUPCHAT:
			if msg.To == 0 {
				c.SendMessage(
					objects.Message{
						Type: objects.MSG_TYPE_ERROR,
						Result: &objects.MessageResult{
							Error: "field \"to\" required",
						},
					},
				)

				continue
			}
			if msg.Chat == nil {
				c.SendMessage(
					objects.Message{
						Type: objects.MSG_TYPE_ERROR,
						Result: &objects.MessageResult{
							Error: "field \"chat\" required",
						},
					},
				)

				continue
			}

			switch msg.Chat.Type {
			case objects.CHAT_TYPE_TEXT:
				if msg.Chat.Text == nil {
					c.SendMessage(
						objects.Message{
							Type: objects.MSG_TYPE_ERROR,
							Result: &objects.MessageResult{
								Error: "field \"chat.text\" required",
							},
						},
					)

					continue
				}
			default:
				c.SendMessage(
					objects.Message{
						Type: objects.MSG_TYPE_ERROR,
						Result: &objects.MessageResult{
							Error: "field \"chat.type\" unknown",
						},
					},
				)

				continue
			}

			// Выставляем отправителя
			msg.From = c.ID

			// Отправляем echo
			c.SendMessage(msg)
		default:
			c.SendMessage(
				objects.Message{
					Type: objects.MSG_TYPE_ERROR,
					Result: &objects.MessageResult{
						Error: "Unknown type",
					},
				},
			)
		}
	}
}
