package apiErrors

import "errors"

var (
	// ErrUnexpected represents an unexpected error
	ErrUnexpected = errors.New("An unexpected error as occur")
)
