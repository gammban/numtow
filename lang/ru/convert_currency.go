package ru

import (
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

	intDS, err := ds.ParseInt64(amount)
	if err != nil {
		return words, err
	}

	return convCurrency(intDS, ds.Zero, e.currency, e.ignoreMinorUnits, e.convertMinorUnits)
}

func convCurrency(intDS, fracDS ds.DigitString, c cur.Currency, hideMU, convMU bool) (words string, err error) {
	err = c.Validate()
	if err != nil {
		return words, err
	}

	err = intDS.Validate()
	if err != nil {
		return words, err
	}

	err = fracDS.Validate()
	if err != nil {
		return words, err
	}

	intFirstTriplet, err := intDS.FirstTriplet()
	if err != nil {
		return words, err
	}

	info, err := getCurrencyInfo(c)
	if err != nil {
		return words, err
	}

	intName := info.GetCurrencyName(getTripletDeclination(intFirstTriplet))

	intWords, err := convert(intDS, info.GetCurrencyGender())
	if err != nil {
		return words, err
	}

	sb := strings.Builder{}
	sb.WriteString(intWords)
	sb.WriteString(sep)
	sb.WriteString(intName)
	sb.WriteString(sep)

	if hideMU {
		return strings.TrimSpace(sb.String()), nil
	}

	fracFirstTriplet, err := fracDS.FirstTriplet()
	if err != nil {
		return words, err
	}

	unitName := info.GetCurrencyUnitName(getTripletDeclination(fracFirstTriplet))

	fracWords := ""

	if convMU {
		fracWords, err = convert(fracDS, info.GetCurrencyUnitGender())
		if err != nil {
			return words, err
		}
	} else {
		fracWords = fracDS.String()
	}

	sb.WriteString(fracWords)
	sb.WriteString(sep)
	sb.WriteString(unitName)

	return strings.TrimSpace(sb.String()), nil
}
