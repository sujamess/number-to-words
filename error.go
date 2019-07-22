package numbertowords

import "errors"

// List of errors
var (
	ErrInvalidNumber    = errors.New("Invalid a given number")
	ErrLangNotSupported = errors.New("This language is not supported")
	ErrConversion       = errors.New("conversion error")
)
