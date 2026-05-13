package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"video-feed/internal/service"
	"video-feed/pkg/errorno"
	"video-feed/pkg/response"
)

type FeedHandler struct {
	feedService *service.FeedService
}

func NewFeedHandler(feedService *service.FeedService) *FeedHandler {
	return &FeedHandler{feedService: feedService}
}

// Following 关注 Feed
// @Summary 获取关注 Feed
// @Description 获取当前登录用户关注作者发布的视频流，使用 cursor 分页
// @Tags Feed
// @Produce json
// @Security BearerAuth
// @Param cursor query string false "游标，第一页不传"
// @Param limit query int false "每页数量，默认 10，最大 50"
// @Success 200 {object} response.Resp
// @Router /feed/following [get]
func (h *FeedHandler) Following(c *gin.Context) {
	userIDValue, exists := c.Get("user_id")
	if !exists {
		response.Error(c, errorno.CodeUnauthorized, "Unauthorized")
		return
	}
	userID, ok := userIDValue.(uint64)
	if !ok {
		response.Error(c, errorno.CodeUnauthorized, "Unauthorized")
		return
	}

	limit := 10
	if limitStr := c.Query("limit"); limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err != nil || l <= 0 || l > 50 {
			response.Error(c, errorno.CodeInvalidParam, "Invalid limit")
			return
		}
		limit = l
	}

	cursor := c.Query("cursor")

	resp, err := h.feedService.GetFollowingFeed(userID, cursor, limit)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidCursor):
			response.Error(c, errorno.CodeInvalidParam, "invalid cursor")
			return
		default:
			response.Error(c, errorno.CodeInternalErr, "internal error")
			return
		}
	}

	response.Success(c, resp)
}
