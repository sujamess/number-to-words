package numbertowords

import (
	"fmt"
	"math"
	"strconv"
)

func (c *Converter) convertEN() (string, error) {
	if isSmallNumber(c.Decimal) {
		return getNumberText(c.Decimal, c.Lang), nil
	}

	decimal, _ := []byte(strconv.Itoa(c.Decimal)), strconv.Itoa(c.Floating)
	decimalLeft, _ := c.Decimal, c.Floating

	for idx, b := range decimal {
		if decimalLeft == 0 {
			break
		}

		if isSmallNumber(decimalLeft) {
			c.Text += fmt.Sprintf("%s ", getNumberText(decimalLeft, c.Lang))
			break
		} else {
			num, err := strconv.Atoi(string(b))
			if err != nil {
				return "", ErrConversion
			}

			digit := len(decimal) - idx - 1

			if isSkipDigit(digit, c.Lang) {
				continue
			}

			if digit == thousandDigitKey || digit == millionDigitKey {
				var text string

				if len(decimal) >= digit+3 || len(decimal) >= digit+2 {
					if len(decimal) >= digit+3 {
						num, err = strconv.Atoi(string(decimal[idx-2]) + string(decimal[idx-1]) + string(b))
						if err != nil {
							return "", ErrConversion
						}

					} else if len(decimal) >= digit+2 {
						num, err = strconv.Atoi(string(decimal[idx-1]) + string(b))
						if err != nil {
							return "", ErrConversion
						}
					}

					text, err = Convert(float64(num), c.Lang)
					if err != nil {
						return "", err
					}
				} else {
					text = getNumberText(num, c.Lang)
				}

				c.Text += fmt.Sprintf("%s %s ", text, getDigitText(digit, c.Lang))
			} else if digit == tenDigitKey {
				numberText, isSpecialNumber := isSpecialNumberCase(num, c.Lang)

				if !isSpecialNumber {
					numberText = getNumberText(num, c.Lang)
				}

				c.Text += fmt.Sprintf("%s%s ", numberText, getDigitText(digit, c.Lang))
			} else {
				c.Text += fmt.Sprintf("%s %s ", getNumberText(num, c.Lang), getDigitText(digit, c.Lang))
			}

			decimalLeft -= num * int(math.Pow(10, float64(digit)))
		}
	}

	if len(c.Text) == 0 {
		return "", nil
	}

	return c.Text[:len(c.Text)-1], nil
}

func (c *Converter) getSpecialTextEN(number, digit int) string {
	return ""
}
