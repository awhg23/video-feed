package repository

import (
	"errors"

	"gorm.io/gorm"

	"video-feed/internal/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByID(id uint64) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := r.db.Model(&model.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

// IsNotFound 判断错误是否是数据库未找到记录导致的错误
func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// 4.17 增加按id批量查找
func (r *UserRepository) GetByIDs(ids []uint64) ([]model.User, error) {
	if len(ids) == 0 {
		return []model.User{}, nil
	}

	var users []model.User
	err := r.db.Where("id IN ?", ids).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
