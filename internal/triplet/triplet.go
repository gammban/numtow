package triplet

import (
	"errors"

	"github.com/dantedenis/numtow/internal/digit"
)

type Triplet struct {
	hundreds digit.Digit
	tens     digit.Digit
	units    digit.Digit
}

var (
	ErrBadTriplet = errors.New("bad triplet")
)

func New(hundreds, tens, units digit.Digit) Triplet {
	return Triplet{
		hundreds: hundreds,
		tens:     tens,
		units:    units,
	}
}

func NewZero() Triplet {
	return Triplet{}
}

func (t Triplet) Validate() error {
	if t.hundreds == digit.DigitUnknown {
		return ErrBadTriplet
	}

	if t.tens == digit.DigitUnknown {
		return ErrBadTriplet
	}

	if t.units == digit.DigitUnknown {
		return ErrBadTriplet
	}

	return nil
}

func (t Triplet) Hundreds() digit.Digit {
	return t.hundreds
}

func (t Triplet) Tens() digit.Digit {
	return t.tens
}

func (t Triplet) Units() digit.Digit {
	return t.units
}

func (t *Triplet) SetHundreds(d digit.Digit) {
	t.hundreds = d
}

func (t *Triplet) SetTens(d digit.Digit) {
	t.tens = d
}

func (t *Triplet) SetUnits(d digit.Digit) {
	t.units = d
}

func (t Triplet) IsZero() bool {
	return t.units == 0 && t.tens == 0 && t.hundreds == 0
}

func (t Triplet) String() string {
	return t.hundreds.String() + t.tens.String() + t.units.String()
}
