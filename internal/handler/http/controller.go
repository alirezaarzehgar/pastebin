package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alirezaarzehgar/pastebin/internal/usecase/dto"
)

const (
	DefaultPasteExpiry = time.Duration(time.Hour * 24)
)

func (h *HttpHandler) createPaste(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(h.config.Handler.MaxFileSize)
	if err != nil {
		h.JSON(w, http.StatusBadRequest, ResponseError{
			Error:   err,
			Message: ResponseErrorParseMaltiPartForm.Error(),
		})
		return
	}
	var files []dto.UploadedFile

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, err := fileHeader.Open()
			if err != nil {
				continue
			}

			files = append(files, dto.UploadedFile{
				Reader:      file,
				Size:        fileHeader.Size,
				Filename:    fileHeader.Filename,
				ContentType: fileHeader.Header.Get("Content-Type"),
			})
		}
	}

	content := r.FormValue("content")
	if content == "" && len(files) == 0 {
		h.JSON(w, http.StatusBadRequest, ResponseError{
			Error:   fmt.Errorf("empty request"),
			Message: ResponseErrorCreatePasteEmptyRequest.Error(),
		})
		return
	}

	paste, err := h.pastebinUsecase.CreatePaste(dto.CreatePasteArgs{
		Content: content,
		Files:   files,
		Expiry:  DefaultPasteExpiry,
	})
	if err != nil {
		h.JSON(w, http.StatusBadRequest, ResponseError{
			Error:   err,
			Message: "",
		})
		return
	}

	h.JSON(w, http.StatusOK, struct {
		PasteID string `json:"paste_id"`
	}{
		PasteID: paste.PasteID,
	})
}

func (h *HttpHandler) getPaste(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET PASTE"))
}
