package kz

import (
	"testing"

	"github.com/gammban/numtow/curtow/cur"
)

func TestParseCurrencyOpts_Variadic(t *testing.T) {
	opts := ParseCurrencyOpts(WithCur(cur.EUR), WithCurIgnoreMU(true), WithCurConvMU(true))

	o := prepareCurrencyOptions(opts...)

	if o.currency != cur.EUR {
		t.Fatal("mismatch")
	}

	if !o.ignoreMinorUnits {
		t.Fatal("mismatch")
	}

	if !o.convertMinorUnits {
		t.Fatal("mismatch")
	}

	opts = ParseCurrencyOpts([]CurrencyOpt{WithCur(cur.EUR), WithCurIgnoreMU(true), WithCurConvMU(true)})

	o = prepareCurrencyOptions(opts...)

	if o.currency != cur.EUR {
		t.Fatal("mismatch")
	}

	if !o.ignoreMinorUnits {
		t.Fatal("mismatch")
	}

	if !o.convertMinorUnits {
		t.Fatal("mismatch")
	}

	opts = ParseCurrencyOpts()

	if len(opts) != 0 {
		t.Fatal("mismatch")
	}
}
