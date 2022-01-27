package en

import "github.com/gammban/numtow/internal/ds"

func Float64(decimal float64, options ...OptFunc) (result string, err error) {
	o := prepareOptions(options...)

	intDS, fracDS, err := ds.ParseFloat64(decimal, ds.WithFracLen(o.Parse.FracLen))
	if err != nil {
		return result, err
	}

	return convertDecimal(intDS, fracDS, options...)
}

func MustFloat64(decimal float64, options ...OptFunc) string {
	result, err := Float64(decimal, options...)
	if err != nil {
		return ""
	}

	return result
}

func String(decimal string, options ...OptFunc) (result string, err error) {
	o := prepareOptions(options...)

	intDS, fracDS, err := ds.ParseDecimal(decimal, ds.WithFracLen(o.Parse.FracLen), ds.WithSep(o.Parse.DecimalSeparator))
	if err != nil {
		return result, err
	}

	return convertDecimal(intDS, fracDS, options...)
}

func MustString(decimal string, options ...OptFunc) string {
	result, err := String(decimal, options...)
	if err != nil {
		return ""
	}

	return result
}

func Int64(num int64, options ...OptFunc) (result string, err error) {
	return convert(ds.ParseInt64(num))
}

func MustInt64(num int64, options ...OptFunc) string {
	result, err := Int64(num, options...)
	if err != nil {
		return ""
	}

	return result
}
