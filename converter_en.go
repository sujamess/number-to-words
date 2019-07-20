package numbertowords

import (
	"fmt"
	"math"
	"strconv"
)

func (c *Converter) convertEN() (*Converter, error) {
	if isSmallNumber(c.Decimal) {
		c.Text = getNumberText(c.Decimal, c.Lang)

		return c, nil
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
				return nil, ErrConversion
			}

			digit := len(decimal) - idx - 1

			if isSkipDigit(digit, c.Lang) {
				continue
			}

			numberText := getNumberText(num, c.Lang)
			digitText := getDigitText(digit, c.Lang)
			isSpecialNumber := false

			if digit == thousandDigitKey || digit == millionDigitKey {
				if len(decimal) >= digit+3 || len(decimal) >= digit+2 {
					if len(decimal) >= digit+3 {
						num, err = strconv.Atoi(string(decimal[idx-2]) + string(decimal[idx-1]) + string(b))
						if err != nil {
							return nil, ErrConversion
						}

					} else if len(decimal) >= digit+2 {
						num, err = strconv.Atoi(string(decimal[idx-1]) + string(b))
						if err != nil {
							return nil, ErrConversion
						}
					}

					cvt, err := Convert(float64(num), c.Lang)
					if err != nil {
						return nil, err
					}

					cvt, err = cvt.convertEN()
					if err != nil {
						return nil, err
					}

					numberText = cvt.GetText()
				}
			}

			if digit == tenDigitKey {
				numberText, isSpecialNumber = isSpecialNumberCase(num, c.Lang)

				if !isSpecialNumber {
					numberText = getNumberText(num, c.Lang)
				}

				c.Text += fmt.Sprintf("%s%s ", numberText, digitText)
			} else {
				c.Text += fmt.Sprintf("%s %s ", numberText, digitText)
			}

			decimalLeft -= num * int(math.Pow(10, float64(digit)))
		}
	}

	if len(c.Text) == 0 {
		return nil, nil
	}

	c.Text = c.Text[:len(c.Text)-1]

	return c, nil
}

func (c *Converter) getSpecialTextEN(number, digit int) string {
	return ""
}
