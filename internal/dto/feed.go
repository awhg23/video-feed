package dto

import "time"

type FeedItem struct {
	ID           uint64      `json:"id"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	VideoURL     string      `json:"video_url"`
	CoberURL     string      `json:"cover_url"`
	LikeCount    uint64      `json:"like_count"`
	CommentCount uint64      `json:"comment_count"`
	Author       VideoAuthor `json:"author"`
	CreatedAt    time.Time   `json:"created_at"`
}

type FeedResponse struct {
	List       []FeedItem `json:"list"`
	NextCursor string     `json:"next_cursor"`
	HasMore    bool       `json:"has_more"`
}
