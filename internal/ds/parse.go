package ds

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gammban/numtow/internal/digit"
)

var (
	ErrParse = errors.New("parse number error")
)

const (
	base                    = 10
	bitSize                 = 64
	defaultExp              = 0
	defaultDecimalSeparator = '.'
	minusSign               = "-"
	float64ParseLimitMax    = 1000000000000
	float64ParseLimitMin    = -1000000000000
)

type ParseOpt struct {
	fracLen uint // Decimal fracLen
	sep     rune // Decimal separator
}

type ParseOptFunc func(o *ParseOpt)

func prepareParseOpt(options ...ParseOptFunc) *ParseOpt {
	o := &ParseOpt{
		fracLen: defaultExp,
		sep:     defaultDecimalSeparator,
	}

	for _, opt := range options {
		opt(o)
	}

	return o
}

func WithFracLen(fracLen uint) ParseOptFunc {
	return func(o *ParseOpt) {
		o.fracLen = fracLen
	}
}

// WithSep sets decimal separator
func WithSep(separator rune) ParseOptFunc {
	return func(o *ParseOpt) {
		o.sep = separator
	}
}

// ParseString parses string and returns a valid DigitString. In case when string is not a valid DigitString returns ErrParse.
func ParseString(s string) (ds DigitString, err error) {
	if s == "" {
		return ds, ErrParse
	}

	data := make([]digit.Digit, 0, len(s))
	isSignMinus := false

	for k, v := range s {
		if k == 0 && v == '-' {
			isSignMinus = true
			continue
		}

		d := digit.ParseRune(v)
		if d == digit.DigitUnknown {
			return ds, ErrParse
		}

		data = append(data, d)
	}

	ds.DS = data
	ds.IsSignMinus = isSignMinus

	return ds, nil
}

// ParseInt64 returns DigitString from int64
func ParseInt64(num int64) (ds DigitString) {
	s := strconv.FormatInt(num, base)

	data := make([]digit.Digit, 0, len(s))
	isSignMinus := false

	for k, v := range s {
		if k == 0 && v == '-' {
			isSignMinus = true
			continue
		}

		data = append(data, digit.ParseRune(v))
	}

	ds.DS = data
	ds.IsSignMinus = isSignMinus

	return ds
}

// ParseFloat64 returns integer and fractional parts.
// 	Parsing float64 limit is 1000000000000.
func ParseFloat64(number float64, options ...ParseOptFunc) (intDS, fracDS DigitString, err error) {
	if math.IsNaN(number) {
		return intDS, fracDS, fmt.Errorf("%w: float64 is NaN", ErrParse)
	}

	if math.IsInf(number, 0) {
		return intDS, fracDS, fmt.Errorf("%w: float64 is Inf", ErrParse)
	}

	if number >= float64ParseLimitMax {
		return intDS, fracDS, fmt.Errorf("%w: float64 has reached max parse limit", ErrParse)
	}

	if number <= float64ParseLimitMin {
		return intDS, fracDS, fmt.Errorf("%w: float64 has reached min parse limit", ErrParse)
	}

	sp := strconv.FormatFloat(number, 'f', -1, bitSize)

	return ParseDecimal(sp, options...)
}

// ParseDecimal returns integer and fractional parts
// 	Examples:
// 	ParseDecimal("123") // "123","00", nil
// 	ParseDecimal("123.00") // "123","00", nil
// 	ParseDecimal("-.45") // "-45","00", nil
// 	ParseDecimal(".87") // "0","87", nil
// 	ParseDecimal("-548725300538597.89") // "-548725300538597","89", nil
// 	ParseDecimal("-123.5", WithFracLen(2)) // "-123","50", nil
// 	ParseDecimal("548725300538597.89", WithFracLen(3)) // "548725300538597","89", nil
func ParseDecimal(number string, options ...ParseOptFunc) (intDS, fracDS DigitString, err error) {
	o := prepareParseOpt(options...)

	sep, exp := string(o.sep), o.fracLen

	var intDSr, fracDSr DigitString

	intPart, fracPart, err := splitDecimal(number, sep, exp)
	if err != nil {
		return intDS, fracDS, ErrParse
	}

	if exp != 0 && int(exp) > len(fracPart) {
		fracPart += strings.Repeat("0", int(exp)-len(fracPart))
	}

	if exp != 0 && int(exp) < len(fracPart) {
		fracPart = fracPart[:exp]
	}

	intDSr, err = ParseString(intPart)
	if err != nil {
		return intDS, fracDS, ErrParse
	}

	if fracPart != "" {
		fracDSr, err = ParseString(fracPart)
		if err != nil {
			return intDS, fracDS, ErrParse
		}

		if fracDSr.IsSignMinus {
			return intDS, fracDS, ErrParse
		}
	}

	intDS = intDSr
	fracDS = fracDSr

	return intDS, fracDS, nil
}

const (
	splitParts1 = 1
	splitParts2 = 2
)

func splitDecimal(number, sep string, fracLen uint) (intPart, fracPart string, err error) {
	ss := strings.Split(number, sep)
	switch len(ss) {
	case splitParts1:
		intPart = number
		fracPart = ""

		if fracLen != 0 {
			fracPart = strings.Repeat("0", int(fracLen))
		}
	case splitParts2:
		intPart = ss[0]
		fracPart = ss[1]

		if fracPart == "" {
			intPart = number
		}

		if fracPart != "" && intPart == "" {
			intPart = "0"
		}

		if fracPart != "" && intPart == minusSign {
			intPart = "-0"
		}
	default:
		return "", "", ErrParse
	}

	return intPart, fracPart, nil
}
