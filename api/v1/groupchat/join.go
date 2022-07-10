package groupchat

import (
	"chat-server/api/v1/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	// joinRequest model info
	// @Description Данные группы
	joinRequest struct {
		// ID группы
		ID int64 `form:"id" binding:"required" example:"232"`
	} // @name groupchat.joinRequest
)

// Join godoc
// @Summary Присоединиться к групповому чату
// @Description Метод добавляет текущего пользователя в участники в группового чата
// @Tags groupchat
// @Security Basic
// @Accept multipart/form-data
// @Produce json
// @Param groupchatData formData groupchat.joinRequest true "Группа"
// @Success 200 {object} objects.Groupchat
// @Failure 400 "Неверный формат данных"
// @Failure 401 "Неверные данные для авторизации"
// @Router /api/v1/groupchat/join/ [post]
func Join(c *gin.Context) {
	var req joinRequest
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
		Name:        "Общая",
		UserOwnerID: 1,
		Members: []objects.GroupchatMember{
			{
				User: objects.User{
					ID:      1,
					Name:    "Александр",
					Surname: "М",
				},
				PermissionLevel: 9,
			},
			{
				User: objects.User{
					ID:      2,
					Name:    "Алиса",
					Surname: "Ш",
				},
			},
			{
				User: objects.User{
					ID:      3,
					Name:    "Виталий",
					Surname: "П",
				},
			},
			{
				User: objects.User{
					ID:      4,
					Name:    "Владимир",
					Surname: "В",
				},
			},
			{
				User: objects.User{
					ID:      5,
					Name:    "Лада",
					Surname: "С",
				},
			},
			{
				User: objects.User{
					ID:      6,
					Name:    "Семён",
					Surname: "Ш",
				},
			},
		},
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}
