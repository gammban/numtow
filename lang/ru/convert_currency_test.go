package ru

import (
	"math"
	"strings"
	"testing"

	"github.com/gammban/numtow/internal/ds"

	"github.com/gammban/numtow/curtow/cur"
	"github.com/gammban/numtow/internal/testdata"
)

func TestConvCurrencyInt64_CurrencyKZT(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUCurrencyKZTInt {
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

func TestConvCurrencyString_CurrencyKZT(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUCurrencyKZTString {
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

func TestConvCurrencyInt64_CurrencyRUB(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUCurrencyRUBInt {
		got, err := CurrencyInt64(k, WithCur(cur.RUB))
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

func TestConvCurrencyString_CurrencyRUB(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUCurrencyRUBString {
		got, err := CurrencyString(k, WithCur(cur.RUB))
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(got, v) {
			t.Errorf("%s: \nexp: '%s' \ngot: '%s'", k, v, got)
			return
		}
	}
}

func TestConvCurrencyInt64_CurrencyUSD(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUCurrencyUSDInt {
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

func TestConvCurrencyString_CurrencyUSD(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUCurrencyUSDString {
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

func TestConvCurrencyInt64_CurrencyEUR(t *testing.T) {
	for k, v := range testdata.TestCaseLangRUCurrencyEURInt {
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

func TestConvCurrencyString_CurrencyEUR(t *testing.T) {
	for _, v := range testdata.TestCaseLangRUCurrencyEURString {
		gotWords, err := CurrencyString(v.GiveAmount, WithCur(cur.EUR), WithCurConvMU(v.ConvertMinorUnits))
		if err != nil {
			t.Error(err)
			return
		}

		if !strings.EqualFold(gotWords, v.WantWords) {
			t.Errorf("%s: \nexp: '%s' \ngot: '%s'", v.GiveAmount, v.WantWords, gotWords)
			return
		}
	}
}

func TestConvCurrency_Errors(t *testing.T) {
	_, err := CurrencyString("a", WithCur(cur.EUR))
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = CurrencyString("a", WithCur(cur.Currency(20)))
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = CurrencyFloat64(math.NaN(), WithCur(cur.EUR))
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = CurrencyFloat64(123, WithCur(cur.Currency(20)))
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = CurrencyString("1.-2")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestConvCurrencyString(t *testing.T) {
	res, err := CurrencyString("100", WithCur(cur.RUB), WithCurConvMU(true), WithCurIgnoreMU(false))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}

func TestConvCurrency(t *testing.T) {
	_, err := convCurrency(ds.New(50), ds.Empty, cur.USD, false, false)
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = convCurrency(ds.Empty, ds.New(50), cur.USD, false, false)
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = convCurrency(ds.Empty, ds.Empty, cur.USD, false, false)
	if err == nil {
		t.Fatal("expected error")
	}
}
