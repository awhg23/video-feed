package model

import "time"

type Video struct {
	ID           uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	AuthorID     uint64    `gorm:"column:author_id;not null;index"`
	Title        string    `gorm:"column:title;type:varchar(128);not null"`
	Description  string    `gorm:"column:description;type:varchar(500)"`
	VideoURL     string    `gorm:"column:video_url;type:varchar(255);not null"`
	CoverURL     string    `gorm:"column:cover_url;type:varchar(255)"`
	LikeCount    uint64    `gorm:"column:like_count;not null;default:0"`
	CommentCount uint64    `gorm:"column:comment_count;not null;default:0"`
	Status       uint8     `gorm:"column:status;not null;default:1"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Video) TableName() string {
	return "videos"
}
