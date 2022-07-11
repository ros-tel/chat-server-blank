package ws

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	r.GET("/ws/",
		func(c *gin.Context) {
			wsHandler(c.Writer, c.Request)
		},
	)
}
