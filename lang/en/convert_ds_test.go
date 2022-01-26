package en

import (
	"errors"
	"strconv"
	"strings"
	"testing"

	"github.com/gammban/numtow/internal/digit"
	"github.com/gammban/numtow/internal/ds"
)

//nolint:gochecknoglobals
var testCaseConvertDS = []struct {
	GiveDS          ds.DigitString
	GiveFmtGroupSep string
	WantResult      string
	WantErr         error
}{
	{GiveDS: ds.New(0), WantResult: "zero", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(0, 0, 0, 0, 0, 0), WantResult: "zero", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(0), WantResult: "zero", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(10, 11, 12), WantResult: "", WantErr: digit.ErrBadDigit, GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.DigitString{DS: []digit.Digit{1}, IsSignMinus: true}, WantResult: "minus one", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1), WantResult: "one", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(2), WantResult: "two", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(3), WantResult: "three", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(4), WantResult: "four", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(5), WantResult: "five", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(6), WantResult: "six", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(7), WantResult: "seven", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(8), WantResult: "eight", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(9), WantResult: "nine", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0), WantResult: "ten", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(2, 0), WantResult: "twenty", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(0, 0, 0, 0, 0, 0, 0, 2), WantResult: "two", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0, 0), WantResult: "one hundred", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 1), WantResult: "eleven", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 2), WantResult: "twelve", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 3), WantResult: "thirteen", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 4), WantResult: "fourteen", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 5), WantResult: "fifteen", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 6), WantResult: "sixteen", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 7), WantResult: "seventeen", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 8), WantResult: "eighteen", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 9), WantResult: "nineteen", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(2, 2), WantResult: "twenty-two", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 2, 2), WantResult: "one hundred and twenty-two", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(9, 9, 9), WantResult: "nine hundred and ninety-nine", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(7, 0, 9), WantResult: "seven hundred and nine", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(7, 0, 0), WantResult: "seven hundred", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(6, 0, 0), WantResult: "six hundred", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(6, 7, 0), WantResult: "six hundred and seventy", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(0, 7, 0, 9), WantResult: "seven hundred and nine", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(0, 0, 7, 0, 9), WantResult: "seven hundred and nine", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(0, 0, 0, 7, 0, 9), WantResult: "seven hundred and nine", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(0, 0, 0, 0, 7, 0, 9), WantResult: "seven hundred and nine", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(5, 1, 2, 2), WantResult: "five thousand, one hundred and twenty-two", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(2, 5, 1, 3, 3), WantResult: "twenty-five thousand, one hundred and thirty-three", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(7, 2, 5, 1, 3, 3), WantResult: "seven hundred and twenty-five thousand, one hundred and thirty-three", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(2, 0, 5, 5, 1, 3, 3), WantResult: "two million, fifty-five thousand, one hundred and thirty-three", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(2, 0, 0, 0, 0, 0, 0), WantResult: "two million", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(4, 2, 0, 5, 5, 1, 3, 3), WantResult: "forty-two million, fifty-five thousand, one hundred and thirty-three", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(), WantResult: "", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2), WantResult: "", WantErr: strconv.ErrRange, GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 5, 1, 1, 3, 8, 8, 8, 7, 3, 0, 0, 9, 9), WantResult: "fifteen trillion, one hundred and thirteen billion, eight hundred and eighty-eight million, seven hundred and thirty thousand and ninety-nine", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 1, 2, 3, 9, 8, 0, 1, 9, 1, 4, 2, 5), WantResult: "one trillion, one hundred and twenty-three billion, nine hundred and eighty million, one hundred and ninety-one thousand, four hundred and twenty-five", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 9, 8, 9, 8, 0, 0, 6, 2, 3, 3, 3), WantResult: "one hundred and ninety-eight billion, nine hundred and eighty million, sixty-two thousand, three hundred and thirty-three", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 4, 3, 6, 0, 1, 2, 2, 0, 0, 8), WantResult: "fourteen billion, three hundred and sixty million, one hundred and twenty-two thousand and eight", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0, 2), WantResult: "one hundred and two", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 2, 0), WantResult: "one hundred and twenty", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0, 0, 2), WantResult: "one thousand and two", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 2, 0, 3), WantResult: "one thousand, two hundred and three", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0, 2, 0, 0, 3), WantResult: "one hundred and two thousand and three", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0, 2, 3, 0, 4), WantResult: "one hundred and two thousand, three hundred and four", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0, 0, 0, 0, 0, 2), WantResult: "one million and two", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0, 0, 0, 0, 2, 0), WantResult: "one million and twenty", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0, 0, 0, 2, 0, 0), WantResult: "one million, two hundred", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0, 0, 2, 0, 0, 0), WantResult: "one million, two thousand", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0, 0, 2, 0, 0, 3), WantResult: "one million, two thousand and three", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0, 2, 3, 0, 4, 5), WantResult: "one million, twenty-three thousand and forty-five", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 2, 0, 3, 4, 5, 0), WantResult: "one million, two hundred and three thousand, four hundred and fifty", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0, 0, 0, 0, 0, 3, 0, 0), WantResult: "one hundred million, three hundred", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0, 2, 0, 0, 0, 0, 0, 3), WantResult: "one hundred and two million and three", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(1, 0, 2, 3, 0, 4, 5, 6, 7), WantResult: "one hundred and two million, three hundred and four thousand, five hundred and sixty-seven", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(5, 1, 4, 1, 4, 3, 6, 0, 1, 2, 2, 0, 0, 8), WantResult: "fifty-one trillion, four hundred and fourteen billion, three hundred and sixty million, one hundred and twenty-two thousand and eight", GiveFmtGroupSep: defaultFormatGroupSeparator},
	{GiveDS: ds.New(5, 1, 4, 1, 4, 3, 6, 0, 1, 2, 2, 0, 0, 8), WantResult: "fifty-one trillion four hundred and fourteen billion three hundred and sixty million one hundred and twenty-two thousand and eight", GiveFmtGroupSep: ""},
	//	{GiveDS: ds.New(), WantResult: ""},

}

func TestConvert(t *testing.T) {
	for _, v := range testCaseConvertDS {
		v := v

		t.Run("", func(tt *testing.T) {
			gotRes, gotErr := convert(v.GiveDS,
				WithFmtGroupSep(v.GiveFmtGroupSep),
			)
			if !errors.Is(gotErr, v.WantErr) {
				tt.Fatalf("exp: '%s' \ngot: '%s'", v.WantErr, gotErr)
			}

			if !strings.EqualFold(v.WantResult, gotRes) {
				tt.Fatalf("%s\nexp: '%s'\ngot: '%s'", v.GiveDS.String(), v.WantResult, gotRes)
			}
		})
	}
}

func TestConvt(t *testing.T) {
	res, err := convertDecimal(ds.New(1, 5, 0, 1, 0, 1, 2), ds.New(1, 5), FormatWithoutAnd...)
	if err != nil {
		t.Fatal()
	}

	t.Log(res)
}

//nolint:gochecknoglobals
var testCaseConvertDecimal = []struct {
	GiveIntDS  ds.DigitString
	GiveFracDS ds.DigitString
	GiveOpts   []OptFunc
	WantResult string
	WantErr    error
}{
	{
		GiveIntDS:  ds.New(1),
		GiveFracDS: ds.New(1),
		WantResult: "one point one",
		GiveOpts:   FormatDefault,
	},
	{
		GiveIntDS:  ds.New(5, 1, 4, 1, 4, 3, 6, 0, 1, 2, 2, 0, 0, 8),
		GiveFracDS: ds.New(1),
		WantResult: "fifty-one trillion four hundred fourteen billion three hundred sixty million one hundred twenty-two thousand eight point one",
		GiveOpts:   FormatWithoutAnd,
	},
	{
		GiveIntDS:  ds.New(5, 1, 4, 1, 4, 3, 6, 0, 1, 2, 2, 0, 0, 8),
		GiveFracDS: ds.New(1),
		WantResult: "fifty-one trillion, four hundred and fourteen billion, three hundred and sixty million, one hundred and twenty-two thousand and eight point one",
		GiveOpts:   FormatDefault,
	},
	{
		GiveIntDS:  ds.New(0),
		GiveFracDS: ds.New(1),
		WantResult: "zero point one",
		GiveOpts:   FormatDefault,
	},
}

func TestConvertDecimal(t *testing.T) {
	for _, v := range testCaseConvertDecimal {
		v := v

		t.Run("", func(tt *testing.T) {
			gotRes, gotErr := convertDecimal(v.GiveIntDS, v.GiveFracDS,
				v.GiveOpts...,
			)
			if !errors.Is(gotErr, v.WantErr) {
				tt.Fatalf("exp: '%s' \ngot: '%s'", v.WantErr, gotErr)
			}

			if !strings.EqualFold(v.WantResult, gotRes) {
				tt.Fatalf("%s.%s\nexp: '%s'\ngot: '%s'", v.GiveIntDS.String(), v.GiveFracDS.String(), v.WantResult, gotRes)
			}
		})
	}
}
