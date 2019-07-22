package numbertowords

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func convertEN(decimal, floating int, lang string) (string, error) {
	var text string

	if isSmallNumber(decimal) {
		text = getNumberText(decimal, lang)

		return text, nil
	}

	dcm, _ := []byte(strconv.Itoa(decimal)), strconv.Itoa(floating)
	decimalLeft, _ := decimal, floating

	for idx, b := range dcm {
		if decimalLeft == 0 {
			break
		}

		fmt.Printf("decimalLeft: %d\n", decimalLeft)

		if isSmallNumber(decimalLeft) {
			numberText := getNumberText(decimalLeft, lang)

			switch lang {
			case th:
				if decimalLeft == 1 && digitKey == digitKey {
					numberText = getSpecialNumber(decimalLeft, lang)
				}
			}

			text += fmt.Sprintf("%s ", numberText)
			break
		} else {
			num, err := strconv.Atoi(string(b))
			if err != nil {
				return "", err
			}

			digit := len(dcm) - idx - 1

			if isSkipDigit(digit, lang) {
				continue
			}

			fmt.Printf("lang = %s, digit: %d, num = %d\n", lang, digit, num)

			numberText := getNumberText(num, lang)
			digitText := getDigitText(digit, lang)
			isSpecialNumber := false

			if (digit == thousandDigitKey && lang == en) || (digit == millionDigitKey) {
				if len(dcm) >= digit+3 || len(dcm) >= digit+2 {
					if len(dcm) >= digit+3 {
						num, err = strconv.Atoi(string(dcm[idx-2]) + string(dcm[idx-1]) + string(b))
						if err != nil {
							return "", ErrConversion
						}

					} else if len(dcm) >= digit+2 {
						num, err = strconv.Atoi(string(dcm[idx-1]) + string(b))
						if err != nil {
							return "", ErrConversion
						}
					}

					numberText, err = convertEN(num, 0.00, lang)
					if err != nil {
						return "", err
					}
				}
			}

			if digit == tenDigitKey {
				numberText, isSpecialNumber = isSpecialNumberCase(num, lang)

				if !isSpecialNumber {
					numberText = getNumberText(num, lang)
				}

				text += fmt.Sprintf("%s%s ", numberText, digitText)
			} else {
				text += fmt.Sprintf("%s %s ", numberText, digitText)
			}

			decimalLeft -= num * int(math.Pow(10, float64(digit)))
		}
	}

	if lang == th {
		text = strings.ReplaceAll(text, " ", "")
	}

	if len(text) > 0 && lang != th {
		text = text[:len(text)-1]
	}

	return text, nil
}
