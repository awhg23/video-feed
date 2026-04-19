package service

import (
	"errors"
	"strings"

	"video-feed/internal/dto"
	"video-feed/internal/model"
	"video-feed/internal/repository"

	"gorm.io/gorm"
)

var ErrInvalidCommentContent = errors.New("invalid comment content")

type CommentService struct {
	commentRepo *repository.CommentRepository
	videoRepo   *repository.VideoRepository
	userRepo    *repository.UserRepository
	db          *gorm.DB
}

func NewCommentService(db *gorm.DB, commentRepo *repository.CommentRepository, videoRepo *repository.VideoRepository, userRepo *repository.UserRepository) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		videoRepo:   videoRepo,
		userRepo:    userRepo,
		db:          db,
	}
}

func (s *CommentService) CreateComment(userID, videoID uint64, req *dto.CreateCommentRequest) (uint64, error) {
	_, err := s.videoRepo.GetByID(videoID)
	if err != nil {
		if repository.IsNotFound(err) {
			return 0, ErrVideoNotFound
		}
		return 0, err
	}

	content := strings.TrimSpace(req.Content)
	if content == "" {
		return 0, ErrInvalidCommentContent
	}

	var commentID uint64
	err = s.db.Transaction(func(tx *gorm.DB) error {
		commentRepo := repository.NewCommentRepository(tx)
		videoRepo := repository.NewVideoRepository(tx)

		comment := &model.Comment{
			UserID:  userID,
			VideoID: videoID,
			Content: content,
		}

		if err := commentRepo.Create(comment); err != nil {
			return err
		}

		if err := videoRepo.IncrCommentCount(videoID); err != nil {
			return err
		}

		commentID = comment.ID
		return nil
	})
	if err != nil {
		return 0, err
	}

	return commentID, nil
}

func (s *CommentService) ListComments(videoID uint64) (*dto.CommentListResponse, error) {
	_, err := s.videoRepo.GetByID(videoID)
	if err != nil {
		if repository.IsNotFound(err) {
			return nil, ErrVideoNotFound
		}
		return nil, err
	}

	comments, err := s.commentRepo.ListByVideoID(videoID)
	if err != nil {
		return nil, err
	}

	list := make([]dto.CommentItem, 0, len(comments))
	for _, comment := range comments {
		user, err := s.userRepo.GetByID(comment.UserID)
		if err != nil {
			return nil, err
		}

		list = append(list, dto.CommentItem{
			ID:      comment.ID,
			Content: comment.Content,
			User: dto.CommentUser{
				ID:        user.ID,
				Username:  user.Username,
				Nickname:  user.Nickname,
				AvatarURL: user.AvatarURL,
			},
			CreatedAt: comment.CreatedAt,
		})
	}

	return &dto.CommentListResponse{List: list}, nil
}
