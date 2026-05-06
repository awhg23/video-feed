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

	userIDs := make([]uint64, 0, len(comments))
	for _, comment := range comments {
		userIDs = append(userIDs, comment.UserID)
	}

	userIDs = uniqueUint64(userIDs)

	users, err := s.userRepo.GetByIDs(userIDs)
	if err != nil {
		return nil, err
	}

	userMap := make(map[uint64]model.User, len(users))
	for _, user := range users {
		userMap[user.ID] = user
	}

	list := make([]dto.CommentItem, 0, len(comments))
	for _, comment := range comments {
		user, ok := userMap[comment.UserID]
		if !ok {
			// 理论上不该发生，除非评论关联的用户被删了。
			// V1 先跳过，避免整个评论列表失败。
			continue
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
