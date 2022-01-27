package kz

import (
	"fmt"
	"strings"

	"github.com/gammban/numtow/curtow/cur"
	"github.com/gammban/numtow/internal/ds"
)

func CurrencyString(amount string, o ...CurrencyOpt) (words string, err error) {
	e := prepareCurrencyOptions(o...)

	intDS, fracDS, err := ds.ParseDecimal(amount, ds.WithFracLen(uint(e.currency.MinorUnits())))
	if err != nil {
		return words, err
	}

	return convCurrency(intDS, fracDS, e.currency, e.ignoreMinorUnits, e.convertMinorUnits)
}

func CurrencyFloat64(amount float64, o ...CurrencyOpt) (words string, err error) {
	e := prepareCurrencyOptions(o...)

	intDS, fracDS, err := ds.ParseFloat64(amount, ds.WithFracLen(uint(e.currency.MinorUnits())))
	if err != nil {
		return words, err
	}

	return convCurrency(intDS, fracDS, e.currency, e.ignoreMinorUnits, e.convertMinorUnits)
}

func CurrencyInt64(amount int64, o ...CurrencyOpt) (words string, err error) {
	e := prepareCurrencyOptions(o...)

	e.ignoreMinorUnits = true

	return convCurrency(ds.ParseInt64(amount), ds.Zero, e.currency, e.ignoreMinorUnits, e.convertMinorUnits)
}

func convCurrency(intDS, fracDS ds.DigitString, c cur.Currency, hideMU, convMU bool) (words string, err error) {
	err = c.Validate()
	if err != nil {
		return words, err
	}

	curName, ok := curNamesKZ[c]
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

	sb := strings.Builder{}

	intWords, err := convert(intDS)
	if err != nil {
		return words, err
	}

	sb.WriteString(intWords)
	sb.WriteString(sep)
	sb.WriteString(curName)
	sb.WriteString(sep)

	if hideMU {
		return strings.TrimSpace(sb.String()), nil
	}

	minorUnitName, ok := curUnitNamesKZ[c]
	if !ok {
		return words, fmt.Errorf("%w: mu not implemented", cur.ErrBadCurrency)
	}

	fracWords := ""

	if convMU {
		fracWords, err = convert(fracDS)
		if err != nil {
			return words, err
		}
	} else {
		fracWords = fracDS.String()
	}

	sb.WriteString(fracWords)
	sb.WriteString(sep)
	sb.WriteString(minorUnitName)

	return strings.TrimSpace(sb.String()), nil
}
