package model

import "time"

type CreatePasteMetadataArgs struct {
	ID        string        `json:"_id"`
	Metadatas FileMetadatas `json:"metadata"`
	ExpiredAt time.Time     `json:"expired_at"`
}
