package numbertowords

var currency = map[string]string{
	en: "baht",
	th: "บาท",
}

func getCurrency(lang string) string {
	return currency[lang]
}
