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

	authService := service.NewAuthService(userRepo, cfg.JWT.Secret, cfg.JWT.ExpireHours)
	userService := service.NewUserService(userRepo)

	healthHandler := handler.NewHealthHandler()
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)

	engine := router.NewRouter(&router.Handlers{
		Health: healthHandler,
		Auth:   authHandler,
		User:   userHandler,
	}, cfg.JWT.Secret)

	return &App{
		Config: cfg,
		DB:     db,
		Logger: log,
		Engine: engine,
	}, nil
}
