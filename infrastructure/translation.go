package infrastructure

import (
	"path/filepath"

	"github.com/dalu/i18n"
)

const (
	// DirTranslation is translation path directory.
	DirTranslation = "translation"
)

// Translation struct.
type Translation struct {
	Middleware *i18n.I18nMiddleware
}

// NewTranslation returns new Translation.
// repository: https://github.com/dalu/i18n
func NewTranslation() *Translation {
	files, err := filepath.Glob(DirTranslation + "/*.json")
	if err != nil {
		panic(err)
	}
	c := i18n.Config{DefaultLanguage: GetConfigString("language.default"),
		Files: files,
		Debug: false,
	}
	return &Translation{Middleware: i18n.New(c)}
}
