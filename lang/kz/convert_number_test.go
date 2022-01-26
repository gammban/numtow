package kz

import (
	"math"
	"strings"
	"testing"

	"github.com/gammban/numtow/internal/testdata"
)

func TestFloat64_Opts(t *testing.T) {
	res, err := Float64(1.01, WithFmtFracIgnore(true))
	if err != nil {
		t.Fatal(err)
	}

	if res != "бір" {
		t.Fatal("mismatch")
	}

	res, err = Float64(-1.01, WithFmtFracIgnore(true))
	if err != nil {
		t.Fatal(err)
	}

	if res != "минус бір" {
		t.Fatal("mismatch")
	}

	res, err = Float64(-100.974, WithFmtFracIgnore(false), WithFmtFracUseDigits(true))
	if err != nil {
		t.Fatal(err)
	}

	if res != "минус жүз бүтін мыңнан 974" {
		t.Fatal("mismatch")
	}

	res, err = Float64(-0.974, WithFmtFracIgnore(false), WithFmtFracUseDigits(true))
	if err != nil {
		t.Fatal(err)
	}

	if res != "минус нөл бүтін мыңнан 974" {
		t.Fatal("mismatch")
	}
}

func TestFloat64(t *testing.T) {
	for k, v := range testdata.TestCaseLangKZDecimalFloat64 {
		got, err := Float64(k)
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(got, v) {
			t.Errorf("%f: expected '%s' got '%s'", k, v, got)
			return
		}
	}
}

func TestString(t *testing.T) {
	for k, v := range testdata.TestCaseLangKZDecimalString {
		got, err := String(k)
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(got, v) {
			t.Errorf("%s: \nexp: '%s'\ngot: '%s'", k, v, got)
			return
		}
	}
}

func TestInt64(t *testing.T) {
	for k, v := range testdata.TestCaseLangKZNumbersInt64 {
		got, err := Int64(k)
		if err != nil {
			t.Error(err)
			return
		}

		if got != v {
			t.Errorf("%d: expected '%s' got '%s'", k, v, got)
			return
		}
	}
}

func TestInt64_9999(t *testing.T) {
	for k, v := range testdata.TestCaseKZWords9999 {
		got, err := Int64(k)
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(got, v) {
			t.Errorf("%d: expected '%s' got '%s'", k, v, got)
			return
		}
	}
}

func BenchmarkKZ_ConvertInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, v := range testdata.TestCaseLangKZNumbersInt64 {
			got, err := Int64(k)
			if err != nil {
				b.Error(err)
				return
			}

			if !strings.EqualFold(got, v) {
				b.Errorf("%d: \nexp: '%s'\ngot: '%s'", k, v, got)
				return
			}
		}
	}
}

func TestMustFloat64(t *testing.T) {
	if res := MustFloat64(math.NaN()); res != "" {
		t.Fatal("result mismatch, expected empty string")
	}

	if res := MustFloat64(math.Inf(0)); res != "" {
		t.Fatal("result mismatch, expected empty string")
	}

	if res := MustFloat64(math.Inf(-1)); res != "" {
		t.Fatal("result mismatch, expected empty string")
	}

	if res := MustFloat64(2); res != "екі" {
		t.Fatal("result mismatch, expected бір")
	}
}

func TestMustString(t *testing.T) {
	if res := MustString("50"); res != "елу" {
		t.Fatal("result mismatch, expected бір")
	}

	if res := MustString("bad"); res != "" {
		t.Fatal("result mismatch, expected empty string")
	}
}

func TestMustInt64(t *testing.T) {
	if res := MustInt64(5); res != "бес" {
		t.Fatal("result mismatch, expected бір")
	}
}
