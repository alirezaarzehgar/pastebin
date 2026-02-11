package usecase

import "github.com/alirezaarzehgar/pastebin/internal/usecase/dto"

type Pastebin interface {
	CreatePaste(args dto.CreatePasteArgs) (*dto.CreatePasteResp, error)
	GetPasteContent(args dto.GetPasteContentArgs) (*dto.GetPasteContentReply, error)
	GetPasteFile(args dto.GetPasteFileArgs) (*dto.GetPasteFileReply, error)
}
