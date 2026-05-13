package main

import (
	"fmt"
	"log"

	"video-feed/internal/app"
)

// @title video-feed API
// @version 1.0
// @description 短视频社区后端服务，支持用户注册登录、视频发布、关注、点赞、评论和关注 Feed。
// @termsOfService http://swagger.io/terms/

// @contact.name awhg23
// @contact.url https://github.com/awhg23/video-feed

// @host localhost:18080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description 输入格式：Bearer {token}
func main() {
	application, err := app.New()
	if err != nil {
		log.Fatalf("init app failed: %v", err)
	}

	addr := fmt.Sprintf(":%d", application.Config.App.Port)
	if err := application.Engine.Run(addr); err != nil {
		log.Fatalf("run server failed: %v", err)
	}
}
