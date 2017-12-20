package article

import "errors"

var (
	// ErrorInternalServer add errors.
	ErrorInternalServer = errors.New("internal Server error")
	// ErrorNotFound add errors.
	ErrorNotFound = errors.New("your requested Item is not found")
	// ErrorConflict add errors.
	ErrorConflict = errors.New("your Item already exist")
)
