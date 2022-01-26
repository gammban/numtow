package digit

import (
	"errors"
	"strconv"
	"testing"
)

func TestDigit_Validate(t *testing.T) {
	for i := uint8(0); i < 10; i++ {
		err := Digit(i).Validate()
		if err != nil {
			t.Fatal("expected nil got", err)
		}
	}

	if err := DigitUnknown.Validate(); !errors.Is(err, ErrBadDigit) {
		t.Fatal("expected ErrBadDigit")
	}

	if err := Digit(11).Validate(); !errors.Is(err, ErrBadDigit) {
		t.Fatal("expected ErrBadDigit")
	}
}

func TestDigit_IsZero(t *testing.T) {
	if !Digit(0).IsZero() {
		t.Fatal("expected true")
	}

	for i := uint8(1); i < 10; i++ {
		if Digit(i).IsZero() {
			t.Fatal("expected false")
		}
	}
}

func TestDigit_IsNotZero(t *testing.T) {
	if Digit(0).IsNotZero() {
		t.Fatal("expected false")
	}

	for i := uint8(1); i < 10; i++ {
		if !Digit(i).IsNotZero() {
			t.Fatal("expected true")
		}
	}
}

func TestDigit_String(t *testing.T) {
	for i := uint8(0); i < 10; i++ {
		if got := Digit(i).String(); got != strconv.Itoa(int(i)) {
			t.Fatal("mismatch")
		}
	}

	if DigitUnknown.String() != "" {
		t.Fatal("mismatch")
	}

	if Digit(11).String() != "" {
		t.Fatal("mismatch")
	}
}

func TestParseRune(t *testing.T) {
	runes := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, v := range runes {
		if ParseRune(v) == DigitUnknown {
			t.Fatal("mismatch")
		}
	}

	if ParseRune('a') != DigitUnknown {
		t.Fatal("expected DigitUnknown")
	}
}
