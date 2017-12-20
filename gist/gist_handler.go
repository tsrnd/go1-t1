package gist

import (
	"net/http"

	"goweb1/shared/handler"
)

// HTTPGistHandler struct.
type HTTPGistHandler struct {
	handler.BaseHTTPHandler
	GUsecase UsecaseInterface
}

// ListGists call usecase to get git urls.
func (h *HTTPGistHandler) ListGists(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	id := queryValues.Get("id")
	if id == "" {
		// default value
		id = "DaisukeHirata"
	}

	list, err := h.GUsecase.ListGists(id)

	if err != nil || len(list) == 0 {
		list = append(list, "404 not found")
		h.StatusNotFoundRequest(w, list)
		return
	}

	h.ResponseJSON(w, list)
}

// NewGistHTTPHandler responses new HTTPGistHandler instance.
func NewGistHTTPHandler(bh *handler.BaseHTTPHandler) *HTTPGistHandler {
	gr := NewRepository()
	gu := NewUsecase(gr)
	return &HTTPGistHandler{
		BaseHTTPHandler: *bh,
		GUsecase:        gu,
	}
}
