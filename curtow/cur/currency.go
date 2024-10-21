package cur

import (
	"errors"
	"strings"
)

// Currency is a currency type
type Currency uint32

const (
	// Unknown currency
	Unknown Currency = iota
	// KZT - Kazakhstan Tenge
	KZT
	// USD - U.S. Dollar
	USD
	// RUB - Russian Ruble
	RUB
	// EUR - Euro
	EUR
	// GBP - Great Britain Pound
	GBP
)

const (
	CodeKZT = "KZT"
	CodeUSD = "USD"
	CodeRUB = "RUB"
	CodeEUR = "EUR"
	CodeGBP = "GBP"
)

var (
	ErrBadCurrency = errors.New("bad currency")
)

// String return code of currency
//
//	USD.String() // "USD"
func (c Currency) String() string {
	switch c {
	case KZT:
		return CodeKZT
	case RUB:
		return CodeRUB
	case USD:
		return CodeUSD
	case EUR:
		return CodeEUR
	case GBP:
		return CodeGBP
	case Unknown:
		return ""
	default:
		return ""
	}
}

// ParseCurrency returns Currency or Unknown.
func ParseCurrency(s string) Currency {
	switch strings.ToUpper(s) {
	case CodeKZT:
		return KZT
	case CodeUSD:
		return USD
	case CodeRUB:
		return RUB
	case CodeEUR:
		return EUR
	case CodeGBP:
		return GBP
	default:
		return Unknown
	}
}

// MinorUnits returns minor units of currency or MinorUnitsUnknown.
func (c Currency) MinorUnits() MinorUnits {
	d, ok := details[c]
	if ok {
		return d.MinorUnits
	}

	return MinorUnitsUnknown
}

// ISO4217 returns currency ISO4217 standard info or nil
func (c Currency) ISO4217() *ISO4217 {
	d, ok := details[c]
	if ok {
		return &d
	}

	return nil
}

// Validate validates currency. In case of invalid currency returns ErrBadCurrency.
func (c Currency) Validate() error {
	switch c {
	case KZT, RUB, USD, EUR, GBP:
		return nil
	case Unknown:
		return ErrBadCurrency
	default:
		return ErrBadCurrency
	}
}
