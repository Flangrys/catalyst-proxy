package services

import "errors"

var (
	ErrNullPointer       = errors.New("null pointer")
	ErrIlegalServiceInit = errors.New("illegal service initialization")
)
