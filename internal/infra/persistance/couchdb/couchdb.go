package couchdb

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func NewClient(cfg *config.Config, cache repo.Cache) (repo.Persistence, error) {
	protocol := "https"
	if cfg.CouchDB.Insecure {
		protocol = "http"
	}

	db := CouchDB{
		cache: cache,

		hostname: cfg.CouchDB.Hostname,
		port:     cfg.CouchDB.Port,
		username: cfg.CouchDB.Username,
		password: cfg.CouchDB.Password,
		protocol: protocol,

		metadataDBName: "metadata",
	}

	if err := db.connect(); err != nil {
		return nil, err
	}

	return &db, nil
}

func (p CouchDB) url() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/", p.protocol, p.username, p.password, p.hostname, p.port)
}

func (p CouchDB) connect() error {
	resp, err := http.Get(p.url() + "_up")
	if err != nil {
		return fmt.Errorf("request failed for couchdb: %w", err)
	}

	var req struct {
		Status string `json:"status"`
	}
	err = json.NewDecoder(resp.Body).Decode(&req)
	if err != nil {
		return fmt.Errorf("invalid couchdb response: %w", err)
	}

	if req.Status != "ok" {
		return fmt.Errorf("couchdb status is not ok. status: [%s]", req.Status)
	}

	return nil
}
