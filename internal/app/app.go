package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"video-feed/internal/handler"
	"video-feed/internal/repository"
	"video-feed/internal/router"
	"video-feed/internal/service"
)

type App struct {
	Config *Config
	DB     *gorm.DB
	Logger *zap.Logger
	Engine *gin.Engine
}

func New() (*App, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	gin.SetMode(cfg.App.Mode)

	log, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("init logger failed: %w", err)
	}

	db, err := NewDB(&cfg.MySQL)
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepository(db)
	videoRepo := repository.NewVideoRepository(db)
	followRepo := repository.NewFollowRepository(db)

	authService := service.NewAuthService(userRepo, cfg.JWT.Secret, cfg.JWT.ExpireHours)
	userService := service.NewUserService(userRepo)
	videoService := service.NewVideoService(videoRepo, userRepo)
	followService := service.NewFollowService(followRepo, userRepo)

	healthHandler := handler.NewHealthHandler()
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	videoHandler := handler.NewVideoHandler(videoService)
	followHandler := handler.NewFollowHandler(followService)

	engine := router.NewRouter(&router.Handlers{
		Health: healthHandler,
		Auth:   authHandler,
		User:   userHandler,
		Video:  videoHandler,
		Follow: followHandler,
	}, cfg.JWT.Secret)

	return &App{
		Config: cfg,
		DB:     db,
		Logger: log,
		Engine: engine,
	}, nil
}
