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
	d, fp := splits[0], splits[1]

	decimal, err := strconv.Atoi(d)
	if err != nil {
		return nil, ErrConversion
	}

	floating, err := strconv.Atoi(fp)
	if err != nil {
		return nil, ErrConversion
	}

	return NewConverter(decimal, floating, lang), nil
}
