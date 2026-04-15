package router

import (
	"video-feed/internal/handler"
	"video-feed/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Health *handler.HealthHandler
	Auth   *handler.AuthHandler
	User   *handler.UserHandler
}

func NewRouter(h *Handlers, jwtSecret string) *gin.Engine {
	r := gin.New()

	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	api := r.Group("/api/v1")
	{
		api.GET("/ping", h.Health.Ping)

		authGroup := api.Group("/auth")
		{
			authGroup.POST("/register", h.Auth.Register)
			authGroup.POST("/login", h.Auth.Login)
		}

		userGroup := api.Group("/users")
		userGroup.Use(middleware.Auth(jwtSecret))
		{
			userGroup.GET("/me", h.User.Me)
		}

	}

	return r
}
