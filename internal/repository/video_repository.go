package repository

import (
	"time"
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

func (r *VideoRepository) IncrLikeCount(videoID uint64) error {
	return r.db.Model(&model.Video{}).
		Where("id = ? AND status = ?", videoID, 1).
		UpdateColumn("like_count", gorm.Expr("like_count + 1")).Error
}

func (r *VideoRepository) DecrLikeCount(videoID uint64) error {
	return r.db.Model(&model.Video{}).
		Where("id = ? AND status = ?", videoID, 1).
		UpdateColumn("like_count", gorm.Expr("like_count - 1")).Error
}

func (r *VideoRepository) IncrCommentCount(videoID uint64) error {
	return r.db.Model(&model.Video{}).
		Where("id = ? AND status = ?", videoID, 1).
		UpdateColumn("comment_count", gorm.Expr("comment_count + 1")).Error
}

func (r *VideoRepository) DecrCommentCount(videoID uint64) error {
	return r.db.Model(&model.Video{}).
		Where("id = ? AND status = ?", videoID, 1).
		UpdateColumn("comment_count", gorm.Expr("comment_count - 1")).Error
}

func (r *VideoRepository) ListFeedByAuthorIDs(authorIDs []uint64, cursorTime *time.Time, cursorID *uint64, limit int) ([]*model.Video, error) {
	if len(authorIDs) == 0 {
		return []*model.Video{}, nil
	}

	db := r.db.Where("author_id IN ? AND status = ?", authorIDs, 1)

	if cursorTime != nil && cursorID != nil {
		db = db.Where(
			"(created_at < ?) OR (created_at = ? AND id < ?)",
			*cursorTime, *cursorTime, *cursorID,
		)
	}

	var videos []*model.Video
	err := db.Order("created_at DESC, id DESC").
		Limit(limit).
		Find(&videos).Error
	if err != nil {
		return nil, err
	}

	return videos, nil
}
