package repository

import (
	"video-feed/internal/model"

	"gorm.io/gorm"
)

type FollowRepository struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) *FollowRepository {
	return &FollowRepository{db: db}
}

func (r *FollowRepository) Create(follow *model.Follow) error {
	return r.db.Create(follow).Error
}

func (r FollowRepository) Delete(userID, followUserID uint64) error {
	return r.db.Where("user_id = ? AND follow_user_id = ?", userID, followUserID).Delete(&model.Follow{}).Error
}

func (r *FollowRepository) Exists(userID, followUserID uint64) (bool, error) {
	var count int64
	err := r.db.Model(&model.Follow{}).Where("user_id = ? AND follow_user_id = ?", userID, followUserID).Count(&count).Error
	return count > 0, err
}

func (r *FollowRepository) ListFollowingUserIDs(userID uint64) ([]uint64, error) {
	var ids []uint64
	err := r.db.Model(&model.Follow{}).Where("user_id = ?", userID).Pluck("follow_user_id", &ids).Error
	return ids, err
}

func (r *FollowRepository) ListFollowerUserIDs(userID uint64) ([]uint64, error) {
	var ids []uint64
	err := r.db.Model(&model.Follow{}).Where("follow_user_id = ?", userID).Pluck("user_id", &ids).Error
	return ids, err
}
