package http

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ResponseErrorParseMaltiPartForm      = errors.New("unable to parse multipart form data")
	ResponseErrorCreatePasteEmptyRequest = errors.New("unable to process empty request")
)

type ResponseError struct {
	Message string `json:"message"`
}

func (h *HttpHandler) JSON(w http.ResponseWriter, code int, response any) error {
	jsonData, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.WriteHeader(code)
	_, err = w.Write(jsonData)
	return err
}
