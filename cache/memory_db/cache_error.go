package memory_db

import "errors"

var (
	ErrValueNotPointer = errors.New("db value must be pointer type")
	ErrNotFound        = errors.New("can not find the value of string")
)
