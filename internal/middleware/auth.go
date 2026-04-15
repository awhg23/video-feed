package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"video-feed/pkg/response"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, 1002, "unauthorized: missing token")
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.Error(c, 1002, "unauthorized: invalid token format")
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			response.Error(c, 1002, "unauthorized: empty token")
			c.Abort()
			return
		}

		// 这里先简单验证 token 是否存在，后续再完善 JWT 验证逻辑

		c.Set("userID", int64(1))
		c.Next()
	}
}
