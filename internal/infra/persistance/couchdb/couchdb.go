package couchdb

import (
	"github.com/alirezaarzehgar/pastebin/internal/config"
	"github.com/alirezaarzehgar/pastebin/internal/domain/repo"
)

type CouchDB struct {
	cache repo.Cache

	hostname string
	port     int
	username string
	password string
	protocol string

	metadataDBName string
}

func New(cfg *config.Config, cache repo.Cache) repo.Persistence {
	protocol := "https"
	if cfg.CouchDB.Insecure {
		protocol = "http"
	}

	return &CouchDB{
		cache: cache,

		hostname: cfg.CouchDB.Hostname,
		port:     cfg.CouchDB.Port,
		username: cfg.CouchDB.Username,
		password: cfg.CouchDB.Password,
		protocol: protocol,

		metadataDBName: "metadata",
	}
}
