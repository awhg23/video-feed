package model

import "time"

type User struct {
	ID           uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	Username     string    `gorm:"column:username;type:varchar(32);not null;uniqueIndex"`
	PasswordHash string    `gorm:"column:password_hash;type:varchar(255);not null"`
	Nickname     string    `gorm:"column:nickname;type:varchar(64);not null"`
	AvatarURL    string    `gorm:"column:avatar_url;type:varchar(255)"`
	Bio          string    `gorm:"column:bio;type:varchar(255)"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}
