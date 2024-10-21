package ru

import (
	"github.com/dantedenis/numtow/curtow/cur"
)

type CurrencyOptions struct {
	currency          cur.Currency
	ignoreMinorUnits  bool
	convertMinorUnits bool
}

const (
	defaultCurrency = cur.KZT
)

type CurrencyOpt func(o *CurrencyOptions)

// WithCur sets cur.Currency
func WithCur(c cur.Currency) CurrencyOpt {
	return func(o *CurrencyOptions) {
		o.currency = c
	}
}

// WithCurConvMU sets flog to convert minor units to words.
func WithCurConvMU(convert bool) CurrencyOpt {
	return func(o *CurrencyOptions) {
		o.convertMinorUnits = convert
	}
}

// WithCurIgnoreMU sets flog to ignore minor units.
func WithCurIgnoreMU(ignore bool) CurrencyOpt {
	return func(o *CurrencyOptions) {
		o.ignoreMinorUnits = ignore
	}
}

func prepareCurrencyOptions(o ...CurrencyOpt) *CurrencyOptions {
	e := &CurrencyOptions{
		currency:         defaultCurrency,
		ignoreMinorUnits: false,
	}

	for _, opt := range o {
		opt(e)
	}

	return e
}

func ParseCurrencyOpts(options ...interface{}) []CurrencyOpt {
	o := make([]CurrencyOpt, 0)

	for _, v := range options {
		opt, ok := v.(CurrencyOpt)
		if ok {
			o = append(o, opt)
			continue
		}

		opts, ok := v.([]CurrencyOpt)
		if ok {
			o = opts
			break
		}
	}

	return o
}
