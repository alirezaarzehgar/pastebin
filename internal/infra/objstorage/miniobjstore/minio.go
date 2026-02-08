package miniobjstore

import (
	"github.com/alirezaarzehgar/pastebin/internal/config"
	"github.com/alirezaarzehgar/pastebin/internal/domain/repo"
	"github.com/minio/minio-go/v7"
)

type minIO struct {
	cfg    *config.Config
	client *minio.Client
}

func New(cfg *config.Config) repo.ObjectStorage {
	return &minIO{
		cfg: cfg,
	}
}
