package apperror

import "errors"

var (
	ErrInvalidInput  = errors.New("invalid input")
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")

	ErrMissingToken  = errors.New("missing token")
	ErrNotAuthorized = errors.New("not authorized")
	ErrInvalidCred   = errors.New("wrong username or password")

	ErrInternal = errors.New("something went wrong")
)
