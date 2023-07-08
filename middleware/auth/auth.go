package auth

import (
	"github.com/gin-gonic/gin"
	"lendotopia.com/originator/config"
)

// TODO - rewrite this to use token in the request headers.

func TokenAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.Request.FormValue("api_token")
		cfg := config.Get()

		if token == "" {
			c.AbortWithStatus(400)
			return
		}
		if token != cfg.ApiToken {
			c.AbortWithStatus(401)
			return
		}

		c.Next()
	}

}
