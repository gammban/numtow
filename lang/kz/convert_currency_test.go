package kz

import (
	"math"
	"testing"

	"github.com/gammban/numtow/internal/digit"
	"github.com/gammban/numtow/internal/ds"

	"github.com/gammban/numtow/curtow/cur"
	"github.com/gammban/numtow/internal/testdata"
)

func TestCurrencyInt64_KZT(t *testing.T) {
	for k, v := range testdata.TestCaseLangKZCurrencyKZTInt {
		got, err := CurrencyInt64(k, WithCur(cur.KZT))
		if err != nil {
			t.Error(err)
			return
		}

		if got != v {
			t.Errorf("%d: \nexp: '%s' \ngot: '%s'", k, v, got)
			return
		}
	}
}

func TestCurrencyInt64_USD(t *testing.T) {
	for k, v := range testdata.TestCaseLangKZCurrencyUSDInt {
		got, err := CurrencyInt64(k, WithCur(cur.USD))
		if err != nil {
			t.Error(err)
			return
		}

		if got != v {
			t.Errorf("%d: \nexp: '%s' \ngot: '%s'", k, v, got)
			return
		}
	}
}

func TestCurrencyInt64_EUR(t *testing.T) {
	for k, v := range testdata.TestCaseLangKZCurrencyEURInt {
		got, err := CurrencyInt64(k, WithCur(cur.EUR))
		if err != nil {
			t.Error(err)
			return
		}

		if got != v {
			t.Errorf("%d: \nexp: '%s' \ngot: '%s'", k, v, got)
			return
		}
	}
}

func TestCurrencyString_KZT(t *testing.T) {
	for k, v := range testdata.TestCaseLangKZCurrencyKZTString {
		got, err := CurrencyString(k, WithCur(cur.KZT))
		if err != nil {
			t.Error(err)
			return
		}

		if got != v {
			t.Errorf("%s: \nexp: '%s' \ngot: '%s'", k, v, got)
			return
		}
	}
}

func TestCurrencyString_USD(t *testing.T) {
	for k, v := range testdata.TestCaseLangKZCurrencyUSDString {
		got, err := CurrencyString(k, WithCur(cur.USD))
		if err != nil {
			t.Error(err)
			return
		}

		if got != v {
			t.Errorf("%s: \nexp: '%s' \ngot: '%s'", k, v, got)
			return
		}
	}
}

func TestCurrencyString_EUR(t *testing.T) {
	for k, v := range testdata.TestCaseLangKZCurrencyEURString {
		got, err := CurrencyString(k, WithCur(cur.EUR))
		if err != nil {
			t.Error(err)
			return
		}

		if got != v {
			t.Errorf("%s: \nexp: '%s' \ngot: '%s'", k, v, got)
			return
		}
	}
}

func TestCurrencyFloat64_EUR(t *testing.T) {
	for k, v := range testdata.TestCaseLangKZCurrencyEURFloat {
		got, err := CurrencyFloat64(k, WithCur(cur.EUR))
		if err != nil {
			t.Error(err)
			return
		}

		if got != v {
			t.Errorf("%.2f: \nexp: '%s' \ngot: '%s'", k, v, got)
			return
		}
	}
}

func TestCurrencyFloat64_KZT(t *testing.T) {
	for k, v := range testdata.TestCaseLangKZCurrencyKZTFloat {
		got, err := CurrencyFloat64(k, WithCur(cur.KZT))
		if err != nil {
			t.Error(err)
			return
		}

		if got != v {
			t.Errorf("%.2f: \nexp: '%s' \ngot: '%s'", k, v, got)
			return
		}
	}
}

func TestCurrencyFloat64_USD(t *testing.T) {
	for k, v := range testdata.TestCaseLangKZCurrencyUSDFloat {
		got, err := CurrencyFloat64(k, WithCur(cur.USD))
		if err != nil {
			t.Error(err)
			return
		}

		if got != v {
			t.Errorf("%.2f: \nexp: '%s' \ngot: '%s'", k, v, got)
			return
		}
	}
}

func TestCurrencyFloat64_Error(t *testing.T) {
	_, err := CurrencyFloat64(math.NaN())
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCurrencyString_Error(t *testing.T) {
	_, err := CurrencyString("bad")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestConvCurrency(t *testing.T) {
	// currency error
	_, err := convCurrency(ds.Empty, ds.Empty, cur.Currency(500), false, true)
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = convCurrency(ds.DigitString{DS: []digit.Digit{50, 50, 50}}, ds.Empty, cur.KZT, false, true)
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = convCurrency(ds.Empty, ds.DigitString{DS: []digit.Digit{50, 50, 50}}, cur.KZT, false, true)
	if err == nil {
		t.Fatal("expected error")
	}
}
