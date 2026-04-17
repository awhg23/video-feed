package model

import "time"

type Follow struct {
	ID           uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	UserID       uint64    `gorm:"column:user_id;not null"`
	FollowUserID uint64    `gorm:"column:follow_user_id;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (Follow) TableName() string {
	return "follows"
}
