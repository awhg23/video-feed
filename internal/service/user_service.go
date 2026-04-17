package service

import (
	"errors"
	"video-feed/internal/dto"
	"video-feed/internal/repository"
)

var ErrUserNotFound = errors.New("user not found")

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetProfile(userID uint64) (*dto.UserProfileResponse, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		if repository.IsNotFound(err) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &dto.UserProfileResponse{
		ID:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		AvatarURL: user.AvatarURL,
		Bio:       user.Bio,
	}, nil
}
