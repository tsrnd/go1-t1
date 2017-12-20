package gist

import "encoding/json"

// UsecaseInterface interface.
type UsecaseInterface interface {
	ListGists(user string) ([]string, error)
}

// Usecase struct.
type Usecase struct {
	Gister RepositoryInterface
}

// ListGists return url list
func (c *Usecase) ListGists(user string) ([]string, error) {
	r, err := c.Gister.DoGistsRequest(user)
	if err != nil {
		return nil, err
	}

	var gists []Gist
	if err := json.NewDecoder(r).Decode(&gists); err != nil {
		return nil, err
	}

	urls := make([]string, 0, len(gists))
	for _, u := range gists {
		urls = append(urls, u.Rawurl)
	}

	return urls, nil
}

// NewUsecase responses new Usecase instance.
func NewUsecase(g RepositoryInterface) UsecaseInterface {
	return &Usecase{g}
}
