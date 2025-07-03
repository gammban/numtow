package curtow

import (
	"errors"
	"strings"
	"testing"
)

func TestFloat64(t *testing.T) {
	for _, v := range testCases {
		gotAmount, gotErr := Float64(v.giveAmountFloat64, v.giveLang, v.giveOpts...)
		if !errors.Is(gotErr, v.wantErr) {
			t.Errorf("%s: \nexp: '%s' \ngot: '%s'", v.giveAmountString, v.wantErr, gotErr)
			return
		}

		if !strings.EqualFold(gotAmount, v.wantAmount) {
			t.Errorf("%s: \nexp: '%s' \ngot: '%s'", v.giveAmountString, v.wantAmount, gotAmount)
		}
	}
}

func TestFloat64OrZero(t *testing.T) {
	for _, v := range testCases {
		gotAmount := Float64OrZero(v.giveAmountFloat64, v.giveLang, v.giveOpts...)
		if v.wantErr != nil && gotAmount != "" {
			t.Errorf("expected empty string got %s", gotAmount)
			return
		}

		if !strings.EqualFold(gotAmount, v.wantAmount) {
			t.Errorf("%s: \nexp: '%s' \ngot: '%s'", v.giveAmountString, v.wantAmount, gotAmount)
		}
	}
}
