package lang

import "errors"

type Lang uint32

var (
	ErrBadLanguage = errors.New("bad language")
)

const (
	Unknown Lang = iota
	KZ
	RU
	EN
)

const (
	CodeKZ = "KZ"
	CodeRU = "RU"
	CodeEN = "EN"
)

// String returns language code or empty string
func (l Lang) String() string {
	switch l {
	case KZ:
		return CodeKZ
	case RU:
		return CodeRU
	case EN:
		return CodeEN
	case Unknown:
		return ""
	default:
		return ""
	}
}
