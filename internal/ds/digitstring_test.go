package ds

import (
	"errors"
	"testing"

	"github.com/gammban/numtow/internal/digit"
)

//nolint:gochecknoglobals // test cases for DigitString.Split method
var testCaseDSSplit = []struct {
	giveDS     DigitString
	needString string
	needErr    error
}{
	{
		giveDS: DigitString{
			DS: []digit.Digit{1, 2, 3},
		},
		needString: "123",
	},
	{
		giveDS: DigitString{
			DS: []digit.Digit{0},
		},
		needString: "000",
	},
	{
		giveDS: DigitString{
			DS: []digit.Digit{1, 2, 3, 4},
		},

		needString: "001234",
	},
	{
		giveDS: DigitString{
			DS: []digit.Digit{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},

		needString: "123456789",
	},
	{
		giveDS:     DigitString{},
		needString: "",
		needErr:    ErrParse,
	},
	{
		giveDS: DigitString{
			DS: []digit.Digit{1, 11, 2, 3},
		},
		needString: "",
		needErr:    digit.ErrBadDigit,
	},
}

func TestDigitString_Split(t *testing.T) {
	for _, v := range testCaseDSSplit {
		v := v

		t.Run("", func(tt *testing.T) {
			gotTriplets, gotErr := v.giveDS.Split()
			if !errors.Is(gotErr, v.needErr) {
				tt.Fatal(gotErr)
				return
			}

			if gotString := gotTriplets.String(); gotString != v.needString {
				tt.Errorf("got %s, expected %s", gotString, v.needString)
				return
			}
		})
	}
}

func TestDigitString_FirstTriplet(t *testing.T) {
	d := DigitString{
		DS:          []digit.Digit{5, 0, 1, 2, 3},
		IsSignMinus: false,
	}

	tr, err := d.FirstTriplet()
	if err != nil {
		t.Fatal(err)
	}

	if tr.String() != "123" {
		t.Fatal("mismatch")
	}

	emptyDS := DigitString{}

	_, err = emptyDS.FirstTriplet()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestDigitString_Len(t *testing.T) {
	emptyDS := DigitString{DS: []digit.Digit{}}
	if emptyDS.Len() != 0 {
		t.Fatal("mismatch")
	}

	emptyDS = DigitString{DS: nil}
	if emptyDS.Len() != 0 {
		t.Fatal("mismatch")
	}

	d := DigitString{DS: []digit.Digit{digit.Digit0}}
	if d.Len() != 1 {
		t.Fatal("mismatch")
	}
}

func TestDigitString_IsZero(t *testing.T) {
	if Empty.IsZero() {
		t.Fatal("mismatch")
	}

	if !Zero.IsZero() {
		t.Fatal("mismatch")
	}

	d := DigitString{DS: []digit.Digit{0}}
	if !d.IsZero() {
		t.Fatal("mismatch")
	}

	d = DigitString{DS: []digit.Digit{0, 0}}
	if !d.IsZero() {
		t.Fatal("mismatch")
	}

	d = DigitString{DS: []digit.Digit{0, 0, 0, 0, 0}}
	if !d.IsZero() {
		t.Fatal("mismatch")
	}

	d = DigitString{DS: []digit.Digit{0, 0, 0, 0, 1, 0}}
	if d.IsZero() {
		t.Fatal("mismatch")
	}

	d = DigitString{DS: []digit.Digit{1, 0, 0, 0, 0, 0, 0}}
	if d.IsZero() {
		t.Fatal("mismatch")
	}
}

func TestNew(t *testing.T) {
	d := New()
	if !d.IsEmpty() {
		t.Fatal("expected empty ds")
	}

	d = New(1)
	if got := d.String(); got != "1" {
		t.Fatal("expected 1 got", got)
	}

	d = New(1, 2, 3)
	if got := d.String(); got != "123" {
		t.Fatal("expected 123 got", got)
	}
}
