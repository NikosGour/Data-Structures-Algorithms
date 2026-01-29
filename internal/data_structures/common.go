package data_structures

import (
	"errors"
)

var (
	ErrIndexOutOfBounds  = errors.New("index out of bounds")
	ErrNotAllItemsCopied = errors.New("not all items copied")
	ErrUnreachable       = errors.New("unreachable")
)
