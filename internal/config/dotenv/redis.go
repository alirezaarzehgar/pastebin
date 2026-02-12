package dotenv

import (
	"fmt"
	"os"

	"github.com/alirezaarzehgar/pastebin/internal/config"
)

func loadRedisConfig(cfg *config.Config) error {
	cfg.Redis.Address = os.Getenv(KeyCacheRedisAddress)
	if cfg.Redis.Address == "" {
		return fmt.Errorf("%s is empty", KeyCacheRedisAddress)
	}

	var err error
	cfg.Redis.Password = os.Getenv(KeyCacheRedisPassword)
	cfg.Redis.DefaultMetadataTTL, err = parseDuration(KeyCacheRedisMetadataTTL, 60*60*24)
	if err != nil {
		return fmt.Errorf("invalid ttl for redis metadata: %w", err)
	}

	return nil
}
