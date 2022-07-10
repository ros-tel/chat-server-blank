package objects

import "github.com/google/uuid"

type (
	// Groupchat model info
	// @Description Участник группового чата
	GroupchatMember struct {
		// Информация по пользователю
		User User `json:"user"`
		// Уровень прав
		// * 0 - участник
		// * 9 - админ
		PermissionLevel int `json:"permission_level" example:"0"`
	} // @name GroupchatMember

	// Groupchat model info
	// @Description Групповые чаты
	Groupchat struct {
		// ID конференции
		ID int64 `json:"id" example:"103"`
		// Название группы
		Name string `json:"name" example:"Общая группа"`
		// ID пользователя владеющего (создателя) группой
		UserOwnerID int64 `json:"user_owner_id" example:"13"`
		// Информация по пользователю
		Members []GroupchatMember `json:"user"`
		// ID последнего прочитанного мной сообщения
		// * null когда текущий пользователь не читал сообщения
		LastReadMessageID *uuid.UUID `json:"last_read_message_id" extensions:"x-nullable" example:"10a3d4e7-c1b0-471a-bbfe-c168902354e8"`
		// Последнее сообщения в чате
		// * null когда в чате не сообщений
		LastMessage *Message `json:"last_message" extensions:"x-nullable"`
		// Количество непрочитанных сообщений у пользователя для запросившего!
		UnreadCount int64 `json:"unread_count" example:"3"`
	} // @name Groupchat
)
