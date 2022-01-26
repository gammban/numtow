package cur

import (
	"errors"
	"strconv"
	"strings"
)

type MinorUnits uint8

const (
	MinorUnitsUnknown MinorUnits = 255
	MinorUnits0       MinorUnits = 0
	MinorUnits2       MinorUnits = 2
	MinorUnits3       MinorUnits = 3
)

var (
	ErrBadMinorUnit = errors.New("bad currency minor unit")
)

const (
	base = 10
)

func (m MinorUnits) String() string {
	err := m.Validate()
	if err != nil {
		return ""
	}

	return strconv.FormatUint(uint64(m), base)
}

// Zero returns zero value for minor unit.
// 	Example: MinorUnits(2).Zero // "00"
func (m MinorUnits) Zero() string {
	if err := m.Validate(); err != nil {
		return ""
	}

	if m == MinorUnits0 {
		return ""
	}

	return strings.Repeat("0", int(m))
}

// Validate validates minor unit, on error returns ErrBadMinorUnit.
func (m MinorUnits) Validate() error {
	switch m {
	case MinorUnits0, MinorUnits2, MinorUnits3:
		return nil
	case MinorUnitsUnknown:
		return ErrBadMinorUnit
	default:
		return ErrBadMinorUnit
	}
}
