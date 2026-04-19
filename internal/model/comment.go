package model

import "time"

type Comment struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	VideoID   uint64    `gorm:"column:video_id;not null;index"`
	UserID    uint64    `gorm:"column:user_id;not null;index"`
	Content   string    `gorm:"column:content;type:varchar(500);not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Comment) TableName() string {
	return "comments"
}
