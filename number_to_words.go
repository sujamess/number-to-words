package numbertowords

import (
	"strconv"
	"strings"
)

// Convert converts number to text (supporting English and Thai), default lang is English
// Supporting 2 precision
func Convert(number float64, lang string) *Converter {
	converter := NewConverter(0, 0, "")

	formattedNumber := strconv.FormatFloat(number, 'f', 2, 64)

	splits := strings.Split(formattedNumber, ".")

	decimal, err := strconv.Atoi(splits[0])
	if err != nil {
		converter.Error = AddError(converter.Error, ErrConversion)
	}

	floating, err := strconv.Atoi(splits[1])
	if err != nil {
		converter.Error = AddError(converter.Error, ErrConversion)
	}

	converter.Decimal = decimal
	converter.Floating = floating

	return converter
}
