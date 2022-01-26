package gender

import (
	"errors"
	"testing"
)

func TestGender_Validate(t *testing.T) {
	err := Male.Validate()
	if err != nil {
		t.Fatal(err)
	}

	err = Female.Validate()
	if err != nil {
		t.Fatal(err)
	}

	err = Neuter.Validate()
	if err != nil {
		t.Fatal(err)
	}

	err = Unknown.Validate()
	if !errors.Is(err, ErrBadGender) {
		t.Fatal("expected ErrBadGender")
	}

	err = Gender(10).Validate()
	if !errors.Is(err, ErrBadGender) {
		t.Fatal("expected ErrBadGender")
	}
}
