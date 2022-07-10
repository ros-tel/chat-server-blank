package v1

import (
	"chat-server/api/v1/auth"
	"chat-server/api/v1/chat"
	"chat-server/api/v1/groupchat"
	"chat-server/api/v1/list"
	"chat-server/session"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {

	v1authorized := r.Group("/api/v1")
	v1authorized.Use(session.AuthRequired())
	{
		v1authorized.GET("/auth/who/", auth.Who)

		v1authorized.GET("/list/contacts/", list.Contacts)

		v1authorized.GET("/chat/", chat.ChatHistory)
		v1authorized.POST("/chat/add/", chat.Add)
		v1authorized.POST("/chat/hide/", chat.Hide)

		v1authorized.GET("/groupchat/", groupchat.GroupchatHistory)
		v1authorized.PUT("/groupchat/", groupchat.Create)
		v1authorized.PATCH("/groupchat/", groupchat.Edit)
		v1authorized.POST("/groupchat/join/", groupchat.Join)
		v1authorized.POST("/groupchat/leave/", groupchat.Leave)
	}
}
