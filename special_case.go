package numbertowords

// digit
var specialCaseDigitList = map[string]digitText{
	en: {
		tenDigitKey:         "teen",
		tenThousandDigitKey: "ten thousand",
		tenMillionDigitKey:  "teen million",
	},
}

type skipDigit map[int]bool

var skipDigitList = map[string]skipDigit{
	en: {
		digitKey:                false,
		tenDigitKey:             false,
		hundredDigitKey:         false,
		thousandDigitKey:        false,
		tenThousandDigitKey:     true,
		hundredThousandDigitKey: true,
		millionDigitKey:         false,
		tenMillionDigitKey:      true,
	},
	th: {
		digitKey:                false,
		tenDigitKey:             false,
		hundredDigitKey:         false,
		thousandDigitKey:        false,
		tenThousandDigitKey:     false,
		hundredThousandDigitKey: false,
		millionDigitKey:         false,
		tenMillionDigitKey:      true,
	},
}

func isSpecialDigitCase(digit int, lang string) (string, bool) {
	text, exists := specialCaseDigitList[lang][digit]

	return text, exists
}

func isSkipDigit(digit int, lang string) bool {
	return skipDigitList[lang][digit]
}

// number
var specialCaseNumberList = map[string]numberText{
	en: {
		2: "twen",
		3: "thir",
		5: "fif",
		8: "eigh",
	},
	th: {
		1: "เอ็ด",
		2: "ยี่",
	},
}

func isSpecialNumberCase(number int, lang string) (string, bool) {
	text, exists := specialCaseNumberList[lang][number]

	return text, exists
}

func getSpecialNumber(number int, lang string) string {
	return specialCaseNumberList[lang][number]
}
