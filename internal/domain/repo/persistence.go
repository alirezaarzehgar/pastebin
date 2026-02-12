package repo

import (
	"context"

	"github.com/alirezaarzehgar/pastebin/internal/domain/model"
)

type Persistence interface {
	SavePasteMetadata(ctx context.Context, args model.CreatePasteMetadataArgs) error
	GetPasteMetadata(ctx context.Context, args model.GetPasteMetadataArgs) (*model.GetPasteMetadataReply, error)
}
