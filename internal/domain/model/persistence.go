package model

type CreatePasteMetadataArgs struct {
	ID        string        `json:"_id"`
	Metadatas FileMetadatas `json:"metadata"`
}
