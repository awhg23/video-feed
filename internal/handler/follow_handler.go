package handler

import (
	"errors"
	"strconv"

	"video-feed/internal/service"
	"video-feed/pkg/errorno"
	"video-feed/pkg/response"

	"github.com/gin-gonic/gin"
)

type FollowHandler struct {
	followService *service.FollowService
}

func NewFollowHandler(followService *service.FollowService) *FollowHandler {
	return &FollowHandler{
		followService: followService,
	}
}

func (h *FollowHandler) Follow(c *gin.Context) {
	currentUserIDValue, exists := c.Get("user_id")
	if !exists {
		response.Error(c, errorno.CodeUnauthorized, "unauthorized")
		return
	}
	currentUserID, ok := currentUserIDValue.(uint64)
	if !ok {
		response.Error(c, errorno.CodeUnauthorized, "unauthorized")
		return
	}

	targetUserID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, errorno.CodeInvalidParam, "invalid user_id")
		return
	}

	err = h.followService.Follow(currentUserID, targetUserID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrCannotFollowSelf):
			response.Error(c, errorno.CodeInvalidParam, "cannot follow self")
			return
		case errors.Is(err, service.ErrAlreadyFollowed):
			response.Error(c, 4002, "already followed")
			return
		case errors.Is(err, service.ErrUserNotFound):
			response.Error(c, errorno.CodeNotFound, "user not found")
			return
		default:
			response.Error(c, errorno.CodeInternalErr, "internal error")
			return
		}
	}

	response.Success(c, gin.H{
		"message": "follow success",
	})
}

func (h *FollowHandler) Unfollow(c *gin.Context) {
	currentUserIDValue, exists := c.Get("user_id")
	if !exists {
		response.Error(c, errorno.CodeUnauthorized, "unauthorized")
		return
	}
	currentUserID, ok := currentUserIDValue.(uint64)
	if !ok {
		response.Error(c, errorno.CodeUnauthorized, "unauthorized")
		return
	}

	targetUserID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, errorno.CodeInvalidParam, "invalid user_id")
		return
	}

	err = h.followService.Unfollow(currentUserID, targetUserID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrCannotFollowSelf):
			response.Error(c, errorno.CodeInvalidParam, "cannot unfollow self")
			return
		case errors.Is(err, service.ErrNotFollowed):
			response.Error(c, 4003, "not followed")
			return
		default:
			response.Error(c, errorno.CodeInternalErr, "internal error")
			return
		}
	}

	response.Success(c, gin.H{
		"message": "unfollow success",
	})
}

func (h *FollowHandler) ListFollowing(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, errorno.CodeInvalidParam, "invalid user id")
		return
	}

	resp, err := h.followService.ListFollowing(userID)
	if err != nil {
		response.Error(c, errorno.CodeInternalErr, "internal error")
		return
	}

	response.Success(c, resp)
}

func (h *FollowHandler) ListFollowers(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, errorno.CodeInvalidParam, "invalid user id")
		return
	}

	resp, err := h.followService.ListFollowers(userID)
	if err != nil {
		response.Error(c, errorno.CodeInternalErr, "internal error")
		return
	}

	response.Success(c, resp)
}
