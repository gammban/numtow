package numtow

import (
	"errors"
	"math"
	"strings"
	"testing"

	"github.com/gammban/numtow/lang/en"

	"github.com/gammban/numtow/internal/ds"
	"github.com/gammban/numtow/lang"
	"github.com/gammban/numtow/lang/kz"
	"github.com/gammban/numtow/lang/ru"
)

//nolint:gochecknoglobals // test cases
var testCaseDecimalNumbers = []struct {
	GiveDecimal string
	GiveFloat64 float64
	GiveLang    lang.Lang
	GiveOpts    []interface{}
	WantResult  string
	WantErr     error
}{
	{
		GiveDecimal: "2", GiveFloat64: 2, GiveLang: lang.RU,
		WantResult: "две",
	},
	{
		GiveDecimal: "2", GiveFloat64: 2, GiveLang: lang.KZ,
		WantResult: "екі",
	},
	{
		GiveDecimal: "100", GiveFloat64: 100, GiveLang: lang.RU,
		WantResult: "сто",
	},
	{
		GiveDecimal: "-100", GiveFloat64: -100, GiveLang: lang.RU,
		WantResult: "минус сто",
	},
	{
		GiveDecimal: "1.01", GiveFloat64: 1.01, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtFracIgnore(false), ru.WithFmtFracUseDigits(false)},
		WantResult: "Одна целая одна сотая",
	},
	{
		GiveDecimal: "1.01", GiveFloat64: 1.01, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtFracIgnore(true), ru.WithFmtFracUseDigits(false)},
		WantResult: "Одна",
	},
	{
		GiveDecimal: "1.01", GiveFloat64: 1.01, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtGender(2), ru.WithFmtFracIgnore(true), ru.WithFmtFracUseDigits(false)},
		WantResult: "Одна",
	},
	{
		GiveDecimal: "1.01", GiveFloat64: 1.01, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtGender(1), ru.WithFmtFracIgnore(true), ru.WithFmtFracUseDigits(false)},
		WantResult: "Один",
	},
	{
		GiveDecimal: "1.01", GiveFloat64: 1.01, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtGender(3), ru.WithFmtFracIgnore(true), ru.WithFmtFracUseDigits(false)},
		WantResult: "Одно",
	},
	{
		GiveDecimal: "1.01", GiveFloat64: 1.01, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtGender(3), ru.WithFmtFracIgnore(false), ru.WithFmtFracUseDigits(false)},
		WantResult: "Одно целое одна сотая",
	},
	{
		GiveDecimal: "1541.1", GiveFloat64: 1541.1, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtGender(1), ru.WithFmtFracIgnore(false), ru.WithFmtFracUseDigits(false)},
		WantResult: "одна тысяча пятьсот сорок один целых одна десятая",
	},
	{
		GiveDecimal: "1541.1", GiveFloat64: 1541.1, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtGender(1), ru.WithFmtFracIgnore(false), ru.WithFmtFracUseDigits(false)},
		WantResult: "одна тысяча пятьсот сорок один целых одна десятая",
	},
	{
		GiveDecimal: "2.02", GiveFloat64: 2.02, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtFracIgnore(false), ru.WithFmtFracUseDigits(true)},
		WantResult: "Две целых 02 сотых",
	},
	{
		GiveDecimal: "2.02", GiveFloat64: 2.02, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtFracIgnore(true), ru.WithFmtFracUseDigits(true)},
		WantResult: "Две",
	},
	{
		GiveDecimal: "2.02", GiveFloat64: 2.02, GiveLang: lang.RU,
		WantResult: "Две целых две сотых",
	},
	{
		GiveDecimal: "2.02945", GiveFloat64: 2.02945, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithParseFracLen(2)},
		WantResult: "Две целых две сотых",
	},
	{
		GiveDecimal: "2.02945", GiveFloat64: 2.02945, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithParseFracLen(3), ru.WithFmtFracUseDigits(false)},
		WantResult: "Две целых двадцать девять тысячных",
	},
	{
		GiveDecimal: "2.02945", GiveFloat64: 2.02945, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithParseFracLen(0)},
		WantResult: "Две целых две тысячи девятьсот сорок пять стотысячных",
	},
	{
		GiveDecimal: "2.02945", GiveFloat64: 2.02945, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithParseFracLen(3), ru.WithFmtFracUseDigits(true)},
		WantResult: "Две целых 029 тысячных",
	},
	{
		GiveDecimal: "2.02945", GiveFloat64: 2.02945, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtFracIgnore(true), ru.WithFmtFracUseDigits(true), ru.WithParseFracLen(3)},
		WantResult: "Две",
	},
	{
		GiveDecimal: "2,02945", GiveFloat64: 2.02945, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithParseSep(',')},
		WantResult: "Две целых две тысячи девятьсот сорок пять стотысячных",
	},
	{
		GiveDecimal: "-2.4", GiveFloat64: -2.4, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithParseFracLen(2)},
		WantResult: "Минус две целых сорок сотых",
	},
	{
		GiveDecimal: "-2.04", GiveFloat64: -2.04, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithParseFracLen(2)},
		WantResult: "Минус две целых четыре сотых",
	},
	{
		GiveDecimal: "-2.04", GiveFloat64: -2.04, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithParseFracLen(1)},
		WantResult: "Минус две",
	},
	{
		GiveDecimal: "2,+02945", GiveFloat64: math.NaN(), GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithParseSep(',')},
		WantResult: "", WantErr: ds.ErrParse,
	},
	{
		GiveDecimal: "2.+02945", GiveFloat64: math.NaN(), GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtFracIgnore(true)},
		WantResult: "", WantErr: ds.ErrParse,
	},
	{
		GiveDecimal: "bad", GiveFloat64: math.Inf(0), GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtFracIgnore(true)},
		WantResult: "", WantErr: ds.ErrParse,
	},
	{
		GiveDecimal: "--2.4", GiveFloat64: math.Inf(-1), GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtFracIgnore(true)},
		WantResult: "", WantErr: ds.ErrParse,
	},
	{
		GiveDecimal: "--2.4", GiveFloat64: math.Inf(-1), GiveLang: lang.KZ, GiveOpts: []interface{}{kz.WithFmtFracIgnore(true)},
		WantResult: "", WantErr: ds.ErrParse,
	},
	{
		GiveDecimal: "-88.4", GiveFloat64: -88.4, GiveLang: lang.KZ, GiveOpts: []interface{}{kz.WithFmtFracIgnore(true)},
		WantResult: "минус сексен сегіз",
	},
	{
		GiveDecimal: "-88.4", GiveFloat64: -88.4, GiveLang: lang.KZ,
		WantResult: "минус сексен сегіз бүтін оннан төрт",
	},
	{
		GiveDecimal: "-88.4", GiveFloat64: -88.4, GiveLang: lang.KZ, GiveOpts: []interface{}{kz.WithFmtFracUseDigits(true)},
		WantResult: "минус сексен сегіз бүтін оннан 4",
	},
	{
		GiveDecimal: "-88.4", GiveFloat64: -88.4, GiveLang: lang.KZ, GiveOpts: []interface{}{kz.WithParseFracLen(2)},
		WantResult: "минус сексен сегіз бүтін жүзден қырық",
	},
	{
		GiveDecimal: "1.974", GiveFloat64: 1.974, GiveLang: lang.KZ,
		WantResult: "бір бүтін мыңнан тоғыз жүз жетпіс төрт",
	},
	{
		GiveDecimal: "1.974", GiveFloat64: 1.974, GiveLang: lang.KZ, GiveOpts: []interface{}{kz.WithParseFracLen(2)},
		WantResult: "бір бүтін жүзден тоқсан жеті",
	},
	{
		GiveDecimal: "1.974", GiveFloat64: 1.974, GiveLang: lang.KZ, GiveOpts: []interface{}{kz.WithParseFracLen(1)},
		WantResult: "бір бүтін оннан тоғыз",
	},
	{
		GiveDecimal: "1.974", GiveFloat64: 1.974, GiveLang: lang.KZ, GiveOpts: []interface{}{kz.WithParseFracLen(1), kz.WithFmtFracUseDigits(true)},
		WantResult: "бір бүтін оннан 9",
	},
	{
		GiveDecimal: "1.974", GiveFloat64: 1.974, GiveLang: lang.KZ, GiveOpts: []interface{}{kz.WithFmtFracIgnore(true)},
		WantResult: "бір",
	},
	{
		GiveDecimal: "1.974", GiveFloat64: 1.974, GiveLang: lang.Lang(5), GiveOpts: []interface{}{kz.WithFmtFracIgnore(true)},
		WantResult: "", WantErr: lang.ErrBadLanguage,
	},
	{
		GiveDecimal: "1.974", GiveFloat64: 1.974, GiveLang: lang.Unknown, GiveOpts: []interface{}{kz.WithFmtFracIgnore(true)},
		WantResult: "", WantErr: lang.ErrBadLanguage,
	},
	{
		GiveDecimal: "1,974", GiveFloat64: 1.974, GiveLang: lang.KZ, GiveOpts: []interface{}{kz.WithParseFracLen(1), kz.WithParseSep(',')},
		WantResult: "бір бүтін оннан тоғыз",
	},
	{
		GiveDecimal: "1.5", GiveFloat64: 1.5, GiveLang: lang.EN, GiveOpts: []interface{}{en.WithParseFracLen(1)},
		WantResult: "one point five",
	},
}

func TestString(t *testing.T) {
	for _, v := range testCaseDecimalNumbers {
		gotResult, gotErr := String(v.GiveDecimal, v.GiveLang, v.GiveOpts...)
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

func TestStringOrZero(t *testing.T) {
	for _, v := range testCaseDecimalNumbers {
		gotResult := StringOrZero(v.GiveDecimal, v.GiveLang, v.GiveOpts...)
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

func TestString_Int(t *testing.T) {
	for _, v := range testCaseIntNumbers {
		gotResult, gotErr := String(v.GiveInt, v.GiveLang, v.GiveOpts...)
		if !errors.Is(gotErr, v.WantErr) {
			t.Errorf("%s: \nexp: '%s' \ngot: '%s'", v.GiveInt, v.WantErr, gotErr)
		}

		if gotErr == nil {
			if !strings.EqualFold(gotResult, v.WantResult) {
				t.Errorf("%s: \nexp: '%s' \ngot: '%s'", v.GiveInt, v.WantResult, gotResult)
			}
		}
	}
}
