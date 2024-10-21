package en

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dantedenis/numtow/internal/digit"
	"github.com/dantedenis/numtow/internal/ds"
	"github.com/dantedenis/numtow/internal/triplet"
)

// convert DigitalString to english words
//
//	Examples:
//	ds.Empty -> ""
//	ds.Zero -> "zero"
//	ds.New(1) -> "one"
//	ds.New(100) -> "one hundred"
//	ds.New(0, 0, 0, 0, 0, 2) -> "two"
//
//nolint:gocyclo
func convert(d ds.DigitString, options ...OptFunc) (result string, err error) {
	o := prepareOptions(options...)

	if d.IsEmpty() {
		return "", nil
	}

	if d.IsZero() {
		return zero, nil
	}

	sb := strings.Builder{}

	if d.IsSignMinus {
		sb.WriteString(o.Format.MinusSignWord)
		sb.WriteString(sep)
	}

	triplets, err := d.Split()
	if err != nil {
		return result, err
	}

	if triplets.Len() > len(megs) {
		return result, strconv.ErrRange
	}

	for i := triplets.Len() - 1; i >= 0; i-- {
		t := triplets[i]

		if t.IsZero() {
			continue
		}

		isTensAndUnitsZero := t.Tens().IsZero() && t.Units().IsZero()

		if t.Hundreds().IsNotZero() {
			sb.WriteString(units[t.Hundreds()])
			sb.WriteString(sep)
			sb.WriteString(hundred)
			sb.WriteString(sep)
		}

		// разделительное and между тысяными: 108 -> one hundred and eight
		if o.Format.AndWord != "" && t.Hundreds().IsNotZero() && !isTensAndUnitsZero {
			sb.WriteString(o.Format.AndWord)
			sb.WriteString(sep)
		}

		if !isTensAndUnitsZero {
			switch t.Tens() {
			case digit.Digit0:
				sb.WriteString(units[t.Units()])
				sb.WriteString(sep)
			case digit.Digit1:
				sb.WriteString(teens[t.Units()])
				sb.WriteString(sep)
			case digit.Digit2, digit.Digit3, digit.Digit4, digit.Digit5, digit.Digit6, digit.Digit7, digit.Digit8, digit.Digit9, digit.DigitUnknown:
				if t.Units().IsNotZero() {
					sb.WriteString(tens[t.Tens()])
					sb.WriteString("-")
					sb.WriteString(units[t.Units()])
					sb.WriteString(sep)

					break
				}

				sb.WriteString(tens[t.Tens()])
				sb.WriteString(sep)
			}
		}

		sb.WriteString(megs[i])

		if groupSep := getGroupSep(triplets, i, o); groupSep != "" {
			sb.WriteString(groupSep)
		}

		sb.WriteString(sep)
	}

	res := strings.TrimSpace(sb.String())

	if o.Format.GroupSeparator != "" && strings.HasSuffix(res, o.Format.GroupSeparator) {
		res = res[:len(res)-1]
	}

	return res, nil
}

func getGroupSep(triplets triplet.Triplets, i int, o *Options) string {
	if i == 0 {
		return ""
	}

	useAnd := true

	if i-1 == 0 {
		useAnd = false
	}

	for j := i - 1; j > 0; j-- {
		if useAnd && j != 0 {
			useAnd = triplets[j].IsZero()
		}
	}

	if i == 1 && !triplets[0].IsZero() && triplets[0].Hundreds().IsZero() {
		if o.Format.AndWord == "" {
			return ""
		}

		return sep + o.Format.AndWord // and
	}

	if useAnd && !triplets[0].IsZero() && triplets[0].Hundreds().IsZero() {
		if o.Format.AndWord == "" {
			return ""
		}

		return sep + o.Format.AndWord // and
	}

	return o.Format.GroupSeparator
}

func convertDecimal(intDS, fracDS ds.DigitString, options ...OptFunc) (result string, err error) {
	o := prepareOptions(options...)

	if fracDS.IsSignMinus {
		return result, fmt.Errorf("%w: fractional part must be positive", ds.ErrParse)
	}

	intPartWords, err := convert(intDS, options...)
	if err != nil {
		return result, err
	}

	if o.Format.FracIgnore {
		return intPartWords, nil
	}

	fracPartWords := ""
	if !fracDS.IsEmpty() && !fracDS.IsZero() {
		fracPartWords, err = convert(fracDS, options...)
		if err != nil {
			return result, err
		}
	}

	if fracPartWords == "" {
		return intPartWords, nil
	}

	sb := strings.Builder{}

	// если целая часть -0, ир запишем "минус ", потому что ожидается дробная часть
	if intDS.IsZero() && intDS.IsSignMinus {
		sb.WriteString(o.Format.MinusSignWord)
		sb.WriteString(sep)
	}

	sb.WriteString(intPartWords)
	sb.WriteString(sep)
	sb.WriteString(integerPart)
	sb.WriteString(sep)

	if o.Format.FracUseDigits {
		sb.WriteString(fracDS.String())
	} else {
		sb.WriteString(fracPartWords)
	}

	return sb.String(), nil
}
