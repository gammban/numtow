package en

import (
	"errors"
	"math"
	"strings"
	"testing"

	"github.com/gammban/numtow/internal/ds"
)

//nolint:gochecknoglobals // test cases for decimal numbers
var (
	testCaseDecimal = []struct {
		GiveString  string
		GiveFloat64 float64
		GiveOpts    []OptFunc
		WantErr     error
		WantResult  string
	}{
		{GiveString: "0.1", GiveFloat64: 0.1, WantResult: "zero point one", GiveOpts: FormatDefault},
		{GiveString: ".1", GiveFloat64: 0.1, WantResult: "zero point one", GiveOpts: FormatDefault},
		{GiveString: "1", GiveFloat64: 1, WantResult: "one", GiveOpts: FormatDefault},
		{GiveString: "2", GiveFloat64: 2, WantResult: "two", GiveOpts: FormatDefault},
		{GiveString: "3", GiveFloat64: 3, WantResult: "three", GiveOpts: FormatDefault},
		{GiveString: "4", GiveFloat64: 4, WantResult: "four", GiveOpts: FormatDefault},
		{GiveString: "5", GiveFloat64: 5, WantResult: "five", GiveOpts: FormatDefault},
		{GiveString: "6", GiveFloat64: 6, WantResult: "six", GiveOpts: FormatDefault},
		{GiveString: "7", GiveFloat64: 7, WantResult: "seven", GiveOpts: FormatDefault},
		{GiveString: "8", GiveFloat64: 8, WantResult: "eight", GiveOpts: FormatDefault},
		{GiveString: "9", GiveFloat64: 9, WantResult: "nine", GiveOpts: FormatDefault},
		{GiveString: "10", GiveFloat64: 10, WantResult: "ten", GiveOpts: FormatDefault},
		{GiveString: "11", GiveFloat64: 11, WantResult: "eleven", GiveOpts: FormatDefault},
		{GiveString: "12", GiveFloat64: 12, WantResult: "twelve", GiveOpts: FormatDefault},
		{GiveString: "13", GiveFloat64: 13, WantResult: "thirteen", GiveOpts: FormatDefault},
		{GiveString: "14", GiveFloat64: 14, WantResult: "fourteen", GiveOpts: FormatDefault},
		{GiveString: "15", GiveFloat64: 15, WantResult: "fifteen", GiveOpts: FormatDefault},
		{GiveString: "16", GiveFloat64: 16, WantResult: "sixteen", GiveOpts: FormatDefault},
		{GiveString: "17", GiveFloat64: 17, WantResult: "seventeen", GiveOpts: FormatDefault},
		{GiveString: "18", GiveFloat64: 18, WantResult: "eighteen", GiveOpts: FormatDefault},
		{GiveString: "19", GiveFloat64: 19, WantResult: "nineteen", GiveOpts: FormatDefault},
		{GiveString: "20", GiveFloat64: 20, WantResult: "twenty", GiveOpts: FormatDefault},
		{GiveString: "21", GiveFloat64: 21, WantResult: "twenty-one", GiveOpts: FormatDefault},
		{GiveString: "33", GiveFloat64: 33, WantResult: "thirty-three", GiveOpts: FormatDefault},
		{GiveString: "-45", GiveFloat64: -45, WantResult: "minus forty-five", GiveOpts: FormatDefault},
		{GiveString: "-99.124", GiveFloat64: -99.124, WantResult: "minus ninety-nine point one hundred and twenty-four", GiveOpts: FormatDefault},
		{GiveString: "-99.124", GiveFloat64: -99.124, WantResult: "minus ninety-nine point one hundred twenty-four", GiveOpts: FormatWithoutAnd},
		{GiveString: "5234.924", GiveFloat64: 5234.924, WantResult: "five thousand two hundred thirty-four point nine hundred twenty-four", GiveOpts: FormatWithoutAnd},
		{GiveString: "--5234.924", GiveFloat64: math.NaN(), WantResult: "", WantErr: ds.ErrParse, GiveOpts: FormatWithoutAnd},
		{GiveString: "-5234.-924", GiveFloat64: math.Inf(0), WantResult: "", WantErr: ds.ErrParse, GiveOpts: FormatWithoutAnd},
		{GiveString: "5234.-924", GiveFloat64: math.Inf(-1), WantResult: "", WantErr: ds.ErrParse, GiveOpts: FormatWithoutAnd},
	}
)

func TestString(t *testing.T) {
	for _, v := range testCaseDecimal {
		gotRes, gotErr := String(v.GiveString, v.GiveOpts...)
		if !errors.Is(gotErr, v.WantErr) {
			t.Errorf("%s: \nexp err: %s\ngot err: %s", v.GiveString, v.WantErr, gotErr)
		}

		if gotErr == nil {
			if !strings.EqualFold(gotRes, v.WantResult) {
				t.Errorf("%s: \nexp: %s\ngot: %s", v.GiveString, v.WantResult, gotRes)
			}
		}
	}
}

func TestMustString(t *testing.T) {
	for _, v := range testCaseDecimal {
		gotRes := MustString(v.GiveString, v.GiveOpts...)
		if gotRes == "" && v.WantErr == nil {
			t.Errorf("%s: \nexp err: %s\ngot err: %s", v.GiveString, v.WantErr, gotRes)
		}

		if gotRes != "" {
			if !strings.EqualFold(gotRes, v.WantResult) {
				t.Errorf("%s: \nexp: %s\ngot: %s", v.GiveString, v.WantResult, gotRes)
			}
		}
	}
}

func TestFloat64(t *testing.T) {
	for _, v := range testCaseDecimal {
		gotRes, gotErr := Float64(v.GiveFloat64, v.GiveOpts...)
		if !errors.Is(gotErr, v.WantErr) {
			t.Errorf("%s: \nexp err: %s\ngot err: %s", v.GiveString, v.WantErr, gotErr)
		}

		if gotErr == nil {
			if !strings.EqualFold(gotRes, v.WantResult) {
				t.Errorf("%s: \nexp: %s\ngot: %s", v.GiveString, v.WantResult, gotRes)
			}
		}
	}
}

func TestMustFloat64(t *testing.T) {
	for _, v := range testCaseDecimal {
		gotRes := MustFloat64(v.GiveFloat64, v.GiveOpts...)
		if gotRes == "" && v.WantErr == nil {
			t.Errorf("%s: \nexp err: %s\ngot err: %s", v.GiveString, v.WantErr, gotRes)
		}

		if gotRes != "" {
			if !strings.EqualFold(gotRes, v.WantResult) {
				t.Errorf("%s: \nexp: %s\ngot: %s", v.GiveString, v.WantResult, gotRes)
			}
		}
	}
}

func TestInt64(t *testing.T) {
	res, err := Int64(6)
	if err != nil {
		t.Fatal(err)
	}

	if res != "six" {
		t.Fatal("mismatch")
	}
}

func TestMustInt64(t *testing.T) {
	if res := MustInt64(5); res != "five" {
		t.Fatal("mismatch")
	}
}
