package service

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func videoDetailCacheKey(videoID uint64) string {
	return fmt.Sprintf("video:detail:%d", videoID)
}

func deleteVideoDetailCache(redisClient *redis.Client, videoID uint64) {
	if redisClient == nil {
		return
	}

	_ = redisClient.Del(context.Background(), videoDetailCacheKey(videoID)).Err()
}
