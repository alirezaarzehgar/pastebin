package miniobjstore

import (
	"errors"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func (m *minIO) Connect() error {
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
