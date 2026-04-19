package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"video-feed/internal/service"
	"video-feed/pkg/errorno"
	"video-feed/pkg/response"
)

type LikeHandler struct {
	likeService *service.LikeService
}

func NewLikeHandler(likeService *service.LikeService) *LikeHandler {
	return &LikeHandler{likeService: likeService}
}

func (h *LikeHandler) Like(c *gin.Context) {
	userIDValue, exists := c.Get("user_id")
	if !exists {
		response.Error(c, errorno.CodeUnauthorized, "unauthorized")
		return
	}
	userID, ok := userIDValue.(uint64)
	if !ok {
		response.Error(c, errorno.CodeUnauthorized, "invalid user ID")
		return
	}

	videoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, errorno.CodeInvalidParam, "invalid video id")
		return
	}

	err = h.likeService.Like(userID, videoID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrVideoNotFound):
			response.Error(c, errorno.CodeNotFound, "video not found")
			return
		case errors.Is(err, service.ErrAlreadyLiked):
			response.Error(c, 5001, "already liked")
			return
		default:
			response.Error(c, errorno.CodeInternalErr, "internal error")
			return
		}
	}

	response.Success(c, gin.H{"message": "like success"})
}

func (h *LikeHandler) Unlike(c *gin.Context) {
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

	videoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, errorno.CodeInvalidParam, "invalid video id")
		return
	}

	err = h.likeService.Unlike(userID, videoID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrVideoNotFound):
			response.Error(c, errorno.CodeNotFound, "video not found")
			return
		case errors.Is(err, service.ErrNotLiked):
			response.Error(c, 5002, "not liked")
			return
		default:
			response.Error(c, errorno.CodeInternalErr, "internal error")
			return
		}
	}

	response.Success(c, gin.H{"message": "unlike success"})
}
