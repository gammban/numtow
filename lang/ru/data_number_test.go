package ru

import (
	"testing"

	"github.com/gammban/numtow/lang/ru/gender"
)

func TestGetUnitsByGender(t *testing.T) {
	_, err := getUnitsByGender(gender.Unknown, 1)
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = getUnitsByGender(gender.Gender(50), 1)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestGetMegaByDeclination(t *testing.T) {
	if getMegaByDeclination(Declination(55), 1) != "тысяч" {
		t.Fatal("expected тысяч")
	}

	if getMegaByDeclination(DeclinationSingular, 1) != "тысяча" {
		t.Fatal("expected тысяча")
	}

	if getMegaByDeclination(Declination234, 1) != "тысячи" {
		t.Fatal("expected тысячи")
	}

	if getMegaByDeclination(DeclinationPlural, 2) != "миллионов" {
		t.Fatal("expected миллионов")
	}
}
