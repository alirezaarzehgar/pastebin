package model

import (
	"io"
	"time"
)

type UploadedFile struct {
	Reader      io.ReadCloser
	Filename    string
	ContentType string
	Size        int64
}

type UploadedFiles []UploadedFile

type CreatePasteObjectArgs struct {
	Files  UploadedFiles
	Expiry time.Duration
}

type FileMetadata struct {
	Filename    string `json:"filename"`
	ContentType string `json:"content-type"`
	Size        int64  `json:"size"`
	Checksum    string `json:"checksum"`
}

type FileMetadatas []FileMetadata

type CreatePasteObjectReply struct {
	Metadatas FileMetadatas
	ID        string
}
