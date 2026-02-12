package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/alirezaarzehgar/pastebin/internal/config"
	"github.com/alirezaarzehgar/pastebin/internal/domain/repo"
	"github.com/alirezaarzehgar/pastebin/internal/infra/cache"
	redisV9 "github.com/redis/go-redis/v9"
)

type RedisCache struct {
	defaultTTL time.Duration
	redis      *redisV9.Client
}

func New(cfg *config.Config) (repo.Cache, error) {
	r := redisV9.NewClient(&redisV9.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := r.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis connection failed: %w", err)
	}

	c := RedisCache{
		defaultTTL: cfg.Redis.DefaultMetadataTTL,
		redis:      r,
	}
	return &c, nil
}

func (rc RedisCache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	if ttl == 0 {
		ttl = rc.defaultTTL
	}

	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value: %w", err)
	}

	err = rc.redis.Set(ctx, key, data, ttl).Err()
	if err != redisV9.Nil {
		return fmt.Errorf("redis set failed: %w", err)
	}

	return nil
}

func (rc RedisCache) Get(ctx context.Context, key string, dest any) error {
	data, err := rc.redis.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redisV9.Nil) {
			return cache.NotFound
		}
		return fmt.Errorf("redis get failed: %w", err)
	}

	if err := json.Unmarshal(data, dest); err != nil {
		return err
	}

	return nil
}
