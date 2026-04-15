package app

import (
	"fmt"
	"video-feed/internal/handler"
	"video-feed/internal/router"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go.uber.org/zap"
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
		return nil, fmt.Errorf("load config failed: %w", err)
	}

	gin.SetMode(cfg.App.Mode)

	log, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("init logger failed: %w", err)
	}

	db, err := NewDB(&cfg.MySQL)
	if err != nil {
		return nil, fmt.Errorf("init db failed: %w", err)
	}

	healthHandler := handler.NewHealthHandler()

	engine := router.NewRouter(&router.Handlers{
		Health: healthHandler,
	})

	return &App{
		Config: cfg,
		DB:     db,
		Logger: log,
		Engine: engine,
	}, nil
}
