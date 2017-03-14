package apiErrors

import "errors"

var (
	// ErrUnexpected represents an unexpected error
	ErrUnexpected = errors.New("An unexpected error as occur")

	// ErrNotFoundOnCache represents an error when the key was not found on cache
	ErrNotFoundOnCache = errors.New("Not found on cache")

	// ErrGetCacheValue represents an error when an error occur when get cache value
	ErrGetCacheValue = errors.New("Not found on cache")
)
