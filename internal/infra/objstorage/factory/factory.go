package factory

import (
	"github.com/alirezaarzehgar/pastebin/internal/config"
	"github.com/alirezaarzehgar/pastebin/internal/domain/repo"
	"github.com/alirezaarzehgar/pastebin/internal/infra/objstorage/miniobjstore"
)

type ObjectStorage uint

const (
	ObjectStorageMinIO = iota
)

func New(objStore ObjectStorage, cfg *config.Config) (repo.ObjectStorage, error) {
	switch objStore {
	case ObjectStorageMinIO:
		return miniobjstore.New(cfg)
	default:
		panic("invalid object storage backend selected")
	}
}
