//go:generate swag init -o swagger -g app.go

// @schemes http
// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @tag.name auth
// @tag.description Методы авторизации и управления сессией

// @tag.name list
// @tag.description Методы получения контактной информации

// @tag.name chat
// @tag.description Методы получения истории сообщений и управления p2p-чатами

// @tag.name groupchat
// @tag.description Методы получения истории сообщений и управления групповыми чатами

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
package main

import (
	v1 "chat-server/api/v1"
	"chat-server/session"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(
		session.SetSessionVars(),
	)

	v1.InitRoutes(r)

	r.Static("/swagger/", "./swagger")

	r.Run("0.0.0.0:8080")
}
