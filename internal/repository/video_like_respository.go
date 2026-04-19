package repository

import (
	"video-feed/internal/model"

	"gorm.io/gorm"
)

type VideoLikeRepository struct {
	db *gorm.DB
}

func NewVideoLikeRepository(db *gorm.DB) *VideoLikeRepository {
	return &VideoLikeRepository{db: db}
}

func (r *VideoLikeRepository) Create(like *model.VideoLike) error {
	return r.db.Create(like).Error
}

func (r *VideoLikeRepository) Delete(userID, videoID uint64) error {
	return r.db.Where("user_id = ? AND video_id = ?", userID, videoID).Delete(&model.VideoLike{}).Error
}

func (r *VideoLikeRepository) Exists(userID, videoID uint64) (bool, error) {
	var count int64
	err := r.db.Model(&model.VideoLike{}).Where("user_id = ? AND video_id = ?", userID, videoID).Count(&count).Error
	return count > 0, err
}
