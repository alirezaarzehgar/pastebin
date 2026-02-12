package http

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/alirezaarzehgar/pastebin/internal/logger"
	"github.com/alirezaarzehgar/pastebin/internal/usecase"
	"github.com/alirezaarzehgar/pastebin/internal/usecase/dto"
)

const (
	DefaultPasteExpiry = time.Duration(time.Hour * 24)
	MaxPasteExpiry     = time.Duration(time.Hour * 24 * 90)
)

func (h *HttpHandler) createPaste(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(h.config.Handler.MaxFileSize)
	if err != nil {
		h.JSON(w, http.StatusBadRequest, ResponseError{
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
			Message: ResponseErrorCreatePasteEmptyRequest.Error(),
		})
		return
	}

	expiry := DefaultPasteExpiry
	if r.FormValue("expiry") != "" {
		expiry, err = time.ParseDuration(r.FormValue("expiry"))
		if err != nil {
			h.JSON(w, http.StatusBadRequest, ResponseError{
				Message: ResponseErrorInvalidFields.Error(),
			})
			return
		}
	}

	if expiry > MaxPasteExpiry {
		h.JSON(w, http.StatusBadRequest, ResponseError{
			Message: ResponseErrorInvalidFields.Error(),
		})
		return
	}

	paste, err := h.pastebinUsecase.CreatePaste(dto.CreatePasteArgs{
		Content: content,
		Files:   files,
		Expiry:  expiry,
	})
	if err != nil {
		logger.Error("failed to create paste", "error", err)
		h.JSON(w, http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
		return
	}

	h.JSON(w, http.StatusOK, struct {
		PasteID string `json:"paste_id"`
	}{
		PasteID: paste.PasteID,
	})
}

func (h *HttpHandler) getPasteContent(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	reply, err := h.pastebinUsecase.GetPasteContent(dto.GetPasteContentArgs{PasteID: id})
	if err != nil {
		if errors.Is(err, usecase.NotFound) {
			h.JSON(w, http.StatusNotFound, ResponseError{
				Message: err.Error(),
			})
			return
		}

		logger.Error("failed to download given paste id", "error", err, "id", id)
		h.JSON(w, http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
		return
	}

	h.JSON(w, http.StatusOK, reply)
}

func (h *HttpHandler) getPasteFile(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	filename := r.PathValue("filename")

	fileReply, err := h.pastebinUsecase.GetPasteFile(dto.GetPasteFileArgs{
		PasteID:  id,
		Filename: filename,
	})
	if err != nil {
		logger.Error("failed to download given paste id", "error", err, "id", id)
		h.JSON(w, http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Transfer-Encoding", "binary")

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	_, err = io.Copy(w, fileReply.File)
	if err != nil {
		logger.Error("failed to stream file", "error", err, "id", id)
		return
	}
}
