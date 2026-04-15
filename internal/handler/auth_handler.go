package handler

import (
	"errors"
	"video-feed/internal/dto"
	"video-feed/internal/service"
	"video-feed/pkg/errorno"
	"video-feed/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorno.CodeInvalidParam, "invalid params")
		return
	}

	userID, err := h.authService.Register(&req)
	if err != nil {

		if errors.Is(err, service.ErrUserExists) {
			response.Error(c, errorno.CodeUserExists, "user already exists")
			return
		}
		response.Error(c, errorno.CodeInternalErr, "internal error")
		return
	}

	response.Success(c, dto.RegisterResponse{
		UserID: userID,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorno.CodeInvalidParam, "invalid params")
		return
	}

	token, err := h.authService.Login(&req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			response.Error(c, errorno.CodeInvalidPassword, "invalid username or password")
			return
		}
		response.Error(c, errorno.CodeInternalErr, "internal error")
		return
	}

	response.Success(c, dto.LoginResponse{
		Token: token,
	})
}
