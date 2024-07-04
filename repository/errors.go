package repository

import "errors"

var (
	ErrForeignKeyNotFound = errors.New("related resource not found")
)
