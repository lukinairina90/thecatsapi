package catsapi

import "errors"

var (
	ErrBadRequest       = errors.New("bad request")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrNotFound         = errors.New("not found")
	ErrInternal         = errors.New("internal error: something went wrong")
	ErrMethodNotAllowed = errors.New("method not allowed")
)
