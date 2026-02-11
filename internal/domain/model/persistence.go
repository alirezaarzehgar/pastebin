package model

import "time"

type CreatePasteMetadataArgs struct {
	ID        string        `json:"_id"`
	Content   string        `json:"content"`
	Metadatas FileMetadatas `json:"metadata"`
	ExpiredAt time.Time     `json:"expired_at"`
}

type GetPasteMetadataArgs struct {
	ID string `json:"_id"`
}

type GetPasteMetadataReply struct {
	FileMetadatas FileMetadatas
	Content       string
}
