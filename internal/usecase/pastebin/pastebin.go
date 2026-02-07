package pastebin

import (
	"fmt"
	"io"

	"github.com/alirezaarzehgar/pastebin/internal/usecase"
	"github.com/alirezaarzehgar/pastebin/internal/usecase/dto"
)

type pastebin struct {
}

func New() usecase.Pastebin {
	return &pastebin{}
}

func (pb pastebin) CreatePaste(args dto.CreatePasteArgs) (*dto.CreatePasteResp, error) {
	fmt.Println("content:", args.Content)
	for _, f := range args.Files {
		fmt.Println(f.Filename, ":")
		data, _ := io.ReadAll(f.Reader)
		fmt.Println(string(data))
		f.Reader.Close()
	}

	paste := &dto.CreatePasteResp{
		PasteID: "xxxx-xxxx-xxxx-xxxx",
	}
	return paste, nil
}
