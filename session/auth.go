package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Auth struct {
		id       int64
		password string
	}
)

var (
	auths = map[string]Auth{
		"alexander": {
			id:       1,
			password: "pass",
		},
		"alisa": {
			id:       2,
			password: "pass",
		},
		"vitaly": {
			id:       3,
			password: "pass",
		},
		"vladimir": {
			id:       4,
			password: "pass",
		},
		"lada": {
			id:       5,
			password: "pass",
		},
		"semen": {
			id:       6,
			password: "pass",
		},
	}
)

func SetSessionVars() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if ok {
			if a, ok := auths[username]; ok {
				if password == a.password {
					c.Set("user_id", a.id)
				}
			}
		}
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exist := c.Get("user_id"); !exist {
			// Credentials doesn't match, we return 401 and abort handlers chain.
			c.Header("WWW-Authenticate", "Basic realm=\"Secure Area\"")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
