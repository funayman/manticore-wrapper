package manticore

import "errors"

var (
	ErrNotImplemented = errors.New("function not implemented")
	ErrNotSupported   = errors.New("not supported with manticore")
)
