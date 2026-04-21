package router

import (
	"video-feed/internal/handler"
	"video-feed/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Health  *handler.HealthHandler
	Auth    *handler.AuthHandler
	User    *handler.UserHandler
	Video   *handler.VideoHandler
	Follow  *handler.FollowHandler
	Like    *handler.LikeHandler
	Comment *handler.CommentHandler
	Feed    *handler.FeedHandler
}

func NewRouter(h *Handlers, jwtSecret string) *gin.Engine {
	r := gin.New()

	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	api := r.Group("/api/v1")
	{
		api.GET("/ping", h.Health.Ping)

		authGroup := api.Group("/auth")
		{
			authGroup.POST("/register", h.Auth.Register)
			authGroup.POST("/login", h.Auth.Login)
		}

		userGroup := api.Group("/users")
		{
			userGroup.GET("/:id/videos", h.Video.ListUserVideos)
			userGroup.GET("/:id/following", h.Follow.ListFollowing)
			userGroup.GET("/:id/followers", h.Follow.ListFollowers)

			userGroup.Use(middleware.Auth(jwtSecret))
			userGroup.GET("/me", h.User.Me)
			userGroup.POST("/:id/follow", h.Follow.Follow)
			userGroup.DELETE("/:id/follow", h.Follow.Unfollow)
		}

		videoGroup := api.Group("/videos")
		{
			videoGroup.GET("/:id", h.Video.Detail)
			videoGroup.GET("/:id/comments", h.Comment.List)

			videoGroup.Use(middleware.Auth(jwtSecret))
			videoGroup.POST("", h.Video.Create)
			videoGroup.POST("/:id/like", h.Like.Like)
			videoGroup.DELETE("/:id/like", h.Like.Unlike)
			videoGroup.POST("/:id/comments", h.Comment.Create)
		}

		feedGroup := api.Group("/feed")
		feedGroup.Use(middleware.Auth(jwtSecret))
		{
			feedGroup.GET("/following", h.Feed.Following)
		}
	}

	return r
}
