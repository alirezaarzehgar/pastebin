package dto

import (
	"io"
)

type UploadedFile struct {
	Reader      io.ReadCloser
	Filename    string
	ContentType string
	Size        int64
}

type CreatePasteArgs struct {
	Content string
	Files   []UploadedFile
}

type CreatePasteResp struct {
	PasteID string
}
