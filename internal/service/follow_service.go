package service

import (
	"errors"

	"video-feed/internal/dto"
	"video-feed/internal/model"
	"video-feed/internal/repository"
)

var (
	ErrCannotFollowSelf = errors.New("cannot follow self")
	ErrAlreadyFollowed  = errors.New("already followed")
	ErrNotFollowed      = errors.New("not followed")
)

type FollowService struct {
	followRepo *repository.FollowRepository
	userRepo   *repository.UserRepository
}

func NewFollowService(followRepo *repository.FollowRepository, userRepo *repository.UserRepository) *FollowService {
	return &FollowService{
		followRepo: followRepo,
		userRepo:   userRepo,
	}
}

func (s *FollowService) Follow(userID, followUserID uint64) error {
	if userID == followUserID {
		return ErrCannotFollowSelf
	}

	_, err := s.userRepo.GetByID(followUserID)
	if err != nil {
		if repository.IsNotFound(err) {
			return ErrUserNotFound
		}
		return err
	}

	exists, err := s.followRepo.Exists(userID, followUserID)
	if err != nil {
		return err
	}
	if exists {
		return ErrAlreadyFollowed
	}

	follow := &model.Follow{
		UserID:       userID,
		FollowUserID: followUserID,
	}

	return s.followRepo.Create(follow)
}

func (s *FollowService) Unfollow(userID, followUserID uint64) error {
	if userID == followUserID {
		return ErrCannotFollowSelf
	}

	exists, err := s.followRepo.Exists(userID, followUserID)
	if err != nil {
		return err
	}
	if !exists {
		return ErrNotFollowed
	}

	return s.followRepo.Delete(userID, followUserID)
}

func (s *FollowService) ListFollowing(userID uint64) (*dto.FollowListResponse, error) {
	ids, err := s.followRepo.ListFollowingUserIDs(userID)
	if err != nil {
		return nil, err
	}

	users, err := s.userRepo.GetByIDs(ids)
	if err != nil {
		return nil, err
	}

	list := make([]dto.FollowUserItem, 0, len(ids))
	for _, user := range users {
		list = append(list, dto.FollowUserItem{
			ID:        user.ID,
			Username:  user.Username,
			Nickname:  user.Nickname,
			AvatarURL: user.AvatarURL,
			Bio:       user.Bio,
		})
	}

	return &dto.FollowListResponse{
		List: list,
	}, nil
}

func (s *FollowService) ListFollowers(userID uint64) (*dto.FollowListResponse, error) {
	ids, err := s.followRepo.ListFollowerUserIDs(userID)
	if err != nil {
		return nil, err
	}

	users, err := s.userRepo.GetByIDs(ids)
	if err != nil {
		return nil, err
	}

	list := make([]dto.FollowUserItem, 0, len(ids))
	for _, user := range users {
		list = append(list, dto.FollowUserItem{
			ID:        user.ID,
			Username:  user.Username,
			Nickname:  user.Nickname,
			AvatarURL: user.AvatarURL,
			Bio:       user.Bio,
		})
	}

	return &dto.FollowListResponse{
		List: list,
	}, nil
}
