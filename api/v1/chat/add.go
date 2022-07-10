package chat

import (
	"net/http"

	"chat-server/api/v1/objects"

	"github.com/gin-gonic/gin"
)

type (
	// addRequest model info
	// @Description Пользователь
	addRequest struct {
		// ID Пользователя
		ID int64 `form:"id" binding:"required" example:"232"`
	} // @name chat.addRequest
)

// Add godoc
// @Summary Добавить p2p-контакт
// @Description Метод добавляет p2p-контакт в список контактов текущего пользователя
// @Tags chat
// @Security Basic
// @Accept multipart/form-data
// @Produce json
// @Param userData formData chat.addRequest true "Котакт"
// @Success 200 {object} objects.Chat
// @Failure 400 "Неверный формат данных"
// @Failure 401 "Неверные данные для авторизации"
// @Router /api/v1/chat/add/ [post]
func Add(c *gin.Context) {
	var req addRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
	}

	response := objects.Chat{
		User: objects.User{
			ID:      req.ID,
			Name:    "Иван",
			Surname: "Иванов",
		},
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}
