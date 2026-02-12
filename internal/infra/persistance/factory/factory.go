package factory

import (
	"github.com/alirezaarzehgar/pastebin/internal/config"
	"github.com/alirezaarzehgar/pastebin/internal/domain/repo"
	"github.com/alirezaarzehgar/pastebin/internal/infra/persistance/couchdb"
)

type Persistence uint

const (
	PersistenceCouchDB = iota
)

func New(persistence Persistence, cache repo.Cache, cfg *config.Config) repo.Persistence {
	switch persistence {
	case PersistenceCouchDB:
		return couchdb.New(cfg, cache)
	default:
		panic("invalid object storage backend selected")
	}
}
