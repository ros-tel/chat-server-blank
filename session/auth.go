package session

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

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
		username, password, ok := basicAuth(c)
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

// BasicAuth returns the username and password provided in the request's
// Authorization header, if the request uses HTTP Basic Authentication.
// See RFC 2617, Section 2.
func basicAuth(c *gin.Context) (username, password string, ok bool) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		return
	}
	return parseBasicAuth(auth)
}

// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	// Case insensitive prefix match. See Issue 22736.
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}

// Генерация хеша 1С-Connect
func makeAuthHash(username, password string) string {
	h1 := fmt.Sprintf("%s:xbwg64okfdvb34fvi3:%s", username, password)
	hash := md5.Sum([]byte(h1))
	return fmt.Sprintf("%X", hash)
}
