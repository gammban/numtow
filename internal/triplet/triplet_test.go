package triplet

import (
	"errors"
	"testing"

	"github.com/gammban/numtow/internal/digit"
)

func TestTriplet_Validate(t *testing.T) {
	tr := New(10, 0, 0)
	if err := tr.Validate(); !errors.Is(err, ErrBadTriplet) {
		t.Fatal("expected ErrBadTriplet")
	}

	tr = New(0, 10, 0)
	if err := tr.Validate(); !errors.Is(err, ErrBadTriplet) {
		t.Fatal("expected ErrBadTriplet")
	}

	tr = New(0, 0, 10)
	if err := tr.Validate(); !errors.Is(err, ErrBadTriplet) {
		t.Fatal("expected ErrBadTriplet")
	}

	tr = New(0, 0, 0)
	if err := tr.Validate(); err != nil {
		t.Fatal("expected no error")
	}
}

func TestTriplet_SetHundreds(t *testing.T) {
	tr := NewZero()

	tr.SetHundreds(digit.Digit5)

	if tr.Hundreds() != digit.Digit5 {
		t.Fatal("expected 5")
	}

	tr.SetTens(digit.Digit2)

	if tr.Tens() != digit.Digit2 {
		t.Fatal("expected 2")
	}

	tr.SetUnits(digit.Digit9)

	if tr.Units() != digit.Digit9 {
		t.Fatal("expected 9")
	}
}

func TestTriplet_String(t *testing.T) {
	if NewZero().String() != "000" {
		t.Fatal("expected 000")
	}
}

func TestTriplet_IsZero(t *testing.T) {
	if !NewZero().IsZero() {
		t.Log("expected true")
	}
}
