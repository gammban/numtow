package curtow

import (
	"github.com/gammban/numtow/lang"
	"github.com/gammban/numtow/lang/en"
	"github.com/gammban/numtow/lang/kz"
	"github.com/gammban/numtow/lang/ru"
)

// Float64 converts amount to words
func Float64(amount float64, language lang.Lang, options ...interface{}) (words string, err error) {
	switch language {
	case lang.KZ:
		o := kz.ParseCurrencyOpts(options...)

		return kz.CurrencyFloat64(amount, o...)
	case lang.RU:
		o := ru.ParseCurrencyOpts(options...)

		return ru.CurrencyFloat64(amount, o...)
	case lang.EN:
		o := en.ParseCurrencyOpts(options...)

		return en.CurrencyFloat64(amount, o...)
	case lang.Unknown:
		return words, lang.ErrBadLanguage
	default:
		return words, lang.ErrBadLanguage
	}
}

// MustFloat64 converts amount to words, on error returns empty string.
func MustFloat64(amount float64, language lang.Lang, options ...interface{}) string {
	words, err := Float64(amount, language, options...)
	if err != nil {
		return ""
	}

	return words
}
