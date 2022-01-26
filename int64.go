package numtow

import (
	"github.com/gammban/numtow/lang"
	"github.com/gammban/numtow/lang/en"
	"github.com/gammban/numtow/lang/kz"
	"github.com/gammban/numtow/lang/ru"
)

// Int64 converts number to words
func Int64(number int64, language lang.Lang, options ...interface{}) (words string, err error) {
	switch language {
	case lang.KZ:
		o := kz.ParseOpts(options...)

		return kz.Int64(number, o...)
	case lang.RU:
		o := ru.ParseOpts(options...)

		return ru.Int64(number, o...)
	case lang.EN:
		o := en.ParseOpts(options...)

		return en.Int64(number, o...)
	case lang.Unknown:
		return words, lang.ErrBadLanguage
	default:
		return words, lang.ErrBadLanguage
	}
}

// MustInt64 returns float64 number converted to words or empty string on error.
func MustInt64(number int64, language lang.Lang, options ...interface{}) string {
	res, err := Int64(number, language, options...)
	if err != nil {
		return ""
	}

	return res
}
