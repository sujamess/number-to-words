package numbertowords

import (
	"strconv"
	"testing"
)

func TestIntegerTH(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"0", "ศูนย์"},
		{"321", "สามร้อยยี่สิบเอ็ด"},
		{"31", "สามสิบเอ็ด"},
		{"5", "ห้า"},
		{"15", "สิบห้า"},
		{"22", "ยี่สิบสอง"},
		{"555", "ห้าร้อยห้าสิบห้า"},
		{"515", "ห้าร้อยสิบห้า"},
		{"4414", "สี่พันสี่ร้อยสิบสี่"},
		{"15512", "หนึ่งหมื่นห้าพันห้าร้อยสิบสอง"},
		{"52152", "ห้าหมื่นสองพันหนึ่งร้อยห้าสิบสอง"},
		{"15155", "หนึ่งหมื่นห้าพันหนึ่งร้อยห้าสิบห้า"},
		{"141512", "หนึ่งแสนสี่หมื่นหนึ่งพันห้าร้อยสิบสอง"},
		{"1111111", "หนึ่งล้านหนึ่งแสนหนึ่งหมื่นหนึ่งพันหนึ่งร้อยสิบเอ็ด"},
		{"22222222", "ยี่สิบสองล้านสองแสนสองหมื่นสองพันสองร้อยยี่สิบสอง"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := readerTH(tc.input, th)
			if err != nil {
				t.Errorf("Error occurred while reading integer: %v", err)
			}

			if got != tc.output {
				t.Errorf("got %v, expect %v", got, tc.output)
			}
		})
	}
}

func TestDecimalTH(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"00", "ถ้วน"},
		{"01", "หนึ่งสตางค์"},
		{"10", "สิบสตางค์"},
		{"15", "สิบห้าสตางค์"},
		{"20", "ยี่สิบสตางค์"},
		{"25", "ยี่สิบห้าสตางค์"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := decimalTH(tc.input, th)
			if err != nil {
				t.Errorf("Error occurred while converting the decimal in Thai: %v", err)
			}

			if got != tc.output {
				t.Errorf("got %v, expect %v", got, err)
			}
		})
	}
}

func TestConvertTH(t *testing.T) {
	testCases := []struct {
		input  float64
		output string
	}{
		{0.00, "ศูนย์บาทถ้วน"},
		{0.25, "ยี่สิบห้าสตางค์"},
		{9.99, "เก้าบาทเก้าสิบเก้าสตางค์"},
		{10.00, "สิบบาทถ้วน"},
		{22.22, "ยี่สิบสองบาทยี่สิบสองสตางค์"},
		{88.88, "แปดสิบแปดบาทแปดสิบแปดสตางค์"},
		{111.11, "หนึ่งร้อยสิบเอ็ดบาทสิบเอ็ดสตางค์"},
	}

	for _, tc := range testCases {
		t.Run(strconv.FormatFloat(tc.input, 'f', 2, 64), func(t *testing.T) {
			got, err := Convert(tc.input, th)
			if err != nil {
				t.Errorf("Error occurred while converting the number into Thai: %v", err)
			}

			if got != tc.output {
				t.Errorf("got %v, expect %v", got, tc.output)
			}
		})
	}
}

func BenchmarkReaderTH(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readerTH("22222222", th)
	}
}

func BenchmarkDecimalTH(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decimalTH("25", th)
	}
}

func BenchmarkConvertTH(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert(111.11, th)
	}
}
