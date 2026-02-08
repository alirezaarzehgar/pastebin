package repo

import (
	"context"

	"github.com/alirezaarzehgar/pastebin/internal/domain/model"
)

type ObjectStorage interface {
	Connect() error
	UploadOjects(context.Context, model.CreatePasteObjectArgs) (*model.CreatePasteObjectReply, error)
}
