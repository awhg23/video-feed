package dto

import "time"

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required,max=50"`
}

type CreateCommentResponse struct {
	CommentID uint64 `json:"comment_id"`
}

type CommentUser struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
}

type CommentItem struct {
	ID        uint64      `json:"id"`
	Content   string      `json:"content"`
	User      CommentUser `json:"user"`
	CreatedAt time.Time   `json:"created_at"`
}

type CommentListResponse struct {
	List []CommentItem `json:"list"`
}
