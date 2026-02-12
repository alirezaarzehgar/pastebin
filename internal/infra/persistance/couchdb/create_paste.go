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
	"github.com/alirezaarzehgar/pastebin/internal/logger"
)

func (p CouchDB) SavePasteMetadata(ctx context.Context, args model.CreatePasteMetadataArgs) error {
	endpoint, err := url.JoinPath(p.url(), p.metadataDBName, args.ID)
	if err != nil {
		return fmt.Errorf("invalid url path: %w", err)
	}

	body, err := json.Marshal(args)
	if err != nil {
		return fmt.Errorf("failed to encode args: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("unable to create new request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("unable to send request: %w", err)
	}

	var saveResp struct {
		Ok  bool   `json:"ok"`
		ID  string `json:"id"`
		Rev string `json:"rev"`
	}
	err = json.NewDecoder(resp.Body).Decode(&saveResp)
	if err != nil {
		return fmt.Errorf("invalid response: %w", err)
	}

	if !saveResp.Ok {
		return fmt.Errorf("request was not successful")
	}

	key := cacheMetadataKey(args.ID)
	exp := min(time.Until(args.ExpiredAt), MaxMetadataCacheTTL)
	value := CachedPaste{
		Content:   args.Content,
		Metadatas: args.Metadatas,
	}
	err = p.cache.Set(ctx, key, value, exp)
	if err != nil {
		logger.Debug("failed to set metadata in cache", "error", err, "key", key)
	}

	return nil
}
