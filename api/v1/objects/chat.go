package objects

import "github.com/google/uuid"

type (
	// Chat model info
	// @Description Чат p2p
	Chat struct {
		// Информация по пользователю
		User User `json:"user"`
		// ID последнего прочитанного мной сообщения
		// * null когда текущий пользователь не читал сообщения
		LastReadMessageID *uuid.UUID `json:"last_read_message_id" extensions:"x-nullable" example:"10a3d4e7-c1b0-471a-bbfe-c168902354e8"`
		// Последнее сообщения в чате
		// * null когда в чате не сообщений
		LastMessage *Message `json:"last_message" extensions:"x-nullable"`
		// Количество непрочитанных сообщений у пользователя для запросившего!
		UnreadCount int64 `json:"unread_count" example:"3"`
	} // @name Chat
)
