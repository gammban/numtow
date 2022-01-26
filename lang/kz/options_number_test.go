package kz

import (
	"testing"
)

func TestParseOpts(t *testing.T) {
	opts := ParseOpts(WithParseSep(','), WithParseFracLen(5), WithFmtFracUseDigits(true), WithFmtFracIgnore(true))

	o := prepareOptions(opts...)

	if o.ParseSeparator != ',' {
		t.Fatal("ParseSeparator mismatch")
	}

	if o.ParseFracLen != 5 {
		t.Fatal("ParseFracLen mismatch")
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
