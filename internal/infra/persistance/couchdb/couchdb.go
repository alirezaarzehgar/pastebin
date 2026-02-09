package couchdb

import (
	"github.com/alirezaarzehgar/pastebin/internal/config"
	"github.com/alirezaarzehgar/pastebin/internal/domain/repo"
)

type CouchDB struct {
	hostname string
	port     int
	username string
	password string
	protocol string
}

func New(cfg *config.Config) repo.Persistence {
	protocol := "https"
	if cfg.CouchDB.Insecure {
		protocol = "http"
	}

	return &CouchDB{
		hostname: cfg.CouchDB.Hostname,
		port:     cfg.CouchDB.Port,
		username: cfg.CouchDB.Username,
		password: cfg.CouchDB.Password,
		protocol: protocol,
	}
}
