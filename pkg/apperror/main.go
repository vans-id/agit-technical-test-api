package apperror

import "errors"

var (
	ErrInvalidInput  = errors.New("invalid input")
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")

	ErrMissingToken  = errors.New("missing token")
	ErrNotAuthorized = errors.New("not authorized")
	ErrWrongUser     = errors.New("wrong user")

	ErrInternal = errors.New("something went wrong")
)
