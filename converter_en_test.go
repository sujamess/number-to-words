package numbertowords

import (
	"fmt"
	"testing"
)

func TestConvertEN(t *testing.T) {
	var testCases = []struct {
		givenConverter *Converter
		expectText     string
		expectErr      error
	}{
		// without floating case
		// number that contains small number
		{NewConverter(0, 0, en), "zero", nil},
		{NewConverter(19, 0, en), "nineteen", nil},
		{NewConverter(400, 0, en), "four hundred", nil},
		{NewConverter(601, 0, en), "six hundred one", nil},
		{NewConverter(719, 0, en), "seven hundred nineteen", nil},
		{NewConverter(4000, 0, en), "four thousand", nil},
		{NewConverter(6001, 0, en), "six thousand one", nil},
		{NewConverter(7019, 0, en), "seven thousand nineteen", nil},
		{NewConverter(10000, 0, en), "ten thousand", nil},
		{NewConverter(11000, 0, en), "eleven thousand", nil},
		{NewConverter(12000, 0, en), "twelve thousand", nil},
		{NewConverter(19000, 0, en), "nineteen thousand", nil},
		{NewConverter(40000, 0, en), "fourty thousand", nil},
		{NewConverter(61001, 0, en), "sixty one thousand one", nil},
		{NewConverter(72019, 0, en), "seventy two thousand nineteen", nil},
		{NewConverter(400000, 0, en), "four hundred thousand", nil},
		{NewConverter(601001, 0, en), "six hundred one thousand one", nil},
		{NewConverter(719019, 0, en), "seven hundred nineteen thousand nineteen", nil},
		{NewConverter(4000000, 0, en), "four million", nil},
		{NewConverter(6001001, 0, en), "six million one thousand one", nil},
		{NewConverter(7019019, 0, en), "seven million nineteen thousand nineteen", nil},
		{NewConverter(40000000, 0, en), "fourty million", nil},
		{NewConverter(60101001, 0, en), "sixty million one hundred one thousand one", nil},
		{NewConverter(71919019, 0, en), "seventy one million nine hundred nineteen thousand nineteen", nil},

		// special case
		{NewConverter(22, 0, en), "twenty two", nil},
		{NewConverter(33, 0, en), "thirty three", nil},
		{NewConverter(55, 0, en), "fifty five", nil},
		{NewConverter(88, 0, en), "eighty eight", nil},
		{NewConverter(222, 0, en), "two hundred twenty two", nil},
		{NewConverter(333, 0, en), "three hundred thirty three", nil},
		{NewConverter(555, 0, en), "five hundred fifty five", nil},
		{NewConverter(888, 0, en), "eight hundred eighty eight", nil},
		{NewConverter(2222, 0, en), "two thousand two hundred twenty two", nil},
		{NewConverter(3333, 0, en), "three thousand three hundred thirty three", nil},
		{NewConverter(5555, 0, en), "five thousand five hundred fifty five", nil},
		{NewConverter(8888, 0, en), "eight thousand eight hundred eighty eight", nil},
		{NewConverter(22222, 0, en), "twenty two thousand two hundred twenty two", nil},
		{NewConverter(33333, 0, en), "thirty three thousand three hundred thirty three", nil},
		{NewConverter(55555, 0, en), "fifty five thousand five hundred fifty five", nil},
		{NewConverter(88888, 0, en), "eighty eight thousand eight hundred eighty eight", nil},
		{NewConverter(222222, 0, en), "two hundred twenty two thousand two hundred twenty two", nil},
		{NewConverter(333333, 0, en), "three hundred thirty three thousand three hundred thirty three", nil},
		{NewConverter(555555, 0, en), "five hundred fifty five thousand five hundred fifty five", nil},
		{NewConverter(888888, 0, en), "eight hundred eighty eight thousand eight hundred eighty eight", nil},
		{NewConverter(2222222, 0, en), "two million two hundred twenty two thousand two hundred twenty two", nil},
		{NewConverter(3333333, 0, en), "three million three hundred thirty three thousand three hundred thirty three", nil},
		{NewConverter(5555555, 0, en), "five million five hundred fifty five thousand five hundred fifty five", nil},
		{NewConverter(8888888, 0, en), "eight million eight hundred eighty eight thousand eight hundred eighty eight", nil},
		{NewConverter(22222222, 0, en), "twenty two million two hundred twenty two thousand two hundred twenty two", nil},
		{NewConverter(33333333, 0, en), "thirty three million three hundred thirty three thousand three hundred thirty three", nil},
		{NewConverter(55555555, 0, en), "fifty five million five hundred fifty five thousand five hundred fifty five", nil},
		{NewConverter(88888888, 0, en), "eighty eight million eight hundred eighty eight thousand eight hundred eighty eight", nil},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("tc %d", idx+1), func(t *testing.T) {
			var text string

			err := tc.givenConverter.Convert().Scan(&text).Error

			if tc.expectErr != err {
				t.Errorf("Expect `%v` got `%v`", tc.expectErr, err)
			}

			if tc.expectText != text {
				t.Errorf("Expect `%s` got `%s`", tc.expectText, text)
			}
		})
	}
}
