package chat

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	// hideRequest model info
	// @Description Пользователь
	hideRequest struct {
		// ID Пользователя
		ID int64 `form:"id" binding:"required" example:"232"`
	} // @name chat.hideRequest
)

// Hide godoc
// @Summary Удалить p2p-контакт
// @Description Метод удаляет p2p-контакт из списока контактов текущего пользователя
// @Tags chat
// @Security Basic
// @Accept multipart/form-data
// @Produce json
// @Param userData formData chat.hideRequest true "Котакт"
// @Success 200 "Успешно"
// @Failure 400 "Неверный формат данных"
// @Failure 401 "Неверные данные для авторизации"
// @Router /api/v1/chat/hide/ [post]
func Hide(c *gin.Context) {
	var req hideRequest
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
