package kz

import (
	"github.com/dantedenis/numtow/internal/ds"
)

// Float64 converts float64 to words or returns error.
//
//	result, err := Float64(84.13) // result: сексен төрт бүтін жүзден он үш
func Float64(decimal float64, options ...OptFunc) (result string, err error) {
	o := prepareOptions(options...)

	intDS, fracDS, err := ds.ParseFloat64(decimal, ds.WithFracLen(o.ParseFracLen))
	if err != nil {
		return result, err
	}

	return convertDecimal(intDS, fracDS, options...)
}

// MustFloat64 converts float64 to words or returns empty string.
//
//	result := MustFloat64(84.13) // result: сексен төрт бүтін жүзден он үш
func MustFloat64(decimal float64, options ...OptFunc) string {
	result, err := Float64(decimal, options...)
	if err != nil {
		return ""
	}

	return result
}

// String converts string decimal to words or returns error.
//
//	result, err := String("1") // result: бiр
func String(decimal string, options ...OptFunc) (result string, err error) {
	o := prepareOptions(options...)

	intDS, fracDS, err := ds.ParseDecimal(decimal, ds.WithFracLen(o.ParseFracLen), ds.WithSep(o.ParseSeparator))
	if err != nil {
		return result, err
	}

	return convertDecimal(intDS, fracDS, options...)
}

// MustString converts string decimal to words or returns empty string.
//
//	result := MustString("1") // result: бiр
func MustString(decimal string, options ...OptFunc) string {
	result, err := String(decimal, options...)
	if err != nil {
		return ""
	}

	return result
}

// Int64 converts int64 to words or returns error.
//
//	result, err := Int64(1) // result: бiр
func Int64(number int64, options ...OptFunc) (result string, err error) {
	return convert(ds.ParseInt64(number))
}

// MustInt64 converts int64 to words or returns empty string.
//
//	result, err := MustInt64(1) // result: бiр
func MustInt64(number int64, options ...OptFunc) string {
	result, err := Int64(number, options...)
	if err != nil {
		return ""
	}

	return result
}
