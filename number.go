package numbertowords

var numberList = map[string]map[int]string{
	en: {
		0:  "zero",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
	},
	th: {
		0:  "ศูนย์",
		1:  "หนึ่ง",
		2:  "สอง",
		3:  "สาม",
		4:  "สี่",
		5:  "ห้า",
		6:  "หก",
		7:  "เจ็ด",
		8:  "แปด",
		9:  "เก้า",
		10: "สิบ",
		11: "สิบเอ็ด",
		12: "สิบสอง",
		13: "สิบสาม",
		14: "สิบสี่",
		15: "สิบห้า",
		16: "สิบหก",
		17: "สิบเจ็ด",
		18: "สิบแปด",
		19: "สิบเก้า",
	},
}

var specialCaseNumber = map[string]map[int]string{
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

func isSmallNumber(number int) bool {
	return number < 20
}

func getNumber(number int, lang string) string {
	return numberList[lang][number]
}

func getSpecialNumber(number int, lang string) string {
	return specialCaseNumber[lang][number]
}
