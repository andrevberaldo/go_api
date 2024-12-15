package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString != "fake_token_JWT" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized!",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
