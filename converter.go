package numbertowords

import (
	"strings"
)

// Converter is a struct that for converting
type Converter struct {
	Decimal  int
	Floating int
	Lang     string
	Text     string
	Currency string
}

// NewConverter returns a new converter
func NewConverter(decimal, floating int, lang string) *Converter {
	return &Converter{
		Decimal:  decimal,
		Floating: floating,
		Lang:     lang,
	}
}

// Convert converts string to text
func (c *Converter) Convert() (*Converter, error) {
	if !c.validateLang() {
		return nil, ErrLangNotSupported
	}

	switch c.Lang {
	case supportedLang[en]:
		return c.convertEN()
	case supportedLang[th]:
	}

	return nil, ErrLangNotSupported
}

func (c *Converter) validateLang() bool {
	if c.Lang == "" {
		c.Lang = supportedLang[en]
	}

	c.Lang = strings.ToLower(c.Lang)

	return validateLang(c.Lang)
}

// GetText gets text
func (c *Converter) GetText() string {
	return c.Text
}

// SetCurrency append currency to converted text
func (c *Converter) SetCurrency(currency string) *Converter {
	c.Currency = currency

	return c
}
