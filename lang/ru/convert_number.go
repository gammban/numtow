package ru

import (
	"github.com/gammban/numtow/internal/ds"
)

// MustInt64 converts int64 number to russian words or returns empty string.
//  ru.MustInt64(1) // "одна"
//  ru.MustInt64(1, ru.WithFmtGender(gender.Male)) // "один"
//  ru.MustInt64(1, ru.WithFmtGender(gender.Neuter)) // "одно"
func MustInt64(number int64, options ...OptFunc) string {
	result, err := Int64(number, options...)
	if err != nil {
		return ""
	}

	return result
}

// Int64 converts int64 number to russian words or returns error.
//  result, err := ru.Int64(1) // result: "одна"
func Int64(number int64, options ...OptFunc) (string, error) {
	e := prepareOptions(options...)

	return convert(ds.ParseInt64(number), e.FmtGender)
}

// Float64 converts float64 number to russian words or returns error.
//  result, err := ru.Float64(2.54) // result: "две целых пятьдесят четыре сотых"
//  result, err := ru.Float64(1234567.12345) // result: "один миллион двести тридцать четыре тысячи пятьсот шестьдесят семь целых двенадцать тысяч триста сорок пять стотысячных"
func Float64(decimal float64, options ...OptFunc) (result string, err error) {
	o := prepareOptions(options...)

	intDS, fracDS, err := ds.ParseFloat64(decimal, ds.WithFracLen(o.ParseFracLen))
	if err != nil {
		return result, err
	}

	return convertDecimal(intDS, fracDS, options...)
}

// MustFloat64 converts float64 number to russian words or returns empty string.
//  result := ru.MustFloat64(2.54) // result: "две целых пятьдесят четыре сотых"
//  result := ru.MustFloat64(1234567.12345) // result: "один миллион двести тридцать четыре тысячи пятьсот шестьдесят семь целых двенадцать тысяч триста сорок пять стотысячных"
func MustFloat64(decimal float64, options ...OptFunc) string {
	result, err := Float64(decimal, options...)
	if err != nil {
		return ""
	}

	return result
}

// String converts string number to russian words or returns error.
//  result, err := ru.String("2.54") // result: "две целых пятьдесят четыре сотых"
//  result, err := ru.String("1234567.12345") // result: "один миллион двести тридцать четыре тысячи пятьсот шестьдесят семь целых двенадцать тысяч триста сорок пять стотысячных"
func String(decimal string, options ...OptFunc) (result string, err error) {
	o := prepareOptions(options...)

	intDS, fracDS, err := ds.ParseDecimal(decimal, ds.WithFracLen(o.ParseFracLen), ds.WithSep(o.ParseSeparator))
	if err != nil {
		return result, err
	}

	return convertDecimal(intDS, fracDS, options...)
}

// MustString converts string number to russian words or returns empty string.
//  result := ru.MustString("2.54") // result: "две целых пятьдесят четыре сотых"
//  result := ru.MustString("1234567.12345") // result: "один миллион двести тридцать четыре тысячи пятьсот шестьдесят семь целых двенадцать тысяч триста сорок пять стотысячных"
func MustString(decimal string, options ...OptFunc) string {
	result, err := String(decimal, options...)
	if err != nil {
		return ""
	}

	return result
}
