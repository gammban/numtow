package numtow

import (
	"github.com/dantedenis/numtow/lang"
	"github.com/dantedenis/numtow/lang/en"
	"github.com/dantedenis/numtow/lang/kz"
	"github.com/dantedenis/numtow/lang/ru"
)

// Float64 converts float64 number to words.
func Float64(decimal float64, language lang.Lang, options ...interface{}) (words string, err error) {
	switch language {
	case lang.KZ:
		o := kz.ParseOpts(options...)

		return kz.Float64(decimal, o...)
	case lang.RU:
		o := ru.ParseOpts(options...)

		return ru.Float64(decimal, o...)
	case lang.EN:
		o := en.ParseOpts(options...)

		return en.Float64(decimal, o...)
	case lang.Unknown:
		return words, lang.ErrBadLanguage
	default:
		return words, lang.ErrBadLanguage
	}
}

// MustFloat64 returns float64 number converted to words or empty string on error.
func MustFloat64(decimal float64, language lang.Lang, options ...interface{}) string {
	res, err := Float64(decimal, language, options...)
	if err != nil {
		return ""
	}

	return res
}
