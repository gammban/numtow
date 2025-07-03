package en

import "github.com/gammban/numtow/curtow/cur"

type CurrencyOptions struct {
	currency          cur.Currency
	ignoreMinorUnits  bool
	convertMinorUnits bool
	ignoreAnd         bool
	groupSep          string
}

const (
	defaultCurrency = cur.KZT
)

type CurrencyOpt func(o *CurrencyOptions)

func WithIgnoreAnd(ignore bool) CurrencyOpt {
	return func(o *CurrencyOptions) {
		o.ignoreAnd = ignore
	}
}

func WithCur(c cur.Currency) CurrencyOpt {
	return func(o *CurrencyOptions) {
		o.currency = c
	}
}

func WithCurConvMU(convert bool) CurrencyOpt {
	return func(o *CurrencyOptions) {
		o.convertMinorUnits = convert
	}
}

func WithCurIgnoreMU(ignore bool) CurrencyOpt {
	return func(o *CurrencyOptions) {
		o.ignoreMinorUnits = ignore
	}
}

func WithCurGroupSep(sep string) CurrencyOpt {
	return func(o *CurrencyOptions) {
		o.groupSep = sep
	}
}

func prepareCurrencyOptions(o ...CurrencyOpt) *CurrencyOptions {
	e := &CurrencyOptions{
		currency:         defaultCurrency,
		ignoreMinorUnits: false,
		groupSep:         ",",
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
