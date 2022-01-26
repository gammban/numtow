package triplet

import "testing"

func TestTriplets_Len(t *testing.T) {
	var tr Triplets

	if tr.Len() != 0 {
		t.Fatal("expected 0")
	}

	tr = []Triplet{NewZero()}
	if tr.Len() != 1 {
		t.Fatal("expected 1")
	}
}

func TestTriplets_IsZero(t *testing.T) {
	var tr Triplets

	if tr.IsZero() {
		t.Fatal("expected false")
	}

	tr = []Triplet{NewZero()}
	if !tr.IsZero() {
		t.Fatal("expected true")
	}

	tr = []Triplet{NewZero(), NewZero(), NewZero()}
	if !tr.IsZero() {
		t.Fatal("expected true")
	}

	tr = []Triplet{NewZero(), New(0, 0, 1), NewZero()}
	if tr.IsZero() {
		t.Fatal("expected false")
	}
}

func TestTriplets_String(t *testing.T) {
	var tr Triplets

	if tr.String() != "" {
		t.Fatal("expected empty string")
	}

	tr = []Triplet{NewZero()}
	if tr.String() != "000" {
		t.Fatal("expected 000")
	}
}
