package ru

import (
	"github.com/dantedenis/numtow/internal/digit"
	"github.com/dantedenis/numtow/internal/triplet"
)

type Declination int

const (
	DeclinationPlural Declination = iota
	DeclinationSingular
	Declination234
)

func getTripletDeclination(t triplet.Triplet) Declination {
	c := t.Units()

	if t.Tens() == digit.Digit1 {
		return DeclinationPlural
	}

	switch c {
	case digit.Digit1:
		return DeclinationSingular
	case digit.Digit2, digit.Digit3, digit.Digit4:
		return Declination234
	case digit.Digit0, digit.Digit5, digit.Digit6, digit.Digit7, digit.Digit8, digit.Digit9:
		return DeclinationPlural
	case digit.DigitUnknown:
		return DeclinationPlural
	default:
		return DeclinationPlural
	}
}
