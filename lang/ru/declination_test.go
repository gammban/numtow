package ru

import (
	"testing"

	"github.com/gammban/numtow/internal/triplet"
)

//nolint:gochecknoglobals
var testCaseDeclination = []struct {
	giveTriplet     triplet.Triplet
	wantDeclination Declination
}{
	{
		giveTriplet:     triplet.New(1, 2, 1),
		wantDeclination: DeclinationSingular,
	},
	{
		giveTriplet:     triplet.New(1, 2, 2),
		wantDeclination: Declination234,
	},
	{
		giveTriplet:     triplet.New(1, 2, 3),
		wantDeclination: Declination234,
	},
	{
		giveTriplet:     triplet.New(1, 2, 4),
		wantDeclination: Declination234,
	},
	{
		giveTriplet:     triplet.New(1, 2, 5),
		wantDeclination: DeclinationPlural,
	},
	{
		giveTriplet:     triplet.New(1, 2, 6),
		wantDeclination: DeclinationPlural,
	},
	{
		giveTriplet:     triplet.New(1, 2, 7),
		wantDeclination: DeclinationPlural,
	},
	{
		giveTriplet:     triplet.New(1, 2, 8),
		wantDeclination: DeclinationPlural,
	},
	{
		giveTriplet:     triplet.New(1, 2, 9),
		wantDeclination: DeclinationPlural,
	},
	{
		giveTriplet:     triplet.New(1, 2, 10),
		wantDeclination: DeclinationPlural,
	},
	{
		giveTriplet:     triplet.New(1, 2, 11),
		wantDeclination: DeclinationPlural,
	},
}

func TestGetTripletDeclination(t *testing.T) {
	for _, v := range testCaseDeclination {
		if gotDecl := getTripletDeclination(v.giveTriplet); gotDecl != v.wantDeclination {
			t.Errorf("expected %d, got %d", v.wantDeclination, gotDecl)
		}
	}
}
