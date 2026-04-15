# Video-feed
## 项目介绍
video-feed 是一个面向短视频社区的 Go 后端项目，核心支持用户注册登录、视频发布、关注关系、点赞评论和关注 Feed 流拉取。项目采用单体架构，基于 Gin、GORM、MySQL 和 JWT 实现。Feed 流第一版使用拉模式和 cursor 分页，优先保证业务闭环和设计清晰，后续再通过 Redis 和异步化做性能优化
## 后端
Go
Gin
## 数据库
MySQL
## 缓存（未实现）
Redis
## ORM
GORM
## 鉴权
JWT
## 配置
Viper
## 日志
zap / slog 都可以
## 部署（未实现）
Docker + docker-compose
```
video-feed
├─ README.md
├─ VERSION
├─ cmd
│  └─ server
│     └─ main.go
├─ config
│  └─ config.yaml
├─ docs
├─ go.mod
├─ go.sum
├─ internal
│  ├─ app
│  │  ├─ app.go
│  │  ├─ config.go
│  │  ├─ db.go
│  │  └─ load.go
│  ├─ dto
│  ├─ handler
│  │  └─ health_handler.go
│  ├─ middleware
│  │  ├─ auth.go
│  │  ├─ logger.go
│  │  └─ recovery.go
│  ├─ model
│  ├─ repository
│  ├─ router
│  │  └─ router.go
│  └─ service
├─ migrations
├─ pkg
│  ├─ errno
│  │  └─ errorno.go
│  ├─ jwt
│  ├─ logger
│  │  └─ logger.go
│  ├─ password
│  └─ response
│     └─ response.go
└─ scripts

```