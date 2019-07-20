package numbertowords

import "errors"

// List of errors
var (
	ErrConversion       = errors.New("Invalid a given number")
	ErrLangNotSupported = errors.New("This language is not supported")
)
