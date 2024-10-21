package ru

import (
	"testing"

	"github.com/dantedenis/numtow/lang/ru/gender"
)

func TestParseOpts(t *testing.T) {
	opts := ParseOpts(WithParseSep(','), WithParseFracLen(5), WithFmtFracUseDigits(true), WithFmtFracIgnore(true), WithFmtGender(gender.Male))

	o := prepareOptions(opts...)

	if o.ParseSeparator != ',' {
		t.Fatal("ParseSeparator mismatch")
	}

	if o.ParseFracLen != 5 {
		t.Fatal("ParseFracLen mismatch")
	}

	if o.FmtGender != gender.Male {
		t.Fatal("FmtGender mismatch")
	}

	if !o.FmtFracUseDigits {
		t.Fatal("FmtFracUseDigits mismatch")
	}

	if !o.FmtFracIgnore {
		t.Fatal("FmtFracIgnore mismatch")
	}

	opts = ParseOpts([]OptFunc{WithParseSep(','), WithParseFracLen(5)})

	o = prepareOptions(opts...)

	if o.ParseSeparator != ',' {
		t.Fatal("ParseSeparator mismatch")
	}

	if o.ParseFracLen != 5 {
		t.Fatal("ParseFracLen mismatch")
	}

	opts = ParseOpts()
	if len(opts) != 0 {
		t.Fatal("result mismatch")
	}
}
