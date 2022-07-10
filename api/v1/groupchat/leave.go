package groupchat

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	// leaveRequest model info
	// @Description Данные группы
	leaveRequest struct {
		// ID группы
		ID int64 `form:"id" binding:"required" example:"12"`
	} // @name groupchat.leaveRequest
)

// Leave godoc
// @Summary Покинуть групповой чат
// @Description Метод удаляет текущего пользователя из участников в группового чата
// @Tags groupchat
// @Security Basic
// @Accept multipart/form-data
// @Produce json
// @Param groupchatData formData groupchat.leaveRequest true "Группа"
// @Success 200 "Успешно"
// @Failure 400 "Неверный формат данных"
// @Failure 401 "Неверные данные для авторизации"
// @Router /api/v1/groupchat/leave/ [post]
func Leave(c *gin.Context) {
	var req leaveRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
	}

	c.Status(http.StatusOK)
}
