package article

import "mime/multipart"

// GetRequest struct.
type GetRequest struct {
	Title string `validate:"required,max=45"`
}

// GetIDRequest struct.
type GetIDRequest struct {
	ID string `validate:"required,numeric"`
}

// PostAddRequest struct.
type PostAddRequest struct {
	Title   string `form:"title" validate:"required,max=45"`
	Content string `form:"content" validate:"required"`
}

// DeleteIDRequest struct.
type DeleteIDRequest struct {
	ID string `validate:"required,numeric"`
}

// PostVisenzeDiscoverSearchRequest struct.
type PostVisenzeDiscoverSearchRequest struct {
	Page        int            `form:"page" validate:"min=1"`
	ResultLimit int            `form:"result_limit" validate:"min=1"`
	File        multipart.File `form:"file" validate:"required"`
}
