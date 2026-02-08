package dto

import (
	"io"
	"time"

	"github.com/alirezaarzehgar/pastebin/internal/domain/model"
)

type UploadedFile struct {
	Reader      io.ReadCloser
	Filename    string
	ContentType string
	Size        int64
}

type UploadedFiles []UploadedFile

func (uf UploadedFiles) ModelUploadedFile() (uploadedFiles model.UploadedFiles) {
	for _, f := range uf {
		uploadedFiles = append(uploadedFiles, model.UploadedFile{
			Reader:      f.Reader,
			Filename:    f.Filename,
			ContentType: f.ContentType,
			Size:        f.Size,
		})
	}
	return uploadedFiles
}

type CreatePasteArgs struct {
	Content string
	Files   UploadedFiles
	Expiry  time.Duration
}

type CreatePasteResp struct {
	PasteID string
}
