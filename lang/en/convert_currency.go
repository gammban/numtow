package en

import (
	"fmt"
	"strings"

	"github.com/gammban/numtow/curtow/cur"
	"github.com/gammban/numtow/internal/ds"
)

// CurrencyString converts string currency to words or returns error.
//
//	en.CurrencyString("1.1", en.WithCur(cur.USD), en.WithCurConvMU(true)) // result: "one dollar and ten cents"
func CurrencyString(amount string, o ...CurrencyOpt) (words string, err error) {
	e := prepareCurrencyOptions(o...)

	intDS, fracDS, err := ds.ParseDecimal(amount, ds.WithFracLen(uint(e.currency.MinorUnits())))
	if err != nil {
		return words, err
	}

	return convCurrency(intDS, fracDS, o...)
}

// CurrencyFloat64 converts string currency to words or returns error.
//
//	en.CurrencyFloat64(1.1, en.WithCur(cur.USD), en.WithCurConvMU(true)) // result: "one dollar and ten cents"
func CurrencyFloat64(amount float64, o ...CurrencyOpt) (words string, err error) {
	e := prepareCurrencyOptions(o...)

	intDS, fracDS, err := ds.ParseFloat64(amount, ds.WithFracLen(uint(e.currency.MinorUnits())))
	if err != nil {
		return words, err
	}

	return convCurrency(intDS, fracDS, o...)
}

func convCurrency(intDS, fracDS ds.DigitString, opts ...CurrencyOpt) (words string, err error) {
	o := prepareCurrencyOptions(opts...)

	err = o.currency.Validate()
	if err != nil {
		return words, err
	}

	c, ok := curNamesEN[o.currency]
	if !ok {
		return words, fmt.Errorf("%w: not implemented", cur.ErrBadCurrency)
	}

	err = intDS.Validate()
	if err != nil {
		return words, err
	}

	err = fracDS.Validate()
	if err != nil {
		return words, err
	}

	if fracDS.IsSignMinus {
		return words, fmt.Errorf("%w: fractional part must be positive", ds.ErrParse)
	}

	sb := strings.Builder{}

	var cOpts []OptFunc
	if o.ignoreAnd {
		cOpts = append(cOpts, WithFmtAndSep(""))
	}

	cOpts = append(cOpts, WithFmtGroupSep(o.groupSep))

	intPartWords, err := convert(intDS, cOpts...)
	if err != nil {
		return words, err
	}

	sb.WriteString(intPartWords)
	sb.WriteString(sep)

	if intDS.String() == "1" {
		sb.WriteString(c.Singular)
	} else {
		sb.WriteString(c.Plural)
	}

	sb.WriteString(sep)

	if o.ignoreMinorUnits {
		return strings.TrimSpace(sb.String()), nil
	}

	var fracWords string

	if o.convertMinorUnits {
		fracWords, err = convert(fracDS)
		if err != nil {
			return words, err
		}
	} else {
		fracWords = fracDS.String()
	}

	if !o.ignoreAnd {
		sb.WriteString(and)
		sb.WriteString(sep)
	}

	sb.WriteString(fracWords)
	sb.WriteString(sep)

	if fracDS.String() == "1" {
		sb.WriteString(c.UnitSingular)
	} else {
		sb.WriteString(c.UnitPlural)
	}

	return strings.TrimSpace(sb.String()), nil
}
