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

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

func (h *CommentHandler) Create(c *gin.Context) {
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

	var req dto.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorno.CodeInvalidParam, "invalid param")
		return
	}

	commentID, err := h.commentService.CreateComment(userID, videoID, &req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrVideoNotFound):
			response.Error(c, errorno.CodeNotFound, "video not found")
			return
		case errors.Is(err, service.ErrInvalidCommentContent):
			response.Error(c, errorno.CodeInvalidParam, "invalid comment content")
			return
		default:
			response.Error(c, errorno.CodeInternalErr, "internal error")
			return
		}
	}

	response.Success(c, dto.CreateCommentResponse{
		CommentID: commentID,
	})
}

func (h *CommentHandler) List(c *gin.Context) {
	videoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, errorno.CodeInvalidParam, "invalid video id")
		return
	}

	resp, err := h.commentService.ListComments(videoID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrVideoNotFound):
			response.Error(c, errorno.CodeNotFound, "video not found")
			return
		default:
			response.Error(c, errorno.CodeInternalErr, "internal error")
			return
		}
	}

	response.Success(c, resp)
}
