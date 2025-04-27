package helper

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
)

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func (response *Response) ResponseWriter(w http.ResponseWriter, r *http.Request, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	raw, err := json.Marshal(response)
	if err != nil {
		slog.Error(err.Error())
		w.Write([]byte("something went wrong"))
	}
	w.Write(raw)
}

func (response *Response) IsValidMediaType(w http.ResponseWriter, r *http.Request) bool {
	contentType := strings.TrimSpace(r.Header.Get("Content-Type"))
	if contentType == "" || contentType != "application/json" {
		response.Message = "Not valid content-type"
		response.ResponseWriter(w, r, http.StatusUnsupportedMediaType)
		return true
	}
	return false
}
