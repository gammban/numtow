package numtow

import (
	"errors"
	"strings"
	"testing"
)

func TestFloat64(t *testing.T) {
	for _, v := range testCaseDecimalNumbers {
		gotResult, gotErr := Float64(v.GiveFloat64, v.GiveLang, v.GiveOpts...)
		if !errors.Is(gotErr, v.WantErr) {
			t.Errorf("%s: \nexp: '%s' \ngot: '%s'", v.GiveDecimal, v.WantErr, gotErr)
		}

		if gotErr == nil {
			if !strings.EqualFold(gotResult, v.WantResult) {
				t.Errorf("%s: \nexp: '%s' \ngot: '%s'", v.GiveDecimal, v.WantResult, gotResult)
			}
		}
	}
}

func TestFloat64OrZero(t *testing.T) {
	for _, v := range testCaseDecimalNumbers {
		gotResult := Float64OrZero(v.GiveFloat64, v.GiveLang, v.GiveOpts...)
		if v.WantErr != nil && gotResult != "" {
			t.Errorf("%s: expected error", v.GiveDecimal)
		}

		if gotResult != "" {
			if !strings.EqualFold(gotResult, v.WantResult) {
				t.Errorf("%s: \nexp: '%s' \ngot: '%s'", v.GiveDecimal, v.WantResult, gotResult)
			}
		}
	}
}
