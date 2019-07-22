package numbertowords

import (
	"strconv"
	"strings"
)

// Convert converts number to text (supporting English and Thai), default lang is English
// Supporting 2 precision
func Convert(number float64, lang string) (*Converter, error) {
	formattedNumber := strconv.FormatFloat(number, 'f', 2, 64)

	splits := strings.Split(formattedNumber, ".")

	decimal, err := strconv.Atoi(splits[0])
	if err != nil {
		return nil, ErrInvalidNumber
	}

	floating, err := strconv.Atoi(splits[1])
	if err != nil {
		return nil, ErrInvalidNumber
	}

	return NewConverter(decimal, floating, lang), nil
}
