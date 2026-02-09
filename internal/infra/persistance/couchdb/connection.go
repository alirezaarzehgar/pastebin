package couchdb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (p CouchDB) url() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/", p.protocol, p.username, p.password, p.hostname, p.port)
}

func (p CouchDB) Connect() error {
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
