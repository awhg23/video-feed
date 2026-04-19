package service

import (
	"errors"

	"video-feed/internal/model"
	"video-feed/internal/repository"

	"gorm.io/gorm"
)

var (
	ErrAlreadyLiked = errors.New("already liked")
	ErrNotLiked     = errors.New("not liked")
)

type LikeService struct {
	likeRepo  *repository.VideoLikeRepository
	videoRepo *repository.VideoRepository
	db        *gorm.DB
}

func NewLikeService(db *gorm.DB, likeRepo *repository.VideoLikeRepository, videoRepo *repository.VideoRepository) *LikeService {
	return &LikeService{
		likeRepo:  likeRepo,
		videoRepo: videoRepo,
		db:        db,
	}
}

func (s *LikeService) Like(userID, videoID uint64) error {
	_, err := s.videoRepo.GetByID(videoID)
	if err != nil {
		if repository.IsNotFound(err) {
			return ErrVideoNotFound
		}
		return err
	}

	exists, err := s.likeRepo.Exists(userID, videoID)
	if err != nil {
		return err
	}
	if exists {
		return ErrAlreadyLiked
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		likeRepo := repository.NewVideoLikeRepository(tx)
		videoRepo := repository.NewVideoRepository(tx)

		if err := likeRepo.Create(&model.VideoLike{
			UserID:  userID,
			VideoID: videoID,
		}); err != nil {
			return err
		}

		if err := videoRepo.IncrLikeCount(videoID); err != nil {
			return err
		}

		return nil
	})
}

func (s *LikeService) Unlike(userID, videoID uint64) error {
	_, err := s.videoRepo.GetByID(videoID)
	if err != nil {
		if repository.IsNotFound(err) {
			return ErrVideoNotFound
		}
		return err
	}

	exists, err := s.likeRepo.Exists(userID, videoID)
	if err != nil {
		return err
	}
	if !exists {
		return ErrNotLiked
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		likeRepo := repository.NewVideoLikeRepository(tx)
		videoRepo := repository.NewVideoRepository(tx)

		if err := likeRepo.Delete(userID, videoID); err != nil {
			return err
		}

		if err := videoRepo.DecrLikeCount(videoID); err != nil {
			return err
		}

		return nil
	})
}
