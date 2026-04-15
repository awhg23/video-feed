package handler

import (
	"github.com/gin-gonic/gin"

	"video-feed/pkg/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Ping(c *gin.Context) {
	response.Success(c, gin.H{
		"message": "pong",
	})
}
