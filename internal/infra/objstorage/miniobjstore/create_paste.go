package miniobjstore

import (
	"context"
	"fmt"
	"time"

	"github.com/alirezaarzehgar/pastebin/internal/domain/model"
	"github.com/alirezaarzehgar/pastebin/internal/logger"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func (m *minIO) UploadOjects(ctx context.Context, args model.CreatePasteObjectArgs) (*model.CreatePasteObjectReply, error) {
	bucketName := uuid.New().String()

	err := m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return nil, fmt.Errorf("unable to create bucket: %w", err)
	}

	var metadatas model.FileMetadatas
	for _, f := range args.Files {
		uInfo, err := m.client.PutObject(ctx, bucketName, f.Filename, f.Reader, f.Size, minio.PutObjectOptions{
			Expires: time.Now().Add(args.Expiry),
		})
		if err != nil {
			logger.Error("failed to upload file on minio", "error", err, "filename", f.Filename, "bucket_name", bucketName)
		}
		metadatas = append(metadatas, model.FileMetadata{
			Filename:    f.Filename,
			ContentType: f.ContentType,
			Size:        f.Size,
			Checksum:    uInfo.ChecksumSHA1,
		})
	}

	reply := model.CreatePasteObjectReply{
		Metadatas: metadatas,
		ID:        bucketName,
	}
	return &reply, nil
}
