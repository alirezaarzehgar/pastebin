package repo

import (
	"context"

	"github.com/alirezaarzehgar/pastebin/internal/domain/model"
)

type Persistence interface {
	Connect() error
	SavePasteMetadata(ctx context.Context, args model.CreatePasteMetadataArgs) error
}
