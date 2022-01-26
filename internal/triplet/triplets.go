package triplet

import (
	"strings"
)

type Triplets []Triplet

func (t Triplets) Len() int {
	return len(t)
}

func (t Triplets) IsZero() bool {
	if len(t) == 0 {
		return false
	}

	for i := range t {
		if !t[i].IsZero() {
			return false
		}
	}

	return true
}

func (t Triplets) String() string {
	s := strings.Builder{}

	for i := len(t) - 1; i >= 0; i-- {
		s.WriteString(t[i].String())
	}

	return s.String()
}
