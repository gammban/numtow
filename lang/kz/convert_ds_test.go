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
