package model

import "errors"

var (
	ErrDefault  = errors.New("something went wrong")
	ErrNotFound = errors.New("not found")
)
