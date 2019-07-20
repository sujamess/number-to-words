package numbertowords

// ValidateLang validates that lang is supported or not
func validateLang(givenLang string) bool {
	for _, lang := range supportedLang {
		if givenLang == lang {
			return true
		}
	}

	return false
}
