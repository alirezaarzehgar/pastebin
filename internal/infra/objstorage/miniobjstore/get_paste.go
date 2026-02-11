package miniobjstore

import (
	"context"
	"errors"
	"fmt"

	"github.com/alirezaarzehgar/pastebin/internal/domain/model"
	"github.com/alirezaarzehgar/pastebin/internal/infra/objstorage"
	"github.com/alirezaarzehgar/pastebin/internal/logger"
	"github.com/minio/minio-go/v7"
)

func (m *minIO) DownloadOject(ctx context.Context, args model.GetPasteObjectArgs) (*model.GetPasteObjectReply, error) {
	obj, err := m.client.GetObject(ctx, args.PasteID, args.Filename, minio.GetObjectOptions{})
	if err != nil {
		var minioErr minio.ErrorResponse
		if errors.As(err, &minioErr) && minioErr.Code == minio.NoSuchKey || minioErr.Code == minio.NoSuchBucket {
			return nil, objstorage.NotFound
		}
		logger.Error("failed to get object", "error", err, "bucket", args.PasteID, "filename", args.Filename)
		return nil, fmt.Errorf("failed to get object: %w", err)
	}

	reply := model.GetPasteObjectReply{
		File: obj,
	}

	return &reply, nil
}
