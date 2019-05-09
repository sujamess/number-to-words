package numbertowords

const (
	digitKey                int = 0
	tenDigitKey             int = 1
	hundredDigitKey         int = 2
	thousandDigitKey        int = 3
	tenThousandDigitKey     int = 4
	hundredThousandDigitKey int = 5
	millionDigitKey         int = 6
	tenMillionDigitKey      int = 7
)

var digitList = map[string]map[int]string{
	en: {
		digitKey:                "",
		tenDigitKey:             "ty",
		hundredDigitKey:         "hundred",
		thousandDigitKey:        "thousand",
		tenThousandDigitKey:     "tythousand",
		hundredThousandDigitKey: "hundredthousand",
		millionDigitKey:         "million",
		tenMillionDigitKey:      "tymillion",
	},
	th: {
		digitKey:                "",
		tenDigitKey:             "สิบ",
		hundredDigitKey:         "ร้อย",
		thousandDigitKey:        "พัน",
		tenThousandDigitKey:     "หมื่น",
		hundredThousandDigitKey: "แสน",
		millionDigitKey:         "ล้าน",
		tenMillionDigitKey:      "สิบล้าน",
	},
}

var specialCaseDigit = map[string]map[int]string{
	en: {
		tenDigitKey:         "teen",
		tenThousandDigitKey: "tenthousand",
		tenMillionDigitKey:  "teenmillion",
	},
}
