package couchdb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/alirezaarzehgar/pastebin/internal/domain/model"
	"github.com/alirezaarzehgar/pastebin/internal/infra/persistance"
)

func (p CouchDB) GetPasteMetadata(ctx context.Context, args model.GetPasteMetadataArgs) (*model.GetPasteMetadataReply, error) {
	endpoint, err := url.JoinPath(p.url(), p.metadataDBName, args.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid url path: %w", err)
	}

	body, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("failed to encode args: %w", err)
	}

	req, err := http.NewRequest(http.MethodGet, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("unable to create new request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to send request: %w", err)
	}

	var metadataResp struct {
		Error     string              `json:"error"`
		Reason    string              `json:"reason"`
		ID        string              `json:"_id"`
		Rev       string              `json:"_rev"`
		Metadatas model.FileMetadatas `json:"metadata"`
		Content   string              `json:"content"`
		ExpiredAt time.Time           `json:"expired_at"`
	}
	err = json.NewDecoder(resp.Body).Decode(&metadataResp)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	if metadataResp.Error == "not_found" || time.Now().After(metadataResp.ExpiredAt) {
		return nil, persistance.RecordNotFound
	}

	if metadataResp.Error != "" {
		return nil, fmt.Errorf("request failed to get id %s cause %s", args.ID, metadataResp.Reason)
	}

	pasteMetadata := model.GetPasteMetadataReply{
		FileMetadatas: metadataResp.Metadatas,
		Content:       metadataResp.Content,
	}

	return &pasteMetadata, nil
}
