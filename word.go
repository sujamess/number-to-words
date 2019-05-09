package numbertowords

import (
	"fmt"
	"strconv"
)

func getWord(number, digit int, lang string) string {
	switch lang {
	case th:
		// ## special cases (th) ##
		// - case *1: must เอ็ด not หนึ่ง
		// - case 2*: must ยี่ not สอง
		// - case > 100,000,000: e.g. 888,000,000 must แปดร้อยแปดสิบแปดล้าน not แปดร้อยล้าน แปดสิบล้าน แปดล้าน
		if digit == digitKey {
			switch number {
			case 1:
				return getSpecialNumber(number, lang)
			}
		} else if digit == tenDigitKey {
			switch number {
			case 2:
				return fmt.Sprintf("%s%s", getSpecialNumber(number, lang), digitList[lang][digit])
			}
		} else if digit == millionDigitKey {
			ch, err := recursive(number, lang)
			if err != nil {
				panic(err)
			}

			return fmt.Sprintf("%s%s", ch, digitList[lang][digit])
		}
	case en:
		// ## special cases (en) ##
		// - case 10: must ten neither oneten nor oneteen
		// - case 20: must twenty not twoty
		// - case 50: must fifty not fivety
		// - case 10000: must tenthousand neither onetythousand nor onetenthousand
		// - case 15000: must fifteenthousand neighter fiveteenthousand nor tenthousand fivethousand
		// - case 22000: must twentytwothousand neigther twentytwothousand not twotytwothousand nor twentythousand twothousand
		// - case 50000: must fiftythousand not fivetythousand
		if digit == tenDigitKey {
			switch number {
			case 2, 3, 5, 8:
				return fmt.Sprintf("%s%s", specialCaseNumber[lang][number], digitList[lang][digit])
			}
		} else if digit == thousandDigitKey || digit == millionDigitKey {
			ch, err := recursive(number, lang)
			if err != nil {
				panic(err)
			}

			return fmt.Sprintf("%s%s", ch, digitList[lang][digit])
		}
	}

	var numberWord string

	if isSmallNumber(number) {
		numberWord = getSmallNumber(number, lang)
	} else {
		numberWord = getNumber(number, lang)
	}

	return fmt.Sprintf("%s%s", numberWord, digitList[lang][digit])
}

func getSmallNumber(number int, lang string) string {
	return numberList[lang][number]
}

func recursive(number int, lang string) (string, error) {
	str := strconv.Itoa(number)

	switch lang {
	case th:
		return readerTH(str, th)
	case en:
		return integerEN(str, en)
	default:
		return integerEN(str, en)
	}
}
