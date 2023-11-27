package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/oauth2.v3/server"
)

func Auth(srv *server.Server) gin.HandlerFunc {

	return func(c *gin.Context) {
		_, err := srv.ValidationBearerToken(c.Request)
		if err != nil {
			c.JSON(http.StatusBadGateway, map[string]string{"message": "Failed"})
			c.Abort()
			return
		}

		c.Next()
	}
}
