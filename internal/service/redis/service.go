package service_rd

import (
	"context"

	"github.com/go-redis/redis/v9"
)

type RedisService interface {
}

type redisService struct {
	ctx context.Context
	rc  *redis.Client
}

func NewRedisService(ctx context.Context, rc *redis.Client) RedisService {
	return &redisService{
		ctx: ctx,
		rc:  rc,
	}
}
