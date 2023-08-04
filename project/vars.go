package project

import "errors"

var (
	// ErrFileNotFound project not found
	ErrFileNotFound = errors.New("File not found")
	// ErrNotFound project not found
	ErrNotFound = errors.New("Project not found")
	// ErrDuplicate project is allready in list when it should not
	ErrDuplicate = errors.New("Project allready found")
	ErrListEmpty = errors.New("List is empty")
	ErrListTooLong = errors.New("List is too long")
)
