package en

import "github.com/dantedenis/numtow/internal/ds"

// Float64 converts float64 number to english words or returns error.
//
//	result, err := en.Float64(-99.124) // result: minus ninety-nine point one hundred and twenty-four
func Float64(decimal float64, options ...OptFunc) (result string, err error) {
	o := prepareOptions(options...)

	intDS, fracDS, err := ds.ParseFloat64(decimal, ds.WithFracLen(o.Parse.FracLen))
	if err != nil {
		return result, err
	}

	return convertDecimal(intDS, fracDS, options...)
}

// MustFloat64 converts float64 number to english words or returns empty string.
//
//	result := en.MustFloat64(-99.124) // result: minus ninety-nine point one hundred and twenty-four
func MustFloat64(decimal float64, options ...OptFunc) string {
	result, err := Float64(decimal, options...)
	if err != nil {
		return ""
	}

	return result
}

// String converts string number to english words or returns error.
//
//	result, err := en.String("21") // result: twenty-one
func String(decimal string, options ...OptFunc) (result string, err error) {
	o := prepareOptions(options...)

	intDS, fracDS, err := ds.ParseDecimal(decimal, ds.WithFracLen(o.Parse.FracLen), ds.WithSep(o.Parse.DecimalSeparator))
	if err != nil {
		return result, err
	}

	return convertDecimal(intDS, fracDS, options...)
}

// MustString converts string number to english words or returns empty string.
//
//	en.MustString("21") // result: "twenty-one"
//	en.MustString("bad") // result: ""
//	en.MustString("-99.124") // result: "minus ninety-nine point one hundred and twenty-four"
func MustString(decimal string, options ...OptFunc) string {
	result, err := String(decimal, options...)
	if err != nil {
		return ""
	}

	return result
}

// Int64 converts int64 number to english words or returns error.
//
//	result, err := en.Int64(1) // result: one
func Int64(num int64, options ...OptFunc) (result string, err error) {
	return convert(ds.ParseInt64(num))
}

// MustInt64 converts int64 number to english words or returns empty string.
//
//	result, err := en.MustInt64(1) // result: one
func MustInt64(num int64, options ...OptFunc) string {
	result, err := Int64(num, options...)
	if err != nil {
		return ""
	}

	return result
}
