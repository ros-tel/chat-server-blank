package list

import (
	"net/http"
	"time"

	"chat-server/api/v1/objects"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	chat_msg_id = uuid.New()

	chats = []objects.Chat{
		{
			User: objects.User{
				ID:      1,
				Name:    "Александр",
				Surname: "М",
			},
			LastReadMessageID: &chat_msg_id,
			LastMessage: &objects.Message{
				ID:       chat_msg_id,
				Time:     time.Now().Add(-10 * time.Minute),
				Type:     1,
				Message:  "Сообщение 1",
				AuthorID: 1,
			},
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
	}

	groupchat_msg_id = uuid.New()

	groupchats = []objects.Groupchat{
		{
			ID:                2324,
			Name:              "Общая",
			UserOwnerID:       1,
			LastReadMessageID: &groupchat_msg_id,
			LastMessage: &objects.Message{
				ID:       groupchat_msg_id,
				Time:     time.Now().Add(-10 * time.Minute),
				Type:     1,
				Message:  "Сообщение 6",
				AuthorID: 6,
			},
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
		},
	}
)

type (
	ContactsResponse struct {
		Chats      []objects.Chat      `json:"chats"`
		Groupchats []objects.Groupchat `json:"groupchats"`
	}
)

// Contacts godoc
// @Summary Получение информации о контактах
// @Description Метод возвращает информацию о контактах текущего пользователя
// @Tags list
// @Security Basic
// @Accept json
// @Produce json
// @Success 200 {object} list.ContactsResponse
// @Failure 401 "Неверные данные для авторизации"
// @Router /api/v1/list/contacts/ [get]
func Contacts(c *gin.Context) {
	user_id := c.MustGet("user_id").(int64)

	var response = ContactsResponse{}

	for _, chat := range chats {
		// Исключаем запросившего
		if chat.User.ID == user_id {
			continue
		}

		response.Chats = append(response.Chats, chat)
	}

	response.Groupchats = groupchats

	c.JSON(
		http.StatusOK,
		response,
	)
}
