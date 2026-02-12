package couchdb

import (
	"fmt"
	"time"

	"github.com/alirezaarzehgar/pastebin/internal/domain/model"
)

var (
	MaxMetadataCacheTTL = time.Hour
)

type CachedPaste struct {
	Content   string              `json:"content"`
	Metadatas model.FileMetadatas `json:"metadatas"`
}

func cacheMetadataKey(id string) string {
	return fmt.Sprintf("paste:metadata:%s", id)
}
