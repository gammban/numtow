package kz

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gammban/numtow/internal/digit"
	"github.com/gammban/numtow/internal/ds"
)

// convert returns DigitString in kazakh
//
//	examples:
//	"" -> ""
//	"0" -> "нөл"
//	"-0" -> "нөл"
//	"-1" -> "минус бір"
//	"1" -> "бір"
//
//nolint:gocyclo
func convert(d ds.DigitString) (result string, err error) {
	if d.IsEmpty() {
		return "", nil
	}

	if d.IsZero() {
		return zero, nil
	}

	s := strings.Builder{}

	if d.IsSignMinus {
		s.WriteString(minus)
		s.WriteString(sep)
	}

	triplets, err := d.Split()
	if err != nil {
		return result, err
	}

	if triplets.Len() > len(megs) {
		return result, strconv.ErrRange
	}

	for i := len(triplets) - 1; i >= 0; i-- {
		t := triplets[i]

		if t.IsZero() {
			continue
		}

		switch t.Hundreds() {
		case digit.Digit0, digit.DigitUnknown:
		case digit.Digit1:
			s.WriteString(hundred)
			s.WriteString(sep)
		case digit.Digit2, digit.Digit3, digit.Digit4, digit.Digit5, digit.Digit6, digit.Digit7, digit.Digit8, digit.Digit9:
			s.WriteString(units[t.Hundreds()])
			s.WriteString(sep)
			s.WriteString(hundred)
			s.WriteString(sep)
		}

		if !(t.Tens().IsZero() && t.Units().IsZero()) {
			switch t.Tens() {
			case digit.Digit0, digit.DigitUnknown:
				s.WriteString(units[t.Units()])
				s.WriteString(sep)
			case digit.Digit1, digit.Digit2, digit.Digit3, digit.Digit4, digit.Digit5, digit.Digit6, digit.Digit7, digit.Digit8, digit.Digit9:
				if t.Units().IsNotZero() {
					s.WriteString(tens[t.Tens()])
					s.WriteString(sep)
					s.WriteString(units[t.Units()])
					s.WriteString(sep)

					break
				}

				s.WriteString(tens[t.Tens()])
				s.WriteString(sep)
			}
		}

		s.WriteString(megs[i])
		s.WriteString(sep)
	}

	return strings.TrimSpace(s.String()), nil
}

func convertDecimal(intDS, fracDS ds.DigitString, options ...OptFunc) (result string, err error) {
	o := prepareOptions(options...)

	if fracDS.IsSignMinus {
		return result, fmt.Errorf("%w: fractional part must be positive", ds.ErrParse)
	}

	intPartWords, err := convert(intDS)
	if err != nil {
		return result, err
	}

	if o.FmtFracIgnore {
		return intPartWords, nil
	}

	fracPartWords := ""
	if !fracDS.IsEmpty() && !fracDS.IsZero() {
		fracPartWords, err = convert(fracDS)
		if err != nil {
			return result, err
		}
	}

	sb := strings.Builder{}

	if fracPartWords == "" {
		return intPartWords, nil
	}

	// если целая часть -0, ир запишем "минус ", потому что ожидается дробная часть
	if intDS.IsZero() && intDS.IsSignMinus {
		sb.WriteString(minus)
		sb.WriteString(sep)
	}

	sb.WriteString(intPartWords)
	sb.WriteString(sep)
	sb.WriteString(integerPart)
	sb.WriteString(sep)

	fracPartName, err := getFracPartName(fracDS)
	if err != nil {
		return result, err
	}

	if o.FmtFracUseDigits {
		sb.WriteString(fracPartName)
		sb.WriteString(sep)
		sb.WriteString(fracDS.String())
	} else {
		sb.WriteString(fracPartName)
		sb.WriteString(sep)
		sb.WriteString(fracPartWords)
	}

	return sb.String(), nil
}

// 1 -> оннан
// 10 -> жүзден
// 500 -> мыңнан
// 1000 -> миллионнан
func getFracPartName(fracDS ds.DigitString) (fracPartName string, err error) {
	if fracDS.Len() > len(fracPart) {
		return fracPartName, fmt.Errorf("%w: fractional part range error", ds.ErrParse)
	}

	fracPartName = fracPart[fracDS.Len()]

	if fracPartName == "" {
		return fracPartName, fmt.Errorf("%w: cannot detec fractional part name", ds.ErrParse)
	}

	return fracPartName, nil
}
