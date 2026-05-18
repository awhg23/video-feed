# video-feed

## 项目简介

本项目是一个面向短视频社区场景的 Go 后端项目，核心支持用户注册登录、视频发布、关注关系、点赞评论、关注 Feed 流拉取等功能。

项目采用单体架构，基于 Gin + GORM + MySQL + Redis + JWT 实现。  
第一版 Feed 流采用拉模式，并使用 cursor 分页解决时间流场景下 offset 分页容易重复、漏读的问题。后续在项目中补充了 N+1 查询优化、Redis 视频详情缓存、缓存一致性处理、Docker Compose 环境编排和 Swagger API 文档。

---

## 技术栈

| 类型 | 技术 |
|---|---|
| 编程语言 | Go |
| Web 框架 | Gin |
| ORM | GORM |
| 数据库 | MySQL |
| 缓存 | Redis |
| 鉴权 | JWT |
| 配置管理 | Viper |
| API 文档 | Swagger / swaggo |
| 容器化 | Docker / Docker Compose |
| 接口测试 | curl 脚本 |

---

## 核心功能

### 用户模块
- 用户注册
- 用户登录
- JWT 鉴权
- 获取当前登录用户信息

### 视频模块
- 发布视频
- 获取视频详情
- 获取用户发布的视频列表
- 视频详情 Redis 缓存

### 关注模块
- 关注用户
- 取消关注
- 获取关注列表
- 获取粉丝列表

### 点赞模块
- 点赞视频
- 取消点赞
- 维护视频点赞数
- 点赞后删除视频详情缓存，保证缓存最终一致

### 评论模块
- 发表评论
- 获取评论列表
- 维护视频评论数
- 评论后删除视频详情缓存，保证缓存最终一致

### Feed 模块
- 获取关注 Feed
- 基于关注关系拉取作者视频
- 按 `created_at desc, id desc` 排序
- 使用 cursor 分页，避免 offset 分页在动态时间流中的重复和漏读问题

---

## 项目结构

```text
video-feed
├── cmd
│   └── server              # 程序入口
├── config                  # 配置文件
├── docs                    # Swagger 生成文档
├── internal
│   ├── app                 # 应用初始化、配置加载、DB/Redis 初始化
│   ├── dto                 # 请求和响应结构体
│   ├── handler             # HTTP 接口层
│   ├── middleware          # JWT、日志、Recovery 中间件
│   ├── model               # 数据库模型
│   ├── repository          # 数据访问层
│   ├── router              # 路由注册
│   └── service             # 业务逻辑层
├── migrations              # 数据库初始化 SQL
├── pkg
│   ├── errorno             # 业务错误码
│   ├── jwt                 # JWT 工具
│   ├── password            # 密码哈希工具
│   └── response            # 统一响应
├── scripts                 # 接口测试脚本
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

## 本地配置

项目不会提交真实配置文件，请先复制示例配置：

```bash
cp config/config.example.yaml config/config.yaml
cp .env.example .env
```