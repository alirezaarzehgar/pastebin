package dto

import (
	"github.com/alirezaarzehgar/pastebin/internal/domain/model"
)

type FileMetadata struct {
	Filename    string `json:"filename"`
	ContentType string `json:"content-type"`
	Size        int64  `json:"size"`
	Checksum    string `json:"checksum,omitempty"`
}

type FileMetadatas []FileMetadata

func ConvertModelToFileMetadatas(modelFileMetadatas model.FileMetadatas) (fileMetadatas FileMetadatas) {
	for _, f := range modelFileMetadatas {
		fileMetadatas = append(fileMetadatas, FileMetadata{
			Filename:    f.Filename,
			ContentType: f.ContentType,
			Size:        f.Size,
			Checksum:    f.Checksum,
		})
	}
	return fileMetadatas
}

type GetPasteContentArgs struct {
	PasteID string
}

type GetPasteContentReply struct {
	Content       string        `json:"content"`
	FileMetadatas FileMetadatas `json:"metadatas"`
}

type GetPasteFileArgs struct {
	PasteID  string
	Filename string
}

type GetPasteFileReply struct {
}
