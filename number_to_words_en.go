package numbertowords

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func fullWordsEN(number float64) (string, error) {
	var (
		integer = getInteger(number)
		decimal = getDecimal(number)
	)

	integerWord, err := integerEN(integer, en)
	if err != nil {
		return "", err
	}

	decimalWord, err := decimalEN(decimal, en)
	if err != nil {
		return "", err
	}

	currencyWord := getCurrency(en)

	fw := integerWord

	if len(decimalWord) != 0 {
		fw = fmt.Sprintf("%s %s", fw, decimalWord)
	}

	if len(currencyWord) != 0 {
		fw = fmt.Sprintf("%s %s", fw, currencyWord)
	}

	return fw, nil
}

func integerEN(str, lang string) (string, error) {
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
			word = getSmallNumber(int(integer), lang)
			integer = 0.00
		} else {
			digitKey := integerLength - idx - 1

			switch digitKey {
			case tenThousandDigitKey, hundredThousandDigitKey, tenMillionDigitKey:
				continue
			case thousandDigitKey:
				if integerLength >= 6 {
					num = string(str[idx-2]) + string(str[idx-1]) + string(c)
				} else if integerLength >= 5 {
					num = string(str[idx-1]) + string(c)
				}
			case millionDigitKey:
				if integerLength >= 9 {
					num = string(str[idx-2]) + string(str[idx-1]) + string(c)
				} else if integerLength >= 8 {
					num = string(str[idx-1]) + string(c)
				}
			}

			numKey, err := strconv.ParseFloat(num, 64)
			if err != nil {
				panic(err)
			}

			if numKey == 0.00 {
				continue
			}

			integer -= numKey * math.Pow(10, float64(digitKey))

			word = getWord(int(numKey), digitKey, lang)
		}

		wordList = append(wordList, word)
	}

	return strings.Join(wordList, " "), nil
}

func decimalEN(str, lang string) (string, error) {
	if str == "00" {
		return "", nil
	}

	wordList := make([]string, 0)
	wordList = append(wordList, "point")

	for idx, c := range str {
		i, err := strconv.Atoi(string(c))
		if err != nil {
			return "", fmt.Errorf("Cannot convert %s to type integer", string(c))
		}

		if idx == len(str)-1 && i == 0 {
			continue
		}

		wordList = append(wordList, getNumber(i, lang))
	}

	return strings.Join(wordList, " "), nil
}
