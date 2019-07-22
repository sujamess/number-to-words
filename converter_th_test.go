package numbertowords

import (
	"fmt"
	"testing"
)

func TestConvertTH(t *testing.T) {
	var testCases = []struct {
		givenConverter *Converter
		expectText     string
		expectErr      error
	}{
		// without floating case
		// number that contains small number
		{NewConverter(0, 0, th), "ศูนย์", nil},
		{NewConverter(19, 0, th), "สิบเก้า", nil},
		{NewConverter(30, 0, th), "สามสิบ", nil},
		{NewConverter(65, 0, th), "หกสิบห้า", nil},
		{NewConverter(400, 0, th), "สี่ร้อย", nil},
		{NewConverter(435, 0, th), "สี่ร้อยสามสิบห้า", nil},
		{NewConverter(601, 0, th), "หกร้อยเอ็ด", nil},
		{NewConverter(719, 0, th), "เจ็ดร้อยสิบเก้า", nil},
		{NewConverter(4000, 0, th), "สี่พัน", nil},
		{NewConverter(6001, 0, th), "หกพันเอ็ด", nil},
		{NewConverter(7019, 0, th), "เจ็ดพันสิบเก้า", nil},
		{NewConverter(10000, 0, th), "หนึ่งหมื่น", nil},
		{NewConverter(11000, 0, th), "หนึ่งหมื่นหนึ่งพัน", nil},
		{NewConverter(12000, 0, th), "หนึ่งหมื่นสองพัน", nil},
		{NewConverter(19000, 0, th), "หนึ่งหมื่นเก้าพัน", nil},
		{NewConverter(40000, 0, th), "สี่หมื่น", nil},
		{NewConverter(61001, 0, th), "หกหมื่นหนึ่งพันเอ็ด", nil},
		{NewConverter(72019, 0, th), "เจ็ดหมื่นสองพันสิบเก้า", nil},
		{NewConverter(400000, 0, th), "สี่แสน", nil},
		{NewConverter(601001, 0, th), "หกแสนหนึ่งพันเอ็ด", nil},
		{NewConverter(719019, 0, th), "เจ็ดแสนหนึ่งหมื่นเก้าพันสิบเก้า", nil},
		{NewConverter(4000000, 0, th), "สี่ล้าน", nil},
		{NewConverter(6001001, 0, th), "หกล้านหนึ่งพันเอ็ด", nil},
		{NewConverter(7019019, 0, th), "เจ็ดล้านหนึ่งหมื่นเก้าพันสิบเก้า", nil},
		{NewConverter(40000000, 0, th), "สี่สิบล้าน", nil},
		{NewConverter(60101001, 0, th), "หกสิบล้านหนึ่งแสนหนึ่งพันเอ็ด", nil},
		{NewConverter(71919019, 0, th), "เจ็ดสิบเอ็ดล้านเก้าแสนหนึ่งหมื่นเก้าพันสิบเก้า", nil},

		// special case
		{NewConverter(22, 0, th), "ยี่สิบสอง", nil},
		{NewConverter(33, 0, th), "สามสิบสาม", nil},
		{NewConverter(55, 0, th), "ห้าสิบห้า", nil},
		// {NewConverter(88, 0, th), "eighty eight", nil},
		// {NewConverter(222, 0, th), "two hundred twenty two", nil},
		// {NewConverter(333, 0, th), "three hundred thirty three", nil},
		// {NewConverter(555, 0, th), "five hundred fifty five", nil},
		// {NewConverter(888, 0, th), "eight hundred eighty eight", nil},
		// {NewConverter(2222, 0, th), "two thousand two hundred twenty two", nil},
		// {NewConverter(3333, 0, th), "three thousand three hundred thirty three", nil},
		// {NewConverter(5555, 0, th), "five thousand five hundred fifty five", nil},
		// {NewConverter(8888, 0, th), "eight thousand eight hundred eighty eight", nil},
		// {NewConverter(22222, 0, th), "twenty two thousand two hundred twenty two", nil},
		// {NewConverter(33333, 0, th), "thirty three thousand three hundred thirty three", nil},
		// {NewConverter(55555, 0, th), "fifty five thousand five hundred fifty five", nil},
		// {NewConverter(88888, 0, th), "eighty eight thousand eight hundred eighty eight", nil},
		// {NewConverter(222222, 0, th), "two hundred twenty two thousand two hundred twenty two", nil},
		// {NewConverter(333333, 0, th), "three hundred thirty three thousand three hundred thirty three", nil},
		// {NewConverter(555555, 0, th), "five hundred fifty five thousand five hundred fifty five", nil},
		// {NewConverter(888888, 0, th), "eight hundred eighty eight thousand eight hundred eighty eight", nil},
		// {NewConverter(2222222, 0, th), "two million two hundred twenty two thousand two hundred twenty two", nil},
		// {NewConverter(3333333, 0, th), "three million three hundred thirty three thousand three hundred thirty three", nil},
		// {NewConverter(5555555, 0, th), "five million five hundred fifty five thousand five hundred fifty five", nil},
		// {NewConverter(8888888, 0, th), "eight million eight hundred eighty eight thousand eight hundred eighty eight", nil},
		// {NewConverter(22222222, 0, th), "twenty two million two hundred twenty two thousand two hundred twenty two", nil},
		// {NewConverter(33333333, 0, th), "thirty three million three hundred thirty three thousand three hundred thirty three", nil},
		// {NewConverter(55555555, 0, th), "fifty five million five hundred fifty five thousand five hundred fifty five", nil},
		// {NewConverter(88888888, 0, th), "eighty eight million eight hundred eighty eight thousand eight hundred eighty eight", nil},
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
