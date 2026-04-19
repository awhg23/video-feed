package model

import "time"

type VideoLike struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	UserID    uint64    `gorm:"column:user_id;not null"`
	VideoID   uint64    `gorm:"column:video_id;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (VideoLike) TableName() string {
	return "video_likes"
}
