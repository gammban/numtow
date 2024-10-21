package kz

import (
	"testing"

	"github.com/gammban/numtow/internal/digit"
	"github.com/gammban/numtow/internal/ds"
)

func TestConvert(t *testing.T) {
	res, err := convert(ds.Empty)
	if err != nil {
		t.Fatal(err)
	}

	if res != "" {
		t.Fatal("expected empty string")
	}

	_, err = convert(ds.DigitString{DS: []digit.Digit{digit.Digit(15)}})
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = convert(ds.DigitString{DS: []digit.Digit{1, digit.Digit(15), 2}})
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = convert(ds.DigitString{DS: []digit.Digit{digit.Digit(15), 5, 2}})
	if err == nil {
		t.Fatal("expected error")
	}

	bigDS := ds.DigitString{DS: []digit.Digit{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}

	_, err = convert(bigDS)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestConvertDecimal(t *testing.T) {
	_, err := convertDecimal(ds.DigitString{}, ds.DigitString{IsSignMinus: true})
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = convertDecimal(ds.New(50), ds.DigitString{})
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = convertDecimal(ds.Empty, ds.New(60))
	if err == nil {
		t.Fatal("expected error")
	}

	longDS, err := ds.ParseString("123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
	if err != nil {
		t.Fatal("unexpected")
	}

	_, err = convertDecimal(ds.Empty, longDS)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestGetFracPartName(t *testing.T) {
	_, err := getFracPartName(ds.Empty)
	if err == nil {
		t.Fatal("mismatch")
	}

	longDS, err := ds.ParseString("123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
	if err != nil {
		t.Fatal("unexpected")
	}

	_, err = getFracPartName(longDS)
	if err == nil {
		t.Fatal("mismatch")
	}
}
