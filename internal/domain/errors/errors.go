package errors

import "errors"

var (
	ErrInvalidTitle       = errors.New("invalid title")
	ErrInvalidDescription = errors.New("invalid description")
)
