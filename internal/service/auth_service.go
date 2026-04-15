package service

import (
	"errors"

	"video-feed/internal/dto"
	"video-feed/internal/model"
	"video-feed/internal/repository"
	jwtpkg "video-feed/pkg/jwt"
	"video-feed/pkg/password"
)

var (
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid username or password")
)

type AuthService struct {
	userRepo    *repository.UserRepository
	jwtSecret   string
	expireHours int
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string, expireHours int) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		jwtSecret:   jwtSecret,
		expireHours: expireHours,
	}
}

func (s *AuthService) Register(req *dto.RegisterRequest) (uint64, error) {
	exists, err := s.userRepo.ExistsByUsername(req.Username)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, ErrUserExists
	}

	hash, err := password.Hash(req.Password)
	if err != nil {
		return 0, err
	}

	user := &model.User{
		Username:     req.Username,
		PasswordHash: hash,
		Nickname:     req.Nickname,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (s *AuthService) Login(req *dto.LoginRequest) (string, error) {
	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		if repository.IsNotFound(err) {
			return "", ErrInvalidCredentials
		}
		return "", err
	}

	if !password.Verify(user.PasswordHash, req.Password) {
		return "", ErrInvalidCredentials
	}

	token, err := jwtpkg.GenerateToken(s.jwtSecret, user.ID, s.expireHours)
	if err != nil {
		return "", err
	}

	return token, nil

}
