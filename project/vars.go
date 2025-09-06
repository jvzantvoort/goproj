package project

import "errors"

var (
	// ErrFileNotFound project not found
	ErrFileNotFound = errors.New("File not found")
	// ErrNotFound project not found
	ErrNotFound = errors.New("Project not found")
	// ErrDuplicate project is already in list when it should not
	ErrDuplicate = errors.New("Project already found")
	// ErrListEmpty missing godoc.
	ErrListEmpty = errors.New("List is empty")
	// ErrListTooLong missing godoc.
	ErrListTooLong = errors.New("List is too long")
)
