package cur

import (
	"errors"
	"reflect"
	"testing"
)

//nolint:gochecknoglobals
var testCaseString = []struct {
	give Currency
	want string
}{
	{give: 0, want: ""},
	{give: 123, want: ""},
	{give: 1, want: "KZT"},
	{give: 2, want: "USD"},
	{give: 3, want: "RUB"},
	{give: 4, want: "EUR"},
}

func TestLang_String(t *testing.T) {
	for _, v := range testCaseString {
		if got := v.give.String(); got != v.want {
			t.Errorf("got %s, expected %s", got, v.want)
		}
	}
}

//nolint:gochecknoglobals
var testCaseParseCur = []struct {
	give string
	want Currency
}{
	{want: 0, give: "Unknown"},
	{want: 0, give: ""},
	{want: 1, give: "KZT"},
	{want: 2, give: "USD"},
	{want: 3, give: "RUB"},
	{want: 4, give: "EUR"},
}

func TestLang_ParseCurrency(t *testing.T) {
	for _, v := range testCaseParseCur {
		if got := ParseCurrency(v.give); got != v.want {
			t.Errorf("got %s, expected %s", got, v.want)
		}
	}
}

//nolint:gochecknoglobals
var testCaseValidateCur = []struct {
	giveCurrency Currency
	wantErr      error
}{
	{giveCurrency: KZT},
	{giveCurrency: USD},
	{giveCurrency: RUB},
	{giveCurrency: EUR},
	{giveCurrency: Unknown, wantErr: ErrBadCurrency},
	{giveCurrency: Currency(10), wantErr: ErrBadCurrency},
}

func TestCurrency_Validate(t *testing.T) {
	for _, v := range testCaseValidateCur {
		if gotErr := v.giveCurrency.Validate(); !errors.Is(gotErr, v.wantErr) {
			t.Errorf("got %s, expected %s", gotErr, v.wantErr)
		}
	}
}

//nolint:gochecknoglobals
var testCaseMinorUnitsCur = []struct {
	giveCurrency   Currency
	wantMinorUnits MinorUnits
}{
	{giveCurrency: KZT, wantMinorUnits: MinorUnits2},
	{giveCurrency: USD, wantMinorUnits: MinorUnits2},
	{giveCurrency: RUB, wantMinorUnits: MinorUnits2},
	{giveCurrency: EUR, wantMinorUnits: MinorUnits2},
	{giveCurrency: Unknown, wantMinorUnits: MinorUnitsUnknown},
	{giveCurrency: Currency(10), wantMinorUnits: MinorUnitsUnknown},
}

func TestCurrency_MinorUnits(t *testing.T) {
	t.Parallel()

	for _, v := range testCaseMinorUnitsCur {
		if gotMU := v.giveCurrency.MinorUnits(); gotMU != v.wantMinorUnits {
			t.Errorf("got %s, expected %s", gotMU, v.wantMinorUnits)
		}
	}
}

//nolint:gochecknoglobals
var testCaseISO4217 = []struct {
	giveCurrency Currency
	want         *ISO4217
}{
	{giveCurrency: KZT, want: &ISO4217{
		Name:           "Tenge",
		AlphabeticCode: "KZT",
		NumericCode:    398,
		MinorUnits:     MinorUnits2,
	}},
	{giveCurrency: RUB, want: &ISO4217{
		Name:           "Russian Ruble",
		AlphabeticCode: "RUB",
		NumericCode:    643,
		MinorUnits:     MinorUnits2,
	}},
	{giveCurrency: USD, want: &ISO4217{
		Name:           "US Dollar",
		AlphabeticCode: "USD",
		NumericCode:    840,
		MinorUnits:     MinorUnits2,
	}},
	{giveCurrency: EUR, want: &ISO4217{
		Name:           "Euro",
		AlphabeticCode: "EUR",
		NumericCode:    978,
		MinorUnits:     MinorUnits2,
	}},
	{giveCurrency: Unknown, want: nil},
	{giveCurrency: Currency(10), want: nil},
}

func TestCurrency_ISO4217(t *testing.T) {
	for _, v := range testCaseISO4217 {
		if got := v.giveCurrency.ISO4217(); !reflect.DeepEqual(got, v.want) {
			t.Errorf("got %v, expected %v", got.MinorUnits, v.want)
		}
	}
}

func TestCurrency_castISO(t *testing.T) {
	for k, v := range details {
		if d := detailsIso[v.NumericCode]; d != k {
			t.Errorf("got: %v, expected %v", d, k)
		}
	}
}
