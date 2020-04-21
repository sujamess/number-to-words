package numbertowords

import (
	"strconv"
	"testing"
)

func TestIntegerEN(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"0", "zero"},
		{"5", "five"},
		{"11", "eleven"},
		{"15", "fifteen"},
		{"88", "eighty eight"},
		{"44", "fourty four"},
		{"111", "onehundred eleven"},
		{"555", "fivehundred fifty five"},
		{"515", "fivehundred fifteen"},
		{"4414", "fourthousand fourhundred fourteen"},
		{"15512", "fifteenthousand fivehundred twelve"},
		{"52152", "fifty twothousand onehundred fifty two"},
		{"15155", "fifteenthousand onehundred fifty five"},
		{"141512", "onehundred fourty onethousand fivehundred twelve"},
		{"1111111", "onemillion onehundred eleventhousand onehundred eleven"},
		{"22222222", "twenty twomillion twohundred twenty twothousand twohundred twenty two"},
		{"88888888", "eighty eightmillion eighthundred eighty eightthousand eighthundred eighty eight"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := integerEN(tc.input, en)
			if err != nil {
				t.Errorf("Error occurred while converting integer to words: %v", err)
			}

			if got != tc.output {
				t.Errorf("got %v, expect %v", got, tc.output)
			}
		})
	}
}

func TestDecimalEN(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"00", ""},
		{"01", "point zero one"},
		{"10", "point one"},
		{"15", "point one five"},
		{"20", "point two"},
		{"25", "point two five"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := decimalEN(tc.input, en)
			if err != nil {
				t.Errorf("Error occurred while converting decimal to words: %v", err)
			}

			if got != tc.output {
				t.Errorf("got %v, expect %v", got, tc.output)
			}
		})
	}
}

func TestConvertEN(t *testing.T) {
	testCases := []struct {
		input  float64
		output string
	}{
		{0.00, "zero baht"},
		{0.25, "zero point two five baht"},
		{9.99, "nine point nine nine baht"},
		{10.00, "ten baht"},
		{22.22, "twenty two point two two baht"},
		{88.88, "eighty eight point eight eight baht"},
		{111.11, "onehundred eleven point one one baht"},
	}

	for _, tc := range testCases {
		t.Run(strconv.FormatFloat(tc.input, 'f', 2, 64), func(t *testing.T) {
			got, err := Convert(tc.input, en)
			if err != nil {
				t.Errorf("Error occurred while converting: %v", err)
			}

			if got != tc.output {
				t.Errorf("got %v, expect %v", got, tc.output)
			}
		})
	}
}

func BenchmarkIntegerEN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		integerEN("88888888", en)
	}
}

func BenchmarkDecimalEN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decimalEN("25", en)
	}
}

func BenchmarkConvertEN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert(111.11, en)
	}
}
