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
	// CHF - Swiss Franc
	CHF
	// CAD - Canadian dollar
	CAD
	// JPY - Japanese Yen
	JPY
	// CNY - Chinese Yuan
	CNY
	// AUD - Australian dollar
	AUD
	// AED - United Arab Emirates Dirham
	AED
	// TRY - Turkish lira
	TRY
	// KGS - Kyrgyz som
	KGS
	// ZAR - South African rand
	ZAR
	// NOK -Norsk krone
	NOK
	// SEK - Swedish Krona
	SEK
)

const (
	CodeKZT = "KZT"
	CodeUSD = "USD"
	CodeRUB = "RUB"
	CodeEUR = "EUR"
	CodeGBP = "GBP"
	CodeCHF = "CHF"
	CodeCAD = "CAD"
	CodeJPY = "JPY"
	CodeCNY = "CNY"
	CodeAUD = "AUD"
	CodeAED = "AED"
	CodeTRY = "TRY"
	CodeKGS = "KGS"
	CodeZAR = "ZAR"
	CodeNOK = "NOK"
	CodeSEK = "SEK"
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
	case CHF:
		return CodeCHF
	case CAD:
		return CodeCAD
	case JPY:
		return CodeJPY
	case CNY:
		return CodeCNY
	case AUD:
		return CodeAUD
	case AED:
		return CodeAED
	case TRY:
		return CodeTRY
	case KGS:
		return CodeKGS
	case ZAR:
		return CodeZAR
	case NOK:
		return CodeNOK
	case SEK:
		return CodeSEK
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
	case CodeCHF:
		return CHF
	case CodeCAD:
		return CAD
	case CodeJPY:
		return JPY
	case CodeCNY:
		return CNY
	case CodeAUD:
		return AUD
	case CodeAED:
		return AED
	case CodeTRY:
		return TRY
	case CodeKGS:
		return KGS
	case CodeZAR:
		return ZAR
	case CodeNOK:
		return NOK
	case CodeSEK:
		return SEK
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
	case KZT, RUB, USD, EUR, GBP, CHF, CAD, JPY, CNY, AUD, AED, TRY, KGS, ZAR, NOK, SEK:
		return nil
	case Unknown:
		return ErrBadCurrency
	default:
		return ErrBadCurrency
	}
}
