package groupchat

import (
	"net/http"

	"chat-server/api/v1/objects"

	"github.com/gin-gonic/gin"
)

type (
	// createRequest model info
	// @Description Данные группы
	createRequest struct {
		// Название группы
		Name string `form:"name" binding:"required" example:"Новая"`
	} // @name groupchat.createRequest
)

// Create godoc
// @Summary Создание группового чата
// @Description Метод создает групповой чат
// @Tags groupchat
// @Security Basic
// @Accept multipart/form-data
// @Produce json
// @Param groupchatData formData groupchat.createRequest true "Группа"
// @Success 200 {object} objects.Groupchat
// @Failure 400 "Неверный формат данных"
// @Failure 401 "Неверные данные для авторизации"
// @Router /api/v1/groupchat/ [put]
func Create(c *gin.Context) {
	user_id := c.MustGet("user_id").(int64)

	var req createRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
	}

	response := objects.Groupchat{
		ID:          122,
		Name:        req.Name,
		UserOwnerID: user_id,
		Members: []objects.GroupchatMember{
			{
				User: objects.User{
					ID: user_id,
				},
				PermissionLevel: 9,
			},
		},
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}
