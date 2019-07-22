package numbertowords

import (
	"errors"
	"strings"
)

// List of errors
var (
	ErrInvalidNumber    = errors.New("Invalid a given number")
	ErrLangNotSupported = errors.New("This language is not supported")
	ErrConversion       = errors.New("conversion error")
)

type customError struct {
	err errorList
}

type errorList []error

func (cErr *customError) Error() string {
	errMsg := make([]string, 0)

	for _, e := range cErr.err {
		errMsg = append(errMsg, e.Error())
	}

	return strings.Join(errMsg, "; ")
}

// AddError adds an error to slice of error
func AddError(err, newErr error) error {
	if err == nil {
		err = &customError{}
	}

	customErr, ok := err.(*customError)
	if ok {
		return &customError{
			err: append(customErr.err, newErr),
		}
	}

	return &customError{
		err: errorList{newErr},
	}
}
