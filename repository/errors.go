package repository

import "errors"

var (
	ErrForeignKeyNotFound = errors.New("related resource not found")
	ErrDatabaseError      = errors.New("database error")
)
