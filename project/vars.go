package project

import "errors"

var (
	// ErrFileNotFound project not found
	ErrFileNotFound = errors.New("file not found")
	// ErrNotFound project not found
	ErrNotFound = errors.New("project not found")
	// ErrDuplicate project is already in list when it should not
	ErrDuplicate = errors.New("project already found")
	// ErrListEmpty missing godoc.
	ErrListEmpty = errors.New("list is empty")
	// ErrListTooLong missing godoc.
	ErrListTooLong = errors.New("list is too long")
)
