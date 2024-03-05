package middleware

import (
	jwt "golang_testing_grpc/pkg/jtoken"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return JWT(jwt.AccessTokenType)
}

func JWTRefresh() gin.HandlerFunc {
	return JWT(jwt.RefreshTokenType)
}

func JWT(tokenType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		payload, err := jwt.ValidateToken(token)
		if err != nil || payload == nil || payload["type"] != tokenType {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		c.Set("userId", payload["id"])
		c.Set("role", payload["role"])
		c.Next()
	}
}
