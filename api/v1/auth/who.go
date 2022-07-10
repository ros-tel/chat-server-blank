package auth

import (
	"net/http"

	"chat-server/api/v1/objects"

	"github.com/gin-gonic/gin"
)

var (
	users = map[int64]objects.User{
		1: {
			ID:      1,
			Name:    "Александр",
			Surname: "М",
		},
		2: {
			ID:      2,
			Name:    "Алиса",
			Surname: "Ш",
		},
		3: {
			ID:      3,
			Name:    "Виталий",
			Surname: "П",
		},
		4: {
			ID:      4,
			Name:    "Владимир",
			Surname: "В",
		},
		5: {
			ID:      5,
			Name:    "Лада",
			Surname: "С",
		},
		6: {
			ID:      6,
			Name:    "Семён",
			Surname: "Ш",
		},
	}
)

// Who godoc
// @Summary Получение информации о текущем пользователе
// @Description Метод возвращает информацию о текущем пользователе
// @Tags auth
// @Security Basic
// @Accept json
// @Produce json
// @Success 200 {object} objects.User
// @Failure 401 "Неверные данные для авторизации"
// @Failure 500 "Ошибка сервера"
// @Router /api/v1/auth/who/ [get]
func Who(c *gin.Context) {
	user_id := c.MustGet("user_id").(int64)

	if user, ok := users[user_id]; ok {
		c.JSON(
			http.StatusOK,
			user,
		)

		return
	}

	c.Status(http.StatusInternalServerError)
}
