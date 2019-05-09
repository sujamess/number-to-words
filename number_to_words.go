package numbertowords

import (
	"strconv"
	"strings"
)

// Convert converts given number to text bath
// default lang = "en"
func Convert(number float64, lang string) (string, error) {
	switch strings.ToUpper(lang) {
	case th:
		lang = th
	case en:
		lang = en
	default:
		lang = en
	}

	return convert(number, lang)
}

func convert(number float64, lang string) (string, error) {
	switch lang {
	case th:
		return fullWordsTH(number)
	default:
		return fullWordsEN(number)
	}
}

func getInteger(number float64) string {
	return strings.Split(strconv.FormatFloat(number, 'f', 2, 64), ".")[0]
}

func getDecimal(number float64) string {
	return strings.Split(strconv.FormatFloat(number, 'f', 2, 64), ".")[1]
}
