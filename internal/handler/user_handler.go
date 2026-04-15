package handler

import (
	"video-feed/internal/service"
	"video-feed/pkg/errorno"
	"video-feed/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Me(c *gin.Context) {
	userIDValue, exists := c.Get("user_id")
	if !exists {
		response.Error(c, errorno.CodeUnauthorized, "unauthorized")
		return
	}

	userID, ok := userIDValue.(uint64)
	if !ok {
		response.Error(c, errorno.CodeUnauthorized, "unauthorized")
		return
	}

	profile, err := h.userService.GetProfile(userID)
	if err != nil {
		response.Error(c, errorno.CodeInternalErr, "internal error")
		return
	}

	response.Success(c, profile)
}
