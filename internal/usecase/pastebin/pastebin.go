package pastebin

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/alirezaarzehgar/pastebin/internal/domain/model"
	"github.com/alirezaarzehgar/pastebin/internal/domain/repo"
	"github.com/alirezaarzehgar/pastebin/internal/infra/objstorage"
	"github.com/alirezaarzehgar/pastebin/internal/infra/persistance"
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
		Content:   args.Content,
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

func (pb pastebin) GetPasteContent(args dto.GetPasteContentArgs) (*dto.GetPasteContentReply, error) {
	ctx := context.Background()

	reply, err := pb.persistence.GetPasteMetadata(ctx, model.GetPasteMetadataArgs{ID: args.PasteID})
	if err != nil {
		if errors.Is(err, persistance.RecordNotFound) {
			return nil, usecase.NotFound
		}
		return nil, fmt.Errorf("failed to get content metadata: %w", err)
	}

	contentReply := dto.GetPasteContentReply{
		FileMetadatas: dto.ConvertModelToFileMetadatas(reply.FileMetadatas),
		Content:       reply.Content,
	}

	return &contentReply, nil
}

func (pb pastebin) GetPasteFile(args dto.GetPasteFileArgs) (*dto.GetPasteFileReply, error) {
	ctx := context.Background()

	fileReply, err := pb.objStore.DownloadOject(ctx, model.GetPasteObjectArgs{
		PasteID:  args.PasteID,
		Filename: args.Filename,
	})
	if err != nil {
		if errors.Is(err, objstorage.NotFound) {
			return nil, usecase.NotFound
		}
		return nil, fmt.Errorf("failed to get file: %w", err)
	}

	reply := dto.GetPasteFileReply{
		File: fileReply.File,
	}
	return &reply, nil
}
