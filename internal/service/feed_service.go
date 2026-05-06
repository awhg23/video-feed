package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"video-feed/internal/dto"
	"video-feed/internal/model"
	"video-feed/internal/repository"
)

var ErrInvalidCursor = errors.New("invalid cursor")

type FeedService struct {
	followRepo *repository.FollowRepository
	videoRepo  *repository.VideoRepository
	userRepo   *repository.UserRepository
}

func NewFeedService(followRepo *repository.FollowRepository, videoRepo *repository.VideoRepository, userRepo *repository.UserRepository) *FeedService {
	return &FeedService{
		followRepo: followRepo,
		videoRepo:  videoRepo,
		userRepo:   userRepo,
	}
}

func (s *FeedService) GetFollowingFeed(userID uint64, cursor string, limit int) (*dto.FeedResponse, error) {
	authorIDs, err := s.followRepo.ListFollowingUserIDs(userID)
	if err != nil {
		return nil, err
	}
	if len(authorIDs) == 0 {
		return &dto.FeedResponse{
			List:       []dto.FeedItem{},
			NextCursor: "",
			HasMore:    false,
		}, nil
	}

	var cursorTime *time.Time
	var cursorID *uint64

	if cursor != "" {
		t, id, err := parseFeedCursor(cursor)
		if err != nil {
			return nil, ErrInvalidCursor
		}
		cursorTime = &t
		cursorID = &id
	}

	videos, err := s.videoRepo.ListFeedByAuthorIDs(authorIDs, cursorTime, cursorID, limit+1)
	if err != nil {
		return nil, err
	}

	hasMore := false
	if len(videos) > limit {
		hasMore = true
		videos = videos[:limit]
	}

	// N+1 优化：
	// 原来是循环 videos，每条视频都调用 GetByID 查一次作者
	// 现在改成：先收集所有 author_id，再批量查询 users，最后用 map 组装 Feed
	feedAuthorIDs := make([]uint64, 0, len(videos))
	for _, video := range videos {
		feedAuthorIDs = append(feedAuthorIDs, video.AuthorID)
	}

	feedAuthorIDs = uniqueUint64(feedAuthorIDs)

	authors, err := s.userRepo.GetByIDs(feedAuthorIDs)
	if err != nil {
		return nil, err
	}

	authorMap := make(map[uint64]model.User, len(authors))
	for _, author := range authors {
		authorMap[author.ID] = author
	}

	list := make([]dto.FeedItem, 0, len(videos))
	for _, video := range videos {
		author, ok := authorMap[video.AuthorID]
		if !ok {
			// 理论上不该发生，除非视频作者被删了。
			// V1 先跳过，避免整个 Feed 失败。
			continue
		}

		list = append(list, dto.FeedItem{
			ID:           video.ID,
			Title:        video.Title,
			Description:  video.Description,
			VideoURL:     video.VideoURL,
			CoverURL:     video.CoverURL,
			LikeCount:    video.LikeCount,
			CommentCount: video.CommentCount,
			Author: dto.VideoAuthor{
				ID:       author.ID,
				Username: author.Username,
				Nickname: author.Nickname,
			},
			CreatedAt: video.CreatedAt,
		})
	}

	nextCursor := ""
	if len(videos) > 0 {
		last := videos[len(videos)-1]
		nextCursor = buildFeedCursor(last.CreatedAt, last.ID)
	}

	return &dto.FeedResponse{
		List:       list,
		NextCursor: nextCursor,
		HasMore:    hasMore,
	}, nil
}

func buildFeedCursor(t time.Time, id uint64) string {
	return fmt.Sprintf("%s_%d", t.Format(time.RFC3339Nano), id)
}

func parseFeedCursor(cursor string) (time.Time, uint64, error) {
	idx := strings.LastIndex(cursor, "_")
	if idx == -1 {
		return time.Time{}, 0, ErrInvalidCursor
	}

	timePart := cursor[:idx]
	idPart := cursor[idx+1:]

	t, err := time.Parse(time.RFC3339Nano, timePart)
	if err != nil {
		return time.Time{}, 0, ErrInvalidCursor
	}

	id, err := strconv.ParseUint(idPart, 10, 64)
	if err != nil {
		return time.Time{}, 0, ErrInvalidCursor
	}

	return t, id, nil
}
