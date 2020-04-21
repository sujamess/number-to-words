package numbertowords

import (
	"math"
	"strconv"
	"strings"
)

func fullWordsTH(number float64) (string, error) {
	var (
		integer = getInteger(number)
		decimal = getDecimal(number)
	)

	var fw string

	if integer != "0" {
		integerWord, err := readerTH(integer, th)
		if err != nil {
			return "", err
		}

		fw += integerWord
	} else {
		if decimal == "00" {
			fw += "ศูนย์"
		}
	}

	if len(fw) > 0 {
		currencyWord := getCurrency(th)

		fw += currencyWord
	}

	decimalWord, err := decimalTH(decimal, th)
	if err != nil {
		return "", err
	}

	fw += decimalWord

	return fw, nil
}

func readerTH(str, lang string) (string, error) {
	integer, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return "", err
	}

	if isSmallNumber(int(integer)) {
		return getSmallNumber(int(integer), lang), nil
	}

	integerLength := len(str)

	wordList := make([]string, 0)

	for idx, c := range str {
		if integer == 0.00 {
			break
		}

		var (
			word string
			num  = string(c)
		)

		if isSmallNumber(int(integer)) {
			switch int(integer) {
			case 1:
				word = getSpecialNumber(int(integer), lang)
			default:
				word = getSmallNumber(int(integer), lang)
			}

			integer = 0.00
		} else {
			digitKey := integerLength - idx - 1

			switch digitKey {
			case tenMillionDigitKey:
				continue
			case millionDigitKey:
				if integerLength >= 9 {
					num = string(str[idx-2]) + string(str[idx-1]) + string(c)
				} else if integerLength >= 8 {
					num = string(str[idx-1]) + string(c)
				}
			}

			numKey, err := strconv.ParseFloat(num, 64)
			if err != nil {
				return "", err
			}

			if numKey == 0.00 {
				continue
			}

			integer -= numKey * math.Pow(10, float64(digitKey))

			word = getWord(int(numKey), digitKey, lang)
		}

		wordList = append(wordList, word)
	}

	return strings.Join(wordList, ""), nil
}

func decimalTH(str, lang string) (string, error) {
	if str == "00" {
		return "ถ้วน", nil
	}

	result, err := readerTH(str, lang)
	if err != nil {
		return "", err
	}

	return result + "สตางค์", nil
}
