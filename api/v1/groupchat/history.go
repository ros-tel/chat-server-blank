package groupchat

import (
	"net/http"
	"time"

	"chat-server/api/v1/objects"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	msgs = []objects.Message{
		{
			ID:       uuid.New(),
			Time:     time.Now().Add(-2 * 24 * time.Hour),
			Type:     1,
			Message:  "Сообщение 1",
			AuthorID: 1,
		},
		{
			ID:       uuid.New(),
			Time:     time.Now().Add(-1 * 24 * time.Hour),
			Type:     1,
			Message:  "Сообщение 2",
			AuthorID: 2,
		},
		{
			ID:       uuid.New(),
			Time:     time.Now().Add(-23 * time.Hour),
			Type:     1,
			Message:  "Сообщение 3",
			AuthorID: 3,
		},
		{
			ID:       uuid.New(),
			Time:     time.Now().Add(-2 * time.Hour),
			Type:     1,
			Message:  "Сообщение 4",
			AuthorID: 4,
		},
		{
			ID:       uuid.New(),
			Time:     time.Now().Add(-20 * time.Minute),
			Type:     1,
			Message:  "Сообщение 5",
			AuthorID: 5,
		},
		{
			ID:       uuid.New(),
			Time:     time.Now().Add(-10 * time.Minute),
			Type:     1,
			Message:  "Сообщение 6",
			AuthorID: 6,
		},
	}
)

type (
	historyRequest struct {
		ID         int64   `form:"id" binding:"required" example:"12"`
		Date_Start *string `form:"date_start" binding:"omitempty,datetime=2006-01-02T15:04:05.999999Z07:00" example:"2019-12-17T13:16:24.550Z"`
		Dir        string  `form:"dir" binding:"omitempty,eq=prev|eq=next" enums:"prev,next" example:"prev"`
		Limit      *int    `form:"limit" binding:"omitempty,min=0,max=1000" example:"10"`
	}
)

// GroupchatHistory godoc
// @Summary Получение истории сообщений по групповому чату
// @Description Метод возвращает историю сообщений по групповому чату
// @Tags groupchat
// @Security Basic
// @Accept json
// @Produce json
// @Param id query int64 true "ID чата по которому необходимо вернуть данные" example(12)
// @Param date_start query string false "Дата начала интервала выборки, если не указана, то берется начало списка (последнее или первое)"
// @Param dir query string false "Направление запроса" Enums(prev,next) example(prev)
// @Param limit query int false "Количество объектов в результате" example(10)
// @Success 200 {array} objects.Message
// @Failure 400 "Неверный формат данных"
// @Failure 401 "Неверные данные для авторизации"
// @Router /api/v1/groupchat/ [get]
func GroupchatHistory(c *gin.Context) {
	// user_id := c.MustGet("user_id").(int64)

	var req historyRequest
	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
	}

	limit := 50
	if req.Limit != nil {
		limit = *req.Limit
	}

	var messages = []objects.Message{}

	for i, msg := range msgs {
		// Достигли лимита
		if i >= limit {
			break
		}

		messages = append(messages, msg)
	}

	c.JSON(
		http.StatusOK,
		messages,
	)
}
