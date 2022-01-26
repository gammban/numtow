package numtow

import (
	"github.com/gammban/numtow/lang"
	"github.com/gammban/numtow/lang/en"
	"github.com/gammban/numtow/lang/kz"
	"github.com/gammban/numtow/lang/ru"
)

// String converts decimal number to words.
//  String("1", lang.EN)
func String(decimal string, language lang.Lang, options ...interface{}) (words string, err error) {
	switch language {
	case lang.KZ:
		o := kz.ParseOpts(options...)

		return kz.String(decimal, o...)
	case lang.RU:
		o := ru.ParseOpts(options...)

		return ru.String(decimal, o...)
	case lang.EN:
		o := en.ParseOpts(options...)

		return en.String(decimal, o...)
	case lang.Unknown:
		return words, lang.ErrBadLanguage
	default:
		return words, lang.ErrBadLanguage
	}
}

// MustString converts decimal number to words or returns an empty string on error.
func MustString(decimal string, language lang.Lang, options ...interface{}) string {
	res, err := String(decimal, language, options...)
	if err != nil {
		return ""
	}

	return res
}
