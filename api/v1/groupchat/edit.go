package groupchat

import (
	"chat-server/api/v1/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	// editRequest model info
	// @Description Данные группы
	editRequest struct {
		// ID группы
		ID int64 `form:"id" binding:"required" example:"232"`
		// Название группы
		Name string `form:"name" binding:"required" example:"Новая"`
	} // @name groupchat.editRequest
)

// Edit godoc
// @Summary Изменение группового чата
// @Description Метод изменяет групповой чат
// @Tags groupchat
// @Security Basic
// @Accept multipart/form-data
// @Produce json
// @Param groupchatData formData groupchat.editRequest true "Группа"
// @Success 200 {object} objects.Groupchat
// @Failure 400 "Неверный формат данных"
// @Failure 401 "Неверные данные для авторизации"
// @Router /api/v1/groupchat/ [patch]
func Edit(c *gin.Context) {
	user_id := c.MustGet("user_id").(int64)

	var req editRequest
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
		ID:          req.ID,
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
