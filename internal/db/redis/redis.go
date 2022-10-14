package redis

import (
	"context"
	"fmt"
	"trfine/pkg/logging"

	"github.com/go-redis/redis/v9"
)

type CredentialRedis struct {
	Host   string
	Port   string
	Secret string
	Size   int
}

type redisClient struct {
	ctx  context.Context
	cred *CredentialRedis
	log  logging.Logger
}

func NewRedisClient(ctx context.Context, cred *CredentialRedis, log logging.Logger) *redisClient {
	return &redisClient{
		ctx:  ctx,
		cred: cred,
		log:  log,
	}
}

func (rc *redisClient) ConnectToRedis() (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", rc.cred.Host, rc.cred.Port)
	client := redis.NewClient(&redis.Options{
		MaxIdleConns: 3,
		Addr:         addr,
		Password:     "",
		DB:           0,
	})

	_, err := client.Ping(rc.ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// func (rc *redisClient) GetStore() (redisGin.Store, error) {
// 	store, err := redisGin.NewStore(
// 		rc.cred.Size,
// 		"tcp",
// 		fmt.Sprintf("%s:%s", rc.cred.Host, rc.cred.Port),
// 		"",
// 		[]byte(rc.cred.Secret),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return store, nil
// }
