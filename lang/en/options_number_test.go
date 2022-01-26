package en

import "testing"

func TestParseOpts(t *testing.T) {
	opts := ParseOpts(WithParseSep(','), WithParseFracLen(5), WithFmtFracUseDigits(true), WithFmtFracIgnore(true), WithFmtNegative())

	o := prepareOptions(opts...)

	if o.Parse.DecimalSeparator != ',' {
		t.Fatal("ParseSeparator mismatch")
	}

	if o.Format.MinusSignWord != negative {
		t.Fatal("MinusSignWord mismatch")
	}

	if o.Parse.FracLen != 5 {
		t.Fatal("ParseFracLen mismatch")
	}

	if !o.Format.FracUseDigits {
		t.Fatal("FmtFracUseDigits mismatch")
	}

	if !o.Format.FracIgnore {
		t.Fatal("FmtFracIgnore mismatch")
	}

	opts = ParseOpts([]OptFunc{WithParseSep(','), WithParseFracLen(5)})

	o = prepareOptions(opts...)

	if o.Parse.DecimalSeparator != ',' {
		t.Fatal("ParseSeparator mismatch")
	}

	if o.Parse.FracLen != 5 {
		t.Fatal("ParseFracLen mismatch")
	}

	opts = ParseOpts()
	if len(opts) != 0 {
		t.Fatal("result mismatch")
	}
}
