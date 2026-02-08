package dotenv

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alirezaarzehgar/pastebin/internal/config"
)

func loadMinIOConfig(cfg *config.Config) error {
	cfg.MinIO.Endpoint = os.Getenv(KeyObjectStorageMinIOEndpoint)
	if cfg.MinIO.Endpoint == "" {
		return fmt.Errorf("%s is empty", KeyObjectStorageMinIOEndpoint)
	}

	cfg.MinIO.AccessKeyID = os.Getenv(KeyObjectStorageMinIOAccessKeyID)
	if cfg.MinIO.AccessKeyID == "" {
		return fmt.Errorf("%s is empty", KeyObjectStorageMinIOAccessKeyID)
	}

	cfg.MinIO.SecretAccessKey = os.Getenv(KeyObjectStorageMinIOSecretAccessKey)
	if cfg.MinIO.SecretAccessKey == "" {
		return fmt.Errorf("%s is empty", KeyObjectStorageMinIOSecretAccessKey)
	}

	useSSL, err := strconv.ParseBool(os.Getenv(KeyObjectStorageMinIOUseSSL))
	if err != nil {
		return fmt.Errorf("invalid %s value", KeyObjectStorageMinIOUseSSL)
	}
	cfg.MinIO.UseSSL = useSSL

	return nil
}
