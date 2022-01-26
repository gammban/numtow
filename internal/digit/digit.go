package digit

import "errors"

type Digit uint8

const (
	Digit0 Digit = iota
	Digit1
	Digit2
	Digit3
	Digit4
	Digit5
	Digit6
	Digit7
	Digit8
	Digit9
	DigitUnknown
)

var (
	ErrBadDigit = errors.New("bad digit")
)

func (d Digit) Validate() error {
	switch d {
	case Digit0, Digit1, Digit2, Digit3, Digit4, Digit5, Digit6, Digit7, Digit8, Digit9:
		return nil
	case DigitUnknown:
		return ErrBadDigit
	default:
		return ErrBadDigit
	}
}

func (d Digit) IsZero() bool {
	return d == Digit0
}

func (d Digit) IsNotZero() bool {
	return d != Digit0
}

func (d Digit) String() string {
	switch d {
	case Digit0:
		return "0"
	case Digit1:
		return "1"
	case Digit2:
		return "2"
	case Digit3:
		return "3"
	case Digit4:
		return "4"
	case Digit5:
		return "5"
	case Digit6:
		return "6"
	case Digit7:
		return "7"
	case Digit8:
		return "8"
	case Digit9:
		return "9"
	case DigitUnknown:
		return ""
	default:
		return ""
	}
}

func ParseRune(s rune) Digit {
	switch s {
	case '0':
		return Digit0
	case '1':
		return Digit1
	case '2':
		return Digit2
	case '3':
		return Digit3
	case '4':
		return Digit4
	case '5':
		return Digit5
	case '6':
		return Digit6
	case '7':
		return Digit7
	case '8':
		return Digit8
	case '9':
		return Digit9
	default:
		return DigitUnknown
	}
}
