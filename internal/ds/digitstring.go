package ds

import (
	"fmt"
	"strings"

	"github.com/gammban/numtow/internal/digit"
	"github.com/gammban/numtow/internal/triplet"
)

// DigitString holds slice of digits (digit.Digit) and minus sign.
type DigitString struct {
	DS          []digit.Digit
	IsSignMinus bool
}

//nolint:gochecknoglobals
var (
	// Empty DigitString
	Empty = DigitString{}
	// Zero DigitString
	Zero = DigitString{DS: []digit.Digit{digit.Digit0}}
)

// New returns DigitString. Returned DigitString can be invalid. Use Validate to check it.
func New(d ...digit.Digit) DigitString {
	if len(d) == 0 {
		return Empty
	}

	data := make([]digit.Digit, 0, len(d))

	for i := range d {
		data = append(data, d[i])
	}

	return DigitString{DS: data}
}

// IsEmpty returns true if DigitString is empty
// 	Examples:
// 	Empty.IsEmpty() // true
// 	Zero.IsEmpty() // false
// 	New().IsEmpty() // true
// 	New(0).IsEmpty() // false
// 	New(1, 2, 3).IsEmpty() // false
func (ds DigitString) IsEmpty() bool {
	return len(ds.DS) == 0
}

// IsZero returns true if DigitString is zero
// 	Examples:
// 	Empty.IsZero() // false
// 	Zero.IsZero() // true
// 	New(0).IsZero() // true
// 	New(0, 0, 0, 0, 0, 0).IsZero() // true
// 	New(1).IsZero() // false
func (ds DigitString) IsZero() bool {
	if len(ds.DS) == 0 {
		return false
	}

	for i := range ds.DS {
		if ds.DS[i].IsNotZero() {
			return false
		}
	}

	return true
}

// String returns string value of DigitString.
// 	Examples:
// 	Empty.String() // ""
// 	Zero.String() // "0"
// 	New(1).String() // "1"
// 	New(1, 2, 3).String() // "123"
func (ds DigitString) String() string {
	sb := strings.Builder{}

	if len(ds.DS) != 0 && ds.IsSignMinus {
		sb.WriteString("-")
	}

	for k := range ds.DS {
		sb.WriteString(ds.DS[k].String())
	}

	return sb.String()
}

// Validate returns error if DigitString is not valid.
// 	Examples:
// 	Empty.Validate() // nil
// 	Zero.Validate() // nil
// 	New(10).Validate() // digit.ErrBadDigit
// 	New(1).Validate() // nil
// 	New(15).Validate() // digit.ErrBadDigit
func (ds DigitString) Validate() error {
	for i := range ds.DS {
		err := ds.DS[i].Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// Len returns length of DigitString
// 	Examples:
// 	Empty.Len() // 0
// 	Zero.Len() // 1
// 	New(1, 2).Len() // 2
func (ds DigitString) Len() int {
	return len(ds.DS)
}

// FirstTriplet returns first triplet.
// 	Examples:
// 	Empty.FirstTriplet() // error ErrParse
// 	Zero.FirstTriplet() // triplet.Triplet{0,0,0}
// 	New(1, 2, 3, 4, 5, 6).FirstTriplet() // triplet.Triplet{4,5,6}
// 	New(1, 2, 3, 4, 5, 6, 7).FirstTriplet() // triplet.Triplet{5,6,7}
func (ds DigitString) FirstTriplet() (t triplet.Triplet, err error) {
	triplets, err := ds.Split()
	if err != nil {
		return t, err
	}

	return triplets[0], nil
}

// Split returns valid []Triplet or returns an error
// 	Examples:
// 	Empty.Split() // error ErrParse
// 	Zero.Split() // triplet.Triplet{0,0,0}
// 	New(1).Split() // triplet.Triplet{0,0,1}
// 	New(1, 2, 3, 4, 5, 6).Split() // triplet.Triplet{1,2,3}, triplet.Triplet{4,5,6}
// 	New(1, 2, 3, 4, 5, 6, 7).Split() // triplet.Triplet{0,0,1}, triplet.Triplet{2,3,4}, triplet.Triplet{5,6,7}
func (ds DigitString) Split() (triplet.Triplets, error) {
	if ds.IsZero() {
		return []triplet.Triplet{triplet.New(0, 0, 0)}, nil
	}

	if ds.IsEmpty() {
		return nil, fmt.Errorf("%w: cannot split empty ds", ErrParse)
	}

	err := ds.Validate()
	if err != nil {
		return nil, err
	}

	result := make([]triplet.Triplet, 0, 1+len(ds.DS)/3)

	for i := len(ds.DS) - 1; i >= 0; i -= 3 {
		t := triplet.Triplet{}

		if i-2 >= 0 {
			t.SetHundreds(ds.DS[i-2])
		}

		if i-1 >= 0 {
			t.SetTens(ds.DS[i-1])
		}

		t.SetUnits(ds.DS[i])

		result = append(result, t)
	}

	return result, nil
}
