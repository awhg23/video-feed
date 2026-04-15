package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"video-feed/internal/dto"
	"video-feed/internal/service"
	"video-feed/pkg/errorno"
	"video-feed/pkg/response"
)

type VideoHandler struct {
	videoService *service.VideoService
}

func NewVideoHandler(videoService *service.VideoService) *VideoHandler {
	return &VideoHandler{videoService: videoService}
}

func (h *VideoHandler) Create(c *gin.Context) {
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

	var req dto.CreateVideoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorno.CodeInvalidParam, "invalid params")
		return
	}

	videoID, err := h.videoService.CreateVideo(userID, &req)
	if err != nil {
		response.Error(c, errorno.CodeInternalErr, "internal error")
		return
	}

	response.Success(c, dto.CreateVideoResponse{
		VideoID: videoID,
	})
}

func (h *VideoHandler) Detail(c *gin.Context) {
	videoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, errorno.CodeInvalidParam, "invalid video id")
		return
	}

	detail, err := h.videoService.GetVideoDetail(uint64(videoID))
	if err != nil {
		if errors.Is(err, service.ErrVideoNotFound) {
			response.Error(c, errorno.CodeNotFound, "video not found")
			return
		}
		response.Error(c, errorno.CodeInternalErr, "internal error")
		return
	}

	response.Success(c, detail)
}

func (h *VideoHandler) ListUserVideos(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, errorno.CodeInvalidParam, "invalid user id")
		return
	}

	page, pageSize := 1, 10

	if pageStr := c.Query("page"); pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err != nil || p <= 0 {
			response.Error(c, errorno.CodeInvalidParam, "invalid page")
			return
		}
		page = p
	}

	if pageSizeStr := c.Query("page_size"); pageSizeStr != "" {
		ps, err := strconv.Atoi(pageSizeStr)
		if err != nil || ps <= 0 || ps > 100 {
			response.Error(c, errorno.CodeInvalidParam, "invalid page_size")
			return
		}
		pageSize = ps
	}

	resp, err := h.videoService.ListUserVideos(userID, page, pageSize)
	if err != nil {
		response.Error(c, errorno.CodeInternalErr, "internal error")
		return
	}

	response.Success(c, resp)
}
