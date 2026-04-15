package service

import (
	"errors"

	"video-feed/internal/dto"
	"video-feed/internal/model"
	"video-feed/internal/repository"
)

var ErrVideoNotFound = errors.New("video not found")

type VideoService struct {
	videoRepo *repository.VideoRepository
	userRepo  *repository.UserRepository
}

func NewVideoService(videoRepo *repository.VideoRepository, userRepo *repository.UserRepository) *VideoService {
	return &VideoService{
		videoRepo: videoRepo,
		userRepo:  userRepo,
	}
}

func (s *VideoService) CreateVideo(userID uint64, req *dto.CreateVideoRequest) (uint64, error) {
	video := &model.Video{
		AuthorID:     userID,
		Title:        req.Title,
		Description:  req.Description,
		VideoURL:     req.VideoURL,
		CoverURL:     req.CoverURL,
		LikeCount:    0,
		CommentCount: 0,
		Status:       1,
	}

	if err := s.videoRepo.Create(video); err != nil {
		return 0, err
	}

	return video.ID, nil
}

func (s *VideoService) GetVideoDetail(videoID uint64) (*dto.VideoDetailResponse, error) {
	video, err := s.videoRepo.GetByID(videoID)
	if err != nil {
		if repository.IsNotFound(err) {
			return nil, ErrVideoNotFound
		}
		return nil, err
	}

	author, err := s.userRepo.GetByID(video.AuthorID)
	if err != nil {
		return nil, err
	}

	return &dto.VideoDetailResponse{
		ID:           video.ID,
		Title:        video.Title,
		Description:  video.Description,
		VideoURL:     video.VideoURL,
		CoverURL:     video.CoverURL,
		LikeCount:    video.LikeCount,
		CommentCount: video.CommentCount,
		Author: dto.VideoAuthor{
			ID:       author.ID,
			Username: author.Username,
			Nickname: author.Nickname,
		},
		CreatedAt: video.CreatedAt,
	}, nil
}

func (s *VideoService) ListUserVideos(userID uint64, page, pageSize int) (*dto.UserVideoListResponse, error) {
	offset := (page - 1) * pageSize
	videos, total, err := s.videoRepo.ListByAuthorID(userID, offset, pageSize)
	if err != nil {
		return nil, err
	}

	author, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	list := make([]dto.UserVideoItem, 0, len(videos))
	for _, video := range videos {
		list = append(list, dto.UserVideoItem{
			ID:           video.ID,
			Title:        video.Title,
			Description:  video.Description,
			VideoURL:     video.VideoURL,
			CoverURL:     video.CoverURL,
			LikeCount:    video.LikeCount,
			CommentCount: video.CommentCount,
			Author: dto.VideoAuthor{
				ID:       author.ID,
				Username: author.Username,
				Nickname: author.Nickname,
			},
			CreatedAt: video.CreatedAt,
		})
	}

	return &dto.UserVideoListResponse{
		List:     list,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}
