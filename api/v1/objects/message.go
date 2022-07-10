package objects

import (
	"time"

	"github.com/google/uuid"
)

type (
	// Message model info
	// @Description Сообщение в чате
	Message struct {
		// ID сообщения
		ID uuid.UUID `json:"id" example:"d096b756-8b38-4b00-b3b2-e23946400fdf"`
		// Время сообщения
		Time time.Time `json:"time"`
		// Тип сообщения
		// * 1 - текстовое сообщение
		Type int `json:"type" example:"1"`
		// Тело сообщения
		Message string `json:"message" example:"Превед медвед!"`
		// ID автора (агент)
		AuthorID int64 `json:"author_id" example:"50"`
	} // @name Message
)
