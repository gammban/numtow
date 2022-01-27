package gender

import "errors"

var (
	ErrBadGender = errors.New("bad gender")
)

// Gender - gender for russian language.
type Gender uint8

const (
	// Unknown - неизвестный род
	Unknown Gender = iota
	// Male - мужской род
	Male
	// Female - женский род
	Female
	// Neuter - средний род
	Neuter
)

// Validate returns ErrBadGender when Gender is not valid.
func (g Gender) Validate() error {
	switch g {
	case Male, Female, Neuter:
		return nil
	case Unknown:
		return ErrBadGender
	default:
		return ErrBadGender
	}
}
