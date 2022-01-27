package ds

import (
	"errors"
	"math"
	"testing"
)

//nolint:gochecknoglobals
var testCaseParseDecimal = []struct {
	giveDec  string
	giveExp  uint
	giveSep  rune
	wantInt  string
	wantFrac string
	wantErr  error
}{
	{giveDec: "5.8", wantInt: "5", wantFrac: "8"},
	{giveDec: "548725300538597.89", wantInt: "548725300538597", wantFrac: "89"},
	{giveDec: "-5.8", wantInt: "-5", wantFrac: "8"},
	{giveDec: "-5.+8", wantInt: "", wantFrac: "", wantErr: ErrParse},
	{giveDec: "0.8", wantInt: "0", wantFrac: "8"},
	{giveDec: "123", wantInt: "123", wantFrac: ""},
	{giveDec: "-123", wantInt: "-123", wantFrac: ""},
	{giveDec: "-5.", wantInt: "", wantFrac: "", wantErr: ErrParse},
	{giveDec: "0.5", wantInt: "0", wantFrac: "5"},
	{giveDec: "-0.5", wantInt: "-0", wantFrac: "5"},
	{giveDec: "-123444656.89", wantInt: "-123444656", wantFrac: "89"},
	{giveDec: "-123444656.-89", wantInt: "", wantFrac: "", wantErr: ErrParse},
	{giveDec: "5.8", giveExp: 2, wantInt: "5", wantFrac: "80"},
	{giveDec: "-5.82", giveExp: 3, wantInt: "-5", wantFrac: "820"},
	{giveDec: "0.82", giveExp: 3, wantInt: "0", wantFrac: "820"},
	{giveDec: ".82", giveExp: 3, wantInt: "0", wantFrac: "820"},
	{giveDec: ".82", giveExp: 0, wantInt: "0", wantFrac: "82"},
	{giveDec: "-.5", wantInt: "-0", wantFrac: "5"},
	{giveDec: "-.5", giveExp: 4, wantInt: "-0", wantFrac: "5000"},
	{giveDec: "-.-5", wantInt: "", wantFrac: "", wantErr: ErrParse},
	{giveDec: "-.-5", giveExp: 2, wantInt: "", wantFrac: "", wantErr: ErrParse},
	{giveDec: "5", giveExp: 4, wantInt: "5", wantFrac: "0000"},
	{giveDec: "4", giveExp: 0, wantInt: "4", wantFrac: ""},
	{giveDec: "-123", giveExp: 0, wantInt: "-123", wantFrac: ""},
	{giveDec: "-548725300538597.89", giveExp: 3, wantInt: "-548725300538597", wantFrac: "890"},
	{giveDec: "1.a", giveExp: 2, wantErr: ErrParse},
	{giveDec: "1.2a", giveExp: 5, wantErr: ErrParse},
	{giveDec: "1.2a", giveExp: 0, wantErr: ErrParse},
	{giveDec: "1.2.3", giveExp: 0, wantErr: ErrParse},
	{giveDec: "1.2.3", giveExp: 2, wantErr: ErrParse},
	{giveDec: "1.2.3", giveExp: 2, wantErr: ErrParse},
	{giveDec: "1.25a3", giveExp: 3, wantErr: ErrParse},
	{giveDec: "0.123456789", giveExp: 2, wantInt: "0", wantFrac: "12"},
	{giveDec: "0.123456789", giveExp: 0, wantInt: "0", wantFrac: "123456789"},
	{giveDec: ",123456789", giveExp: 1, giveSep: ',', wantInt: "0", wantFrac: "1"},
	{giveDec: "-.123456789", giveExp: 5, wantInt: "-0", wantFrac: "12345"},
	{giveDec: "100", giveExp: 2, wantInt: "100", wantFrac: "00"},
	{giveDec: "00001.02", giveExp: 2, wantInt: "00001", wantFrac: "02"},
	{giveDec: "1.bad", wantErr: ErrParse},
}

func TestParseDecimal(t *testing.T) {
	for _, v := range testCaseParseDecimal {
		sep := defaultDecimalSeparator
		if v.giveSep != 0 {
			sep = v.giveSep
		}

		gotInt, gotFrac, gotErr := ParseDecimal(v.giveDec, WithFracLen(v.giveExp), WithSep(sep))
		if !errors.Is(gotErr, v.wantErr) {
			t.Errorf("%s: err result mismatch: got %s, expected %s", v.giveDec, gotErr, v.wantErr)
			return
		}

		if gotInt.String() != v.wantInt {
			t.Errorf("%s: int result mismatch: got '%s', expected '%s'", v.giveDec, gotInt, v.wantInt)
			return
		}

		if gotFrac.String() != v.wantFrac {
			t.Errorf("%s: frac result mismatch: got '%s', expected '%s'", v.giveDec, gotFrac, v.wantFrac)
			return
		}
	}
}

//nolint:gochecknoglobals
var testCaseParseInt64 = []struct {
	give       int64
	wantString string
	wantErr    error
}{
	{give: 123, wantString: "123", wantErr: nil},
	{give: -123, wantString: "-123", wantErr: nil},
	{give: 0, wantString: "0", wantErr: nil},
	{give: 10, wantString: "10", wantErr: nil},
	{give: -0, wantString: "0", wantErr: nil},
	{give: 45, wantString: "45", wantErr: nil},
	{give: 9223372036854775807, wantString: "9223372036854775807", wantErr: nil},
	{give: -9223372036854775808, wantString: "-9223372036854775808", wantErr: nil},
}

func TestParseInt64(t *testing.T) {
	for _, v := range testCaseParseInt64 {
		gotDS := ParseInt64(v.give)
		if gotDS.IsEmpty() {
			t.Error("mismatch")
			return
		}

		if gotString := gotDS.String(); gotString != v.wantString {
			t.Errorf("mismatch: got %s, expected %s", gotString, v.wantString)
			return
		}
	}
}

//nolint:gochecknoglobals
var testCaseParseString = []struct {
	give       string
	wantString string
	wantErr    error
}{
	{give: "123", wantString: "123", wantErr: nil},
	{give: "-123", wantString: "-123", wantErr: nil},
	{give: "0", wantString: "0", wantErr: nil},
	{give: "-a0", wantString: "", wantErr: ErrParse},
	{give: "--0", wantString: "", wantErr: ErrParse},
	{give: "", wantString: "", wantErr: ErrParse},
	{give: "45\n45", wantString: "", wantErr: ErrParse},
	{give: "-45\n45", wantString: "", wantErr: ErrParse},
	{give: "179769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368", wantString: "179769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368", wantErr: nil},
	{give: "-179769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368", wantString: "-179769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368", wantErr: nil},
}

func TestParseString(t *testing.T) {
	for _, v := range testCaseParseString {
		gotDS, gotErr := ParseString(v.give)
		if !errors.Is(gotErr, v.wantErr) {
			t.Error("mismatch")
			return
		}

		if gotString := gotDS.String(); gotString != v.wantString {
			t.Errorf("mismatch: got %s, expected %s", gotString, v.wantString)
			return
		}
	}
}

//nolint:gochecknoglobals
var testCaseParseFloat64 = []struct {
	giveFloat64 float64
	giveExp     uint
	wantInt     string
	wantFrac    string
	wantErr     error
}{
	{giveFloat64: 5.8, wantInt: "5", wantFrac: "8"},
	{giveFloat64: -5.8, wantInt: "-5", wantFrac: "8"},
	{giveFloat64: 0.8, wantInt: "0", wantFrac: "8"},
	{giveFloat64: 123, wantInt: "123", wantFrac: ""},
	{giveFloat64: -123, wantInt: "-123", wantFrac: ""},
	{giveFloat64: -5., wantInt: "-5", wantFrac: ""},
	{giveFloat64: 0.5, wantInt: "0", wantFrac: "5"},
	{giveFloat64: -0.5, wantInt: "-0", wantFrac: "5"},
	{giveFloat64: -123444656.89, wantInt: "-123444656", wantFrac: "89"},
	{giveFloat64: 5.8, giveExp: 2, wantInt: "5", wantFrac: "80"},
	{giveFloat64: -5.82, giveExp: 3, wantInt: "-5", wantFrac: "820"},
	{giveFloat64: 0.82, giveExp: 3, wantInt: "0", wantFrac: "820"},
	{giveFloat64: 0.001, giveExp: 0, wantInt: "0", wantFrac: "001"},
	{giveFloat64: .82, giveExp: 3, wantInt: "0", wantFrac: "820"},
	{giveFloat64: .82, giveExp: 0, wantInt: "0", wantFrac: "82"},
	{giveFloat64: -.5, wantInt: "-0", wantFrac: "5"},
	{giveFloat64: -.5, giveExp: 4, wantInt: "-0", wantFrac: "5000"},
	{giveFloat64: 5, giveExp: 4, wantInt: "5", wantFrac: "0000"},
	{giveFloat64: 4, giveExp: 0, wantInt: "4", wantFrac: ""},
	{giveFloat64: -123, giveExp: 0, wantInt: "-123", wantFrac: ""},
	{giveFloat64: math.NaN(), giveExp: 3, wantErr: ErrParse},
	{giveFloat64: math.Inf(1), giveExp: 3, wantErr: ErrParse},
	{giveFloat64: math.Inf(-1), giveExp: 3, wantErr: ErrParse},
	{giveFloat64: 548725300538597.89, giveExp: 3, wantErr: ErrParse},
	{giveFloat64: -548725300538597.89, giveExp: 3, wantErr: ErrParse},
	{giveFloat64: 0.123456789, giveExp: 2, wantInt: "0", wantFrac: "12"},
	{giveFloat64: 0.123456789, giveExp: 0, wantInt: "0", wantFrac: "123456789"},
	{giveFloat64: .123456789, giveExp: 1, wantInt: "0", wantFrac: "1"},
	{giveFloat64: -.123456789, giveExp: 5, wantInt: "-0", wantFrac: "12345"},
	{giveFloat64: 999999999999, giveExp: 5, wantInt: "999999999999", wantFrac: "00000"},
	{giveFloat64: 999999999999.9999, giveExp: 5, wantInt: "999999999999", wantFrac: "99990"},
	{giveFloat64: -999999999999.9999, giveExp: 5, wantInt: "-999999999999", wantFrac: "99990"},
	{giveFloat64: 999999999999.9999, giveExp: 0, wantInt: "999999999999", wantFrac: "9999"},
	{giveFloat64: -999999999999.9999, giveExp: 0, wantInt: "-999999999999", wantFrac: "9999"},
}

func TestParseFloat64(t *testing.T) {
	for _, v := range testCaseParseFloat64 {
		gotInt, gotFrac, gotErr := ParseFloat64(v.giveFloat64, WithFracLen(v.giveExp))
		if !errors.Is(gotErr, v.wantErr) {
			t.Errorf("%f: err result mismatch: got %s, expected %s", v.giveFloat64, gotErr, v.wantErr)
			return
		}

		if gotInt.String() != v.wantInt {
			t.Errorf("%f: int result mismatch: got '%s', expected '%s'", v.giveFloat64, gotInt, v.wantInt)
			return
		}

		if gotFrac.String() != v.wantFrac {
			t.Errorf("%f: frac result mismatch: got '%s', expected '%s'", v.giveFloat64, gotFrac, v.wantFrac)
			return
		}
	}
}
