package handler

import (
	"crypto/rand"
	"encoding/json"
	"net/http"
	"os"

	"encoding/hex"
	"path/filepath"

	"io"

	"github.com/go-playground/form"
	"goweb1/infrastructure"
	"github.com/nicksnyder/go-i18n/i18n/bundle"
	"github.com/sirupsen/logrus"
)

const (
	// FilePerm is file permission.
	FilePerm = 0666
)

// BaseHTTPHandler base handler struct.
type BaseHTTPHandler struct {
	Logger *logrus.Logger
}

// Parse parse  form struct.
// https://github.com/go-playground/form
func (h *BaseHTTPHandler) Parse(r *http.Request, i interface{}) {
	// mapping post to struct.
	r.ParseForm()
	decoder := form.NewDecoder()
	decoder.Decode(&i, r.PostForm)
}

// ParseMultipart parse form struct.
// https://github.com/go-playground/form
func (h *BaseHTTPHandler) ParseMultipart(r *http.Request, i interface{}) {
	maxMemory := infrastructure.GetConfigInt64("multipart.maxmemory")
	r.ParseMultipartForm(maxMemory)
	// mapping post to struct.
	decoder := form.NewDecoder()
	decoder.Decode(&i, r.PostForm)
}

// SaveToFile saves uploaded file to new path.
// it only operates the first one of mutil-upload form file field.
func (h *BaseHTTPHandler) SaveToFile(r *http.Request, fromfile, tofile string) error {
	file, _, err := r.FormFile(fromfile)
	if err != nil {
		return err
	}
	defer file.Close()
	f, err := os.OpenFile(tofile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, FilePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(f, file)
	return nil
}

// GetRandomTempFileName get temp random filename.
func (h *BaseHTTPHandler) GetRandomTempFileName(prefix, filename string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	extension := filepath.Ext(filename)
	return filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes)+extension)
}

// ResponseJSON responses status code 200 and json.
func (h *BaseHTTPHandler) ResponseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(data)
	w.Write(json)
}

// StatusRedirect responses status code 302 and redirect.
func (h *BaseHTTPHandler) StatusRedirect(w http.ResponseWriter, url string) {
	w.Header().Set("Location", url)
	// status code 302
	w.WriteHeader(http.StatusFound)
}

// StatusBadRequest responses status code 400 and json.
func (h *BaseHTTPHandler) StatusBadRequest(w http.ResponseWriter, data interface{}) {
	// status code 400
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json, _ := json.Marshal(data)
	w.Write(json)
}

// StatusNotFoundRequest responses status code 404 and json.
func (h *BaseHTTPHandler) StatusNotFoundRequest(w http.ResponseWriter, data interface{}) {
	// status code 404
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json, _ := json.Marshal(data)
	w.Write(json)
}

// StatusServerError responses 500.
func (h *BaseHTTPHandler) StatusServerError(w http.ResponseWriter) {
	// status code 500
	w.WriteHeader(http.StatusInternalServerError)
}

// GetTranslaterFunc returns i18n.TranslateFunc.
// This function is necessary for translation.
func (h *BaseHTTPHandler) GetTranslaterFunc(r *http.Request) bundle.TranslateFunc {
	return r.Context().Value("i18nTfunc").(bundle.TranslateFunc)
}

// NewBaseHTTPHandler returns BaseHTTPHandler instance.
func NewBaseHTTPHandler(logger *logrus.Logger) *BaseHTTPHandler {
	return &BaseHTTPHandler{Logger: logger}
}
