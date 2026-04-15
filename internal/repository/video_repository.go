package repository

import (
	"video-feed/internal/model"

	"gorm.io/gorm"
)

type VideoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepository {
	return &VideoRepository{db: db}
}

func (r *VideoRepository) Create(video *model.Video) error {
	return r.db.Create(video).Error
}

func (r *VideoRepository) GetByID(id uint64) (*model.Video, error) {
	var video model.Video
	err := r.db.Where("id = ? AND status = ?", id, 1).First(&video).Error
	if err != nil {
		return nil, err
	}

	return &video, nil
}

func (r *VideoRepository) ListByAuthorID(authorID uint64, offset, limit int) ([]model.Video, int64, error) {
	var videos []model.Video
	var total int64

	if err := r.db.Model(&model.Video{}).
		Where("author_id = ? AND status = ?", authorID, 1).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Where("author_id = ? AND status = ?", authorID, 1).
		Order("created_at DESC, id DESC").
		Offset(offset).
		Limit(limit).
		Find(&videos).Error
	if err != nil {
		return nil, 0, err
	}

	return videos, total, nil
}
