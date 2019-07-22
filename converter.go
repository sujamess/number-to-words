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
	Error    error
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
func (c *Converter) Convert() *Converter {
	if !c.validateLang() {
		c.Error = AddError(c.Error, ErrLangNotSupported)
	}

	var err error

	switch c.Lang {
	case supportedLang[en]:
		c.Text, err = convertEN(c.Decimal, c.Floating, c.Lang)
	case supportedLang[th]:
	}

	if err != nil {
		c.Error = AddError(c.Error, err)
	}

	return c
}

func (c *Converter) validateLang() bool {
	if c.Lang == "" {
		c.Lang = supportedLang[en]
	}

	c.Lang = strings.ToLower(c.Lang)

	return validateLang(c.Lang)
}

// Scan gets text
func (c *Converter) Scan(text *string) *Converter {
	*text = c.Text

	return c
}

// SetCurrency append currency to converted text
func (c *Converter) SetCurrency(currency string) *Converter {
	c.Currency = currency

	return c
}
