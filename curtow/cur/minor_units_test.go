package cur

import (
	"errors"
	"testing"
)

//nolint:gochecknoglobals // test cases
var testCaseMU = []struct {
	giveMU          MinorUnits
	wantString      string
	wantZero        string
	wantValidateErr error
}{
	{
		giveMU:     MinorUnits0,
		wantString: "0",
		wantZero:   "",
	},
	{
		giveMU:     MinorUnits2,
		wantString: "2",
		wantZero:   "00",
	},
	{
		giveMU:     MinorUnits3,
		wantString: "3",
		wantZero:   "000",
	},
	{
		giveMU:          MinorUnitsUnknown,
		wantString:      "",
		wantZero:        "",
		wantValidateErr: ErrBadMinorUnit,
	},
	{
		giveMU:          MinorUnits(10),
		wantString:      "",
		wantZero:        "",
		wantValidateErr: ErrBadMinorUnit,
	},
}

func TestMinorUnits_String(t *testing.T) {
	for _, v := range testCaseMU {
		if gotString := v.giveMU.String(); gotString != v.wantString {
			t.Errorf("expected %s, got %s", v.wantString, gotString)
		}
	}
}

func TestMinorUnits_Zero(t *testing.T) {
	for _, v := range testCaseMU {
		if gotZero := v.giveMU.Zero(); gotZero != v.wantZero {
			t.Errorf("expected '%s', got '%s'", v.wantZero, gotZero)
		}
	}
}

func TestMinorUnits_Validate(t *testing.T) {
	for _, v := range testCaseMU {
		if gotErr := v.giveMU.Validate(); !errors.Is(gotErr, v.wantValidateErr) {
			t.Errorf("expected '%s', got '%s'", v.wantValidateErr, gotErr)
		}
	}
}
