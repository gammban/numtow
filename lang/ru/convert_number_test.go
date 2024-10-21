package ru

import (
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/dantedenis/numtow/internal/testdata"
	"github.com/dantedenis/numtow/lang/ru/gender"
)

func TestInt64_Female(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUNumbersInt64GenderFemale {
		got, err := Int64(k, WithFmtGender(gender.Female))
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(got, v) {
			t.Errorf("%d: \nexp: '%s'\ngot: '%s'", k, v, got)
			return
		}

		gotInt := MustInt64(k, WithFmtGender(gender.Female))
		if !strings.EqualFold(gotInt, v) {
			t.Errorf("%d: \nexp: '%s'\ngot: '%s'", k, v, gotInt)
			return
		}
	}
}

func BenchmarkRU_ConvertInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, v := range testdata.TestCaseLangRUNumbersInt64GenderFemale {
			got := MustInt64(k)
			if !strings.EqualFold(got, v) {
				b.Errorf("%d: \nexp: '%s'\ngot: '%s'", k, v, got)
				return
			}
		}
	}
}

func TestString_Male(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUNumbersStringGenderMale {
		got, err := String(k, WithFmtGender(gender.Male))
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(got, v) {
			t.Errorf("%s: \nexp: '%s'\ngot: '%s'", k, v, got)
			return
		}

		gotStr := MustString(k, WithFmtGender(gender.Male))
		if !strings.EqualFold(gotStr, v) {
			t.Errorf("%s: \nexp: '%s'\ngot: '%s'", k, v, gotStr)
			return
		}
	}
}

func TestString_Female(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUNumbersStringGenderFemale {
		got, err := String(k, WithFmtGender(gender.Female))
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(got, v) {
			t.Errorf("%s: \nexp: '%s'\ngot: '%s'", k, v, got)
			return
		}

		gotStr := MustString(k, WithFmtGender(gender.Female))
		if !strings.EqualFold(gotStr, v) {
			t.Errorf("%s: \nexp: '%s'\ngot: '%s'", k, v, gotStr)
			return
		}
	}
}

func TestString_Neuter(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUNumbersStringGenderNeuter {
		got, err := String(k, WithFmtGender(gender.Neuter))
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(got, v) {
			t.Errorf("%s: \nexp: '%s'\ngot: '%s'", k, v, got)
			return
		}

		gotStr := MustString(k, WithFmtGender(gender.Neuter))
		if !strings.EqualFold(gotStr, v) {
			t.Errorf("%s: \nexp: '%s'\ngot: '%s'", k, v, gotStr)
			return
		}
	}
}

func TestString_Errors(t *testing.T) {
	str := MustString("a")
	if str != "" {
		t.Fatal("expected empty string")
	}

	str = MustString("123", WithFmtGender(gender.Unknown))
	if str != "" {
		t.Fatal("expected empty string")
	}

	_, err := String("a")
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = String("123", WithFmtGender(gender.Unknown))
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestFloat64_Female(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUDecimalFloat64GenderFemale {
		give := k
		want := v

		s := strconv.FormatFloat(give, 'f', -1, 64)

		got, err := Float64(give, WithFmtGender(gender.Female))
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(got, want) {
			t.Errorf("%s:\nexp: '%s'\ngot: '%s'", s, want, got)
			return
		}
	}
}

func TestFloat64_Male(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUDecimalFloat64GenderMale {
		give := k
		want := v

		s := strconv.FormatFloat(give, 'f', -1, 64)

		got, err := Float64(give, WithFmtGender(gender.Male))
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(got, want) {
			t.Errorf("%s:\nexp: '%s'\ngot: '%s'", s, want, got)
			return
		}
	}
}

func TestFloat64_Neuter(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUDecimalFloat64GenderNeuter {
		give := k
		want := v

		s := strconv.FormatFloat(give, 'f', -1, 64)

		got, err := Float64(give, WithFmtGender(gender.Neuter))
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(got, want) {
			t.Errorf("%s:\nexp: '%s'\ngot: '%s'", s, want, got)
			return
		}
	}
}

func TestString_Decimal_Female(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUDecimalStringGenderFemale {
		give := k
		want := v

		got, err := String(give, WithFmtGender(gender.Female))
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(got, want) {
			t.Errorf("%s:\nexp: '%s'\ngot: '%s'", give, want, got)
			return
		}
	}
}

func TestString_Decimal_Male(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUDecimalStringGenderMale {
		give := k
		want := v

		got, err := String(give, WithFmtGender(gender.Male))
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(got, want) {
			t.Errorf("%s:\nexp: '%s'\ngot: '%s'", give, want, got)
			return
		}
	}
}

func TestString_Decimal_Neuter(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUDecimalStringGenderNeuter {
		give := k
		want := v

		got, err := String(give, WithFmtGender(gender.Neuter))
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(got, want) {
			t.Errorf("%s:\nexp: '%s'\ngot: '%s'", give, want, got)
			return
		}
	}
}

func TestString_Decimal_Errors(t *testing.T) {
	_, err := String("123", WithFmtGender(gender.Gender(10)))
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = String("aa", WithFmtGender(gender.Male))
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = Float64(math.NaN(), WithFmtGender(gender.Female))
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = Float64(1, WithFmtGender(gender.Gender(10)))
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestMustFloat64_Errors(t *testing.T) {
	if res := MustFloat64(math.NaN()); res != "" {
		t.Fatal("expected empty string")
	}

	if res := MustFloat64(math.Inf(-1)); res != "" {
		t.Fatal("expected empty string")
	}

	if res := MustFloat64(math.Inf(0)); res != "" {
		t.Fatal("expected empty string")
	}

	if res := MustFloat64(1); res != "одна" {
		t.Fatal("mismatch")
	}
}
