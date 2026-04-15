package router

import (
	"video-feed/internal/handler"
	"video-feed/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Health *handler.HealthHandler
}

func NewRouter(h *Handlers) *gin.Engine {
	r := gin.New()

	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	api := r.Group("/api/v1")
	{
		api.GET("/ping", h.Health.Ping)
	}

	return r
}
