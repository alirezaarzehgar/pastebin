package pastebin

import (
	"context"
	"fmt"
	"time"

	"github.com/alirezaarzehgar/pastebin/internal/domain/model"
	"github.com/alirezaarzehgar/pastebin/internal/domain/repo"
	"github.com/alirezaarzehgar/pastebin/internal/usecase"
	"github.com/alirezaarzehgar/pastebin/internal/usecase/dto"
)

type pastebin struct {
	persistence repo.Persistence
	objStore    repo.ObjectStorage
	cache       repo.Cache
}

func New(persistence repo.Persistence, objStore repo.ObjectStorage, cache repo.Cache) usecase.Pastebin {
	return &pastebin{
		persistence: persistence,
		objStore:    objStore,
		cache:       cache,
	}
}

func (pb pastebin) CreatePaste(args dto.CreatePasteArgs) (*dto.CreatePasteResp, error) {
	ctx := context.Background()
	expiredAt := time.Now().Add(args.Expiry)

	uploadedFiles := args.Files.ModelUploadedFile()
	reply, err := pb.objStore.UploadOjects(ctx, model.CreatePasteObjectArgs{
		Files:     uploadedFiles,
		ExpiredAt: expiredAt,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload files in object storage: %w", err)
	}

	err = pb.persistence.SavePasteMetadata(ctx, model.CreatePasteMetadataArgs{
		ID:        reply.ID,
		Metadatas: reply.Metadatas,
		ExpiredAt: expiredAt,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to store metadata in persistence backend: %w", err)
	}

	paste := &dto.CreatePasteResp{
		PasteID: reply.ID,
	}
	return paste, nil
}

func (pb pastebin) DownloadPaste(args dto.DownloadPasteArgs) (*dto.DownloadPasteReply, error) {
	return nil, nil
}
