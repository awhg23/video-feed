package repository

import (
	"video-feed/internal/model"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(comment *model.Comment) error {
	return r.db.Create(comment).Error
}

func (r *CommentRepository) ListByVideoID(videoID uint64) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := r.db.Where("video_id = ?", videoID).Order("id DESC").Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}
