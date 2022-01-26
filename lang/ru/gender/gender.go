package gender

import "errors"

var (
	ErrBadGender = errors.New("bad gender")
)

// Gender - род
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
