package gist

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// RepositoryInterface interface.
type RepositoryInterface interface {
	DoGistsRequest(user string) (io.Reader, error)
}

// Repository struct.
type Repository struct{}

// DoGistsRequest responses gists data.
func (g *Repository) DoGistsRequest(user string) (io.Reader, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/gists", user))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return nil, err
	}

	return &buf, nil
}

// NewRepository responses new Repository instance.
func NewRepository() RepositoryInterface {
	return &Repository{}
}
