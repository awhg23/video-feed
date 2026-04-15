package dto

import "time"

type CreateVideoRequest struct {
	Title       string `json:"title" binding:"required,min=1,max=128"`
	Description string `json:"description" binding:"max=500"`
	VideoURL    string `json:"video_url" binding:"required,url,max=255"`
	CoverURL    string `json:"cover_url" binding:"omitempty,url,max=255"`
}

type CreateVideoResponse struct {
	VideoID uint64 `json:"video_id"`
}

type VideoAuthor struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

type VideoDetailResponse struct {
	ID           uint64      `json:"id"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	VideoURL     string      `json:"video_url"`
	CoverURL     string      `json:"cover_url"`
	LikeCount    uint64      `json:"like_count"`
	CommentCount uint64      `json:"comment_count"`
	Author       VideoAuthor `json:"author"`
	CreatedAt    time.Time   `json:"created_at"`
}

type UserVideoItem struct {
	ID           uint64      `json:"id"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	VideoURL     string      `json:"video_url"`
	CoverURL     string      `json:"cover_url"`
	LikeCount    uint64      `json:"like_count"`
	CommentCount uint64      `json:"comment_count"`
	Author       VideoAuthor `json:"author"`
	CreatedAt    time.Time   `json:"created_at"`
}

type UserVideoListResponse struct {
	List     []UserVideoItem `json:"list"`
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
	Total    int64           `json:"total"`
}
