package storage

import "errors"

var (
	ErrOwnerExists   = errors.New("owner already exists")
	ErrOwnerNotFound = errors.New("owner not found")
	ErrAppNotFound   = errors.New("app not found")
)
