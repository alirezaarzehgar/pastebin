package factory

import (
	"github.com/alirezaarzehgar/pastebin/internal/config"
	"github.com/alirezaarzehgar/pastebin/internal/domain/repo"
	"github.com/alirezaarzehgar/pastebin/internal/infra/cache/redis"
)

type Cache uint

const (
	CacheRedis Cache = iota
)

func New(c Cache, cfg *config.Config) (repo.Cache, error) {
	switch c {
	case CacheRedis:
		return redis.NewClient(cfg)
	default:
		panic("invalid cache backend selected")
	}
}
