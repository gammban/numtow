package ru

import (
	"fmt"
	"strings"

	"github.com/gammban/numtow/internal/digit"
	"github.com/gammban/numtow/internal/ds"
	"github.com/gammban/numtow/lang/ru/gender"
)

const (
	tripletIdx1 = 1
	tripletIdx2 = 2
)

//nolint:gocognit,gocyclo
func convert(d ds.DigitString, g gender.Gender) (res string, err error) {
	var un, tn, h digit.Digit

	err = g.Validate()
	if err != nil {
		return res, err
	}

	if d.IsEmpty() {
		return "", nil
	}

	if d.IsZero() {
		return zero, nil
	}

	sb := strings.Builder{}

	if d.IsSignMinus {
		sb.WriteString(minus)

		sb.WriteString(sep)
	}

	triplets, err := d.Split()
	if err != nil {
		return res, err
	}

	if triplets.IsZero() {
		return zero, nil
	}

	for i := len(triplets); i >= 1; i-- {
		t := triplets[i-1]

		if t.IsZero() {
			continue
		}

		h = t.Hundreds()
		tn = t.Tens()
		un = t.Units()

		//-- добавляем сотни
		if h.IsNotZero() {
			sb.WriteString(hundreds[h])

			sb.WriteString(sep)
		}

		// добавляем десятки
		if tn.IsNotZero() {
			if tn == digit.Digit1 && un.IsNotZero() {
				sb.WriteString(teens[un])

				sb.WriteString(sep)

				un = digit.Digit0
			} else {
				sb.WriteString(tens[tn])

				sb.WriteString(sep)
			}
		}

		// добавляем единицы:
		if un.IsNotZero() {
			switch i {
			// до 1 000 - смотрим указанный род
			case tripletIdx1:
				unitWord, err := getUnitsByGender(g, un)
				if err != nil {
					return res, err
				}

				if unitWord != "" {
					sb.WriteString(unitWord)
					sb.WriteString(sep)
				}
			// от 1 000 до 1 000 000 - женский род
			case tripletIdx2:
				sb.WriteString(unitsFemale[un])
				sb.WriteString(sep)
			// свыше 1 000 000 - мужской род
			default:
				sb.WriteString(unitsMale[un])
				sb.WriteString(sep)
			}
		}

		// добавляем наименование триады:
		if i > 1 {
			// полное наименование с учётом склонения:
			megaName := getMegaByDeclination(getTripletDeclination(t), i-1)

			sb.WriteString(megaName)
			sb.WriteString(sep)
		}
	}

	return strings.TrimSpace(sb.String()), nil
}

// Существительным управляет дробная часть числительного
func convertDecimal(intDS, fracDS ds.DigitString, options ...OptFunc) (result string, err error) {
	e := prepareOptions(options...)

	g := e.FmtGender

	if fracDS.IsSignMinus {
		return result, fmt.Errorf("%w: fractional part must be positive", ds.ErrParse)
	}

	intPart, err := convert(intDS, g)
	if err != nil {
		return result, err
	}

	if e.FmtFracIgnore {
		return intPart, nil
	}

	fracPart := ""
	if !fracDS.IsEmpty() && !fracDS.IsZero() {
		// дробная часть всегда женского рода
		fracPart, err = convert(fracDS, gender.Female)
		if err != nil {
			return result, err
		}
	}

	sb := strings.Builder{}

	if fracPart == "" {
		return intPart, nil
	}

	if intDS.IsZero() && intDS.IsSignMinus {
		sb.WriteString(minus)
		sb.WriteString(sep)
	}

	sb.WriteString(intPart)
	sb.WriteString(sep)

	intFirstTriplet, err := intDS.FirstTriplet()
	if err != nil {
		return result, err
	}

	integerPartName := getIntPart(getTripletDeclination(intFirstTriplet), g)
	if integerPartName == "" {
		return result, fmt.Errorf("%w: inetegr part name error", ds.ErrParse)
	}

	sb.WriteString(integerPartName)
	sb.WriteString(sep)

	fracFirstTriplet, err := fracDS.FirstTriplet()
	if err != nil {
		return result, err
	}

	// дробная часть всегда женского рода
	fracPartName := getFracPart(getTripletDeclination(fracFirstTriplet), gender.Female, fracDS.Len())
	if fracPartName == "" {
		return result, fmt.Errorf("%w: fractional part name error", ds.ErrParse)
	}

	if e.FmtFracUseDigits {
		sb.WriteString(fracDS.String())
		sb.WriteString(sep)
		sb.WriteString(fracPartName)
	} else {
		sb.WriteString(fracPart)
		sb.WriteString(sep)
		sb.WriteString(fracPartName)
	}

	return sb.String(), nil
}
