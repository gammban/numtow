package curtow

import (
	"errors"
	"strings"
	"testing"
)

func TestFloat64(t *testing.T) {
	for _, v := range testCases {
		v := v
		t.Run(v.giveAmountString, func(tt *testing.T) {
			gotAmount, gotErr := Float64(v.giveAmountFloat64, v.giveLang, v.giveOpts...)
			if !errors.Is(gotErr, v.wantErr) {
				tt.Errorf("%s: \nexp: '%s' \ngot: '%s'", v.giveAmountString, v.wantErr, gotErr)
				return
			}

			if !strings.EqualFold(gotAmount, v.wantAmount) {
				tt.Errorf("%s: \nexp: '%s' \ngot: '%s'", v.giveAmountString, v.wantAmount, gotAmount)
			}
		})
	}
}

func TestMustFloat64(t *testing.T) {
	for _, v := range testCases {
		v := v
		t.Run(v.giveAmountString, func(tt *testing.T) {
			gotAmount := MustFloat64(v.giveAmountFloat64, v.giveLang, v.giveOpts...)
			if v.wantErr != nil && gotAmount != "" {
				tt.Errorf("expected empty string got %s", gotAmount)
				return
			}

			if !strings.EqualFold(gotAmount, v.wantAmount) {
				tt.Errorf("%s: \nexp: '%s' \ngot: '%s'", v.giveAmountString, v.wantAmount, gotAmount)
			}
		})
	}
}
