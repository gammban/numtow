package numtow

import (
	"errors"
	"strings"
	"testing"

	"github.com/gammban/numtow/lang"
	"github.com/gammban/numtow/lang/kz"
	"github.com/gammban/numtow/lang/ru"
	"github.com/gammban/numtow/lang/ru/gender"
)

//nolint:gochecknoglobals
var testCaseIntNumbers = []struct {
	GiveInt    string
	GiveInt64  int64
	GiveLang   lang.Lang
	GiveOpts   []interface{}
	WantResult string
	WantErr    error
}{
	{
		GiveInt: "1", GiveInt64: 1, GiveLang: lang.RU,
		WantResult: "Одна",
	},
	{
		GiveInt: "1", GiveInt64: 1, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtFracIgnore(true)},
		WantResult: "Одна",
	},
	{
		GiveInt: "1", GiveInt64: 1, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtGender(gender.Female), ru.WithFmtFracIgnore(true)},
		WantResult: "Одна",
	},
	{
		GiveInt: "1", GiveInt64: 1, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtGender(gender.Male), ru.WithFmtFracIgnore(true)},
		WantResult: "Один",
	},
	{
		GiveInt: "1", GiveInt64: 1, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtGender(gender.Neuter), ru.WithFmtFracIgnore(true)},
		WantResult: "Одно",
	},
	{
		GiveInt: "1", GiveInt64: 1, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtGender(gender.Neuter)},
		WantResult: "Одно",
	},
	{
		GiveInt: "1541", GiveInt64: 1541, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtGender(gender.Male)},
		WantResult: "одна тысяча пятьсот сорок один",
	},
	{
		GiveInt: "1541", GiveInt64: 1541, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtGender(gender.Male)},
		WantResult: "одна тысяча пятьсот сорок один",
	},
	{
		GiveInt: "2", GiveInt64: 2, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtFracUseDigits(true)},
		WantResult: "Две",
	},
	{
		GiveInt: "2", GiveInt64: 2, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithFmtFracUseDigits(true), ru.WithFmtFracIgnore(true)},
		WantResult: "Две",
	},
	{
		GiveInt: "2", GiveInt64: 2, GiveLang: lang.RU,
		WantResult: "Две",
	},
	{
		GiveInt: "2", GiveInt64: 2, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithParseFracLen(2)},
		WantResult: "Две",
	},
	{
		GiveInt: "2", GiveInt64: 2, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithParseFracLen(3)},
		WantResult: "Две",
	},
	{
		GiveInt: "2", GiveInt64: 2, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithParseFracLen(0)},
		WantResult: "Две",
	},
	{
		GiveInt: "2", GiveInt64: 2, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithParseFracLen(3)},
		WantResult: "Две",
	},
	{
		GiveInt: "2", GiveInt64: 2, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithParseFracLen(3), ru.WithFmtFracIgnore(true), ru.WithFmtFracUseDigits(true)},
		WantResult: "Две",
	},
	{
		GiveInt: "2", GiveInt64: 2, GiveLang: lang.RU, GiveOpts: []interface{}{ru.WithParseFracLen(2), ru.WithParseSep(',')},
		WantResult: "Две",
	},
	{
		GiveInt: "1", GiveInt64: 1, GiveLang: lang.Lang(5),
		WantResult: "", WantErr: lang.ErrBadLanguage,
	},
	{
		GiveInt: "1", GiveInt64: 1, GiveLang: lang.Unknown,
		WantResult: "", WantErr: lang.ErrBadLanguage,
	},
	{
		GiveInt: "-88", GiveInt64: -88, GiveLang: lang.KZ,
		WantResult: "минус сексен сегіз",
	},
	{
		GiveInt: "-88", GiveInt64: -88, GiveLang: lang.KZ,
		WantResult: "минус сексен сегіз",
	},
	{
		GiveInt: "-88", GiveInt64: -88, GiveLang: lang.KZ, GiveOpts: []interface{}{kz.WithFmtFracUseDigits(true)},
		WantResult: "минус сексен сегіз",
	},
	{
		GiveInt: "-88", GiveInt64: -88, GiveLang: lang.KZ, GiveOpts: []interface{}{kz.WithParseFracLen(2)},
		WantResult: "минус сексен сегіз",
	},
	{
		GiveInt: "1", GiveInt64: 1, GiveLang: lang.KZ,
		WantResult: "бір",
	},
	{
		GiveInt: "1", GiveInt64: 1, GiveLang: lang.KZ, GiveOpts: []interface{}{kz.WithParseFracLen(2)},
		WantResult: "бір",
	},
	{
		GiveInt: "1", GiveInt64: 1, GiveLang: lang.KZ, GiveOpts: []interface{}{kz.WithParseFracLen(1)},
		WantResult: "бір",
	},
	{
		GiveInt: "1", GiveInt64: 1, GiveLang: lang.KZ, GiveOpts: []interface{}{kz.WithParseFracLen(1), kz.WithFmtFracUseDigits(true)},
		WantResult: "бір",
	},
	{
		GiveInt: "1", GiveInt64: 1, GiveLang: lang.KZ, GiveOpts: []interface{}{kz.WithFmtFracIgnore(true)},
		WantResult: "бір",
	},
}

func TestInt64(t *testing.T) {
	for _, v := range testCaseIntNumbers {
		gotResult, gotErr := Int64(v.GiveInt64, v.GiveLang, v.GiveOpts...)
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

func TestMustInt64(t *testing.T) {
	for _, v := range testCaseIntNumbers {
		gotResult := MustInt64(v.GiveInt64, v.GiveLang, v.GiveOpts...)
		if v.WantErr != nil && gotResult != "" {
			t.Errorf("%s: expected error", v.GiveInt)
		}

		if gotResult != "" {
			if !strings.EqualFold(gotResult, v.WantResult) {
				t.Errorf("%s: \nexp: '%s' \ngot: '%s'", v.GiveInt, v.WantResult, gotResult)
			}
		}
	}
}
