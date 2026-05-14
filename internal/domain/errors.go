package domain

import (
	"errors"
)

var (
	ErrNotFound    = errors.New("not found")
	ErrInvalidData = errors.New("invalid data")
	ErrConflict    = errors.New("conflict")
)

type InvalidDataError struct {
	Msg string
}

func (e InvalidDataError) Error() string {
	return e.Msg
}

func (e InvalidDataError) Is(target error) bool {
	return target == ErrInvalidData
}
