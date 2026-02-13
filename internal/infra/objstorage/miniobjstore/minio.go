package miniobjstore

import (
	"errors"
	"fmt"

	"github.com/alirezaarzehgar/pastebin/internal/config"
	"github.com/alirezaarzehgar/pastebin/internal/domain/repo"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minIO struct {
	cfg    *config.Config
	client *minio.Client
}

func NewClient(cfg *config.Config) (repo.ObjectStorage, error) {
	objStore := minIO{
		cfg: cfg,
	}

	if err := objStore.connect(); err != nil {
		return nil, err
	}

	return &objStore, nil
}

func (m *minIO) connect() error {
	client, err := minio.New(m.cfg.MinIO.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.cfg.MinIO.AccessKeyID, m.cfg.MinIO.SecretAccessKey, ""),
		Secure: m.cfg.MinIO.UseSSL,
	})
	if err != nil {
		return fmt.Errorf("connecting to MinIO failed: %w", err)
	}

	if !client.IsOnline() {
		return errors.New("minio is not online")
	}

	m.client = client
	return nil
}
