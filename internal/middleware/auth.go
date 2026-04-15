package middleware

import (
	"strings"

	jwtpkg "video-feed/pkg/jwt"
	"video-feed/pkg/response"

	"github.com/gin-gonic/gin"
)

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, 1002, "unauthorized")
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.Error(c, 1002, "invalid token format")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwtpkg.ParseToken(secret, tokenString)
		if err != nil {
			response.Error(c, 1002, "invalid token")
			c.Abort()
			return
		}

		// 这里先简单验证 token 是否存在，后续再完善 JWT 验证逻辑

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
