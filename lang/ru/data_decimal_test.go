package ru

import (
	"testing"

	"github.com/gammban/numtow/lang/ru/gender"
)

func TestGetFracPart(t *testing.T) {
	if getFracPart(DeclinationSingular, gender.Male, 1) != "десятых" {
		t.Fatal("mismatch")
	}
	// test bad gender
	if getFracPart(DeclinationSingular, gender.Gender(45), 1) != "" {
		t.Fatal("mismatch")
	}
	// test bad gender
	if getFracPart(DeclinationSingular, gender.Unknown, 1) != "" {
		t.Fatal("mismatch")
	}
	// test bad idx
	if getFracPart(DeclinationSingular, gender.Male, 500) != "" {
		t.Fatal("mismatch")
	}
	// test bad declination for gender.Male
	if getFracPart(Declination(-5), gender.Male, 1) != "" {
		t.Fatal("mismatch")
	}
	// test bad declination for gender.Female
	if getFracPart(Declination(-4), gender.Female, 1) != "" {
		t.Fatal("mismatch")
	}
	// test bad declination for gender.Neuter
	if getFracPart(Declination(-3), gender.Neuter, 1) != "" {
		t.Fatal("mismatch")
	}

	if getFracPart(DeclinationSingular, gender.Neuter, 1) != "десятое" {
		t.Fatal("mismatch")
	}

	if getFracPart(DeclinationPlural, gender.Neuter, 2) != "сотых" {
		t.Fatal("mismatch")
	}
}

func TestGetIntPart(t *testing.T) {
	if getIntPart(DeclinationPlural, gender.Gender(10)) != "" {
		t.Fatal("mismatch")
	}

	if getIntPart(DeclinationPlural, gender.Unknown) != "" {
		t.Fatal("mismatch")
	}

	if getIntPart(Declination(-3), gender.Male) != "" {
		t.Fatal("mismatch")
	}

	if getIntPart(Declination(-2), gender.Female) != "" {
		t.Fatal("mismatch")
	}

	if getIntPart(Declination(-1), gender.Neuter) != "" {
		t.Fatal("mismatch")
	}
}
