package repo

import (
	"context"

	"github.com/alirezaarzehgar/pastebin/internal/domain/model"
)

type ObjectStorage interface {
	UploadOjects(context.Context, model.CreatePasteObjectArgs) (*model.CreatePasteObjectReply, error)
	DownloadOject(context.Context, model.GetPasteObjectArgs) (*model.GetPasteObjectReply, error)
}
