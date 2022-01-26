package curtow

import (
	"errors"
	"math"
	"strings"
	"testing"

	"github.com/gammban/numtow/curtow/cur"
	"github.com/gammban/numtow/internal/ds"
	"github.com/gammban/numtow/lang/kz"
	"github.com/gammban/numtow/lang/ru"

	"github.com/gammban/numtow/lang"
)

//nolint:gochecknoglobals
var testCases = []struct {
	giveAmountString  string
	giveAmountFloat64 float64
	giveLang          lang.Lang
	giveOpts          []interface{}
	wantAmount        string
	wantErr           error
}{
	{
		giveAmountString: "0", giveAmountFloat64: 0, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.USD), ru.WithCurConvMU(false), ru.WithCurIgnoreMU(false)},
		wantAmount: "ноль долларов США 00 центов",
	},
	{
		giveAmountString: "0.22", giveAmountFloat64: 0.22, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.USD), ru.WithCurConvMU(false), ru.WithCurIgnoreMU(false)},
		wantAmount: "Ноль долларов США 22 цента",
	},
	{
		giveAmountString: "0.22", giveAmountFloat64: 0.22, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.USD), ru.WithCurConvMU(false), ru.WithCurIgnoreMU(true)},
		wantAmount: "Ноль долларов США",
	},
	{
		giveAmountString: "0.22", giveAmountFloat64: 0.22, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.USD), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Ноль долларов США двадцать два цента",
	},
	{
		giveAmountString: "0.223", giveAmountFloat64: 0.223, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.USD), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Ноль долларов США двадцать два цента",
	},
	{
		giveAmountString: "0.229", giveAmountFloat64: 0.229, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.USD), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Ноль долларов США двадцать два цента",
	},
	{
		giveAmountString: "1.1", giveAmountFloat64: 1.1, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.USD), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Один доллар США десять центов",
	},
	{
		giveAmountString: "1.10", giveAmountFloat64: 1.10, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.USD), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Один доллар США десять центов",
	},
	{
		giveAmountString: "bad", giveAmountFloat64: math.NaN(), giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.USD), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "", wantErr: ds.ErrParse,
	},
	{
		giveAmountString: "-35.42", giveAmountFloat64: -35.42, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.USD), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Минус тридцать пять долларов США сорок два цента",
	},
	{
		giveAmountString: "-35.42", giveAmountFloat64: -35.42, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.USD), ru.WithCurConvMU(false), ru.WithCurIgnoreMU(false)},
		wantAmount: "Минус тридцать пять долларов США 42 цента",
	},
	{
		giveAmountString: "1000000", giveAmountFloat64: 1000000, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.USD), ru.WithCurConvMU(false), ru.WithCurIgnoreMU(false)},
		wantAmount: "Один миллион долларов США 00 центов",
	},
	{
		giveAmountString: "1000000", giveAmountFloat64: 1000000, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.USD), ru.WithCurConvMU(false), ru.WithCurIgnoreMU(true)},
		wantAmount: "Один миллион долларов США",
	},
	{
		giveAmountString: "1999999.99", giveAmountFloat64: 1999999.99, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.USD), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Один миллион девятьсот девяносто девять тысяч девятьсот девяносто девять долларов США девяносто девять центов",
	},
	{
		giveAmountString: "235.75", giveAmountFloat64: 235.75, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.EUR), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Двести тридцать пять евро семьдесят пять евроцентов",
	},
	{
		giveAmountString: "-181.02", giveAmountFloat64: -181.02, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.EUR), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Минус сто восемьдесят один евро два евроцента",
	},
	{
		giveAmountString: "4541782354.87", giveAmountFloat64: 4541782354.87, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.EUR), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Четыре миллиарда пятьсот сорок один миллион семьсот восемьдесят две тысячи триста пятьдесят четыре евро восемьдесят семь евроцентов",
	},
	{
		giveAmountString: "450000", giveAmountFloat64: 450000, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.EUR), ru.WithCurConvMU(false), ru.WithCurIgnoreMU(false)},
		wantAmount: "Четыреста пятьдесят тысяч евро 00 евроцентов",
	},
	{
		giveAmountString: "1.02", giveAmountFloat64: 1.02, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Один рубль две копейки",
	},
	{
		giveAmountString: "2.01", giveAmountFloat64: 2.01, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Два рубля одна копейка",
	},
	{
		giveAmountString: "3.04", giveAmountFloat64: 3.04, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Три рубля четыре копейки",
	},
	{
		giveAmountString: "4.5", giveAmountFloat64: 4.5, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Четыре рубля пятьдесят копеек",
	},
	{
		giveAmountString: "5.69", giveAmountFloat64: 5.69, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Пять рублей шестьдесят девять копеек",
	},
	{
		giveAmountString: "-10.11", giveAmountFloat64: -10.11, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Минус десять рублей одиннадцать копеек",
	},
	{
		giveAmountString: "12.45", giveAmountFloat64: 12.45, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(true)},
		wantAmount: "Двенадцать рублей",
	},
	{
		giveAmountString: "13.14", giveAmountFloat64: 13.14, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Тринадцать рублей четырнадцать копеек",
	},
	{
		giveAmountString: "315.16", giveAmountFloat64: 315.16, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Триста пятнадцать рублей шестнадцать копеек",
	},
	{
		giveAmountString: "5617.18", giveAmountFloat64: 5617.18, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Пять тысяч шестьсот семнадцать рублей восемнадцать копеек",
	},
	{
		giveAmountString: "100", giveAmountFloat64: 100, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(false), ru.WithCurIgnoreMU(false)},
		wantAmount: "Сто рублей 00 копеек",
	},
	{
		giveAmountString: "100", giveAmountFloat64: 100, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Сто рублей ноль копеек",
	},
	{
		giveAmountString: "10.00000", giveAmountFloat64: 10.00000, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Десять рублей ноль копеек",
	},
	{
		giveAmountString: "1000", giveAmountFloat64: 1000, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(true)},
		wantAmount: "Одна тысяча рублей",
	},
	{
		giveAmountString: "10000", giveAmountFloat64: 10000, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(false), ru.WithCurIgnoreMU(false)},
		wantAmount: "Десять тысяч рублей 00 копеек",
	},
	{
		giveAmountString: "100000", giveAmountFloat64: 100000, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Сто тысяч рублей ноль копеек",
	},
	{
		giveAmountString: "10000000", giveAmountFloat64: 10000000, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Десять миллионов рублей ноль копеек",
	},
	{
		giveAmountString: "100000000", giveAmountFloat64: 100000000, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Сто миллионов рублей ноль копеек",
	},
	{
		giveAmountString: "1000000000", giveAmountFloat64: 1000000000, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Один миллиард рублей ноль копеек",
	},
	{
		giveAmountString: "2000000000", giveAmountFloat64: 2000000000, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.RUB), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Два миллиарда рублей ноль копеек",
	},
	{
		giveAmountString: "1.01", giveAmountFloat64: 1.01, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.KZT), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Одна тенге одна тиын",
	},
	{
		giveAmountString: "3.02", giveAmountFloat64: 3.02, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.KZT), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Три тенге две тиын",
	},
	{
		giveAmountString: "2.45", giveAmountFloat64: 2.45, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.KZT), ru.WithCurConvMU(true), ru.WithCurIgnoreMU(false)},
		wantAmount: "Две тенге сорок пять тиын",
	},
	{
		giveAmountString: "125545215.45", giveAmountFloat64: 125545215.45, giveLang: lang.RU, giveOpts: []interface{}{ru.WithCur(cur.KZT), ru.WithCurConvMU(false), ru.WithCurIgnoreMU(false)},
		wantAmount: "Сто двадцать пять миллионов пятьсот сорок пять тысяч двести пятнадцать тенге 45 тиын",
	},
	{
		giveAmountString: "187.51", giveAmountFloat64: 187.51, giveLang: lang.KZ, giveOpts: []interface{}{kz.WithCur(cur.KZT), kz.WithCurConvMU(true), kz.WithCurIgnoreMU(false)},
		wantAmount: "жүз сексен жеті теңге елу бір тиын",
	},
	{
		giveAmountString: "-1.05", giveAmountFloat64: -1.05, giveLang: lang.KZ, giveOpts: []interface{}{kz.WithCur(cur.KZT), kz.WithCurConvMU(true), kz.WithCurIgnoreMU(false)},
		wantAmount: "минус бір теңге бес тиын",
	},
	{
		giveAmountString: "92.5059", giveAmountFloat64: 92.5059, giveLang: lang.KZ, giveOpts: []interface{}{kz.WithCur(cur.KZT), kz.WithCurConvMU(true), kz.WithCurIgnoreMU(false)},
		wantAmount: "тоқсан екі теңге елу тиын",
	},
	{
		giveAmountString: "92.5059", giveAmountFloat64: 92.5059, giveLang: lang.KZ, giveOpts: []interface{}{kz.WithCur(cur.KZT), kz.WithCurConvMU(true), kz.WithCurIgnoreMU(false)},
		wantAmount: "тоқсан екі теңге елу тиын",
	},
	{
		giveAmountString: "bad", giveAmountFloat64: math.Inf(0), giveLang: lang.KZ, giveOpts: []interface{}{kz.WithCur(cur.KZT), kz.WithCurConvMU(true), kz.WithCurIgnoreMU(false)},
		wantAmount: "", wantErr: ds.ErrParse,
	},
	{
		giveAmountString: "100", giveAmountFloat64: 100, giveLang: lang.KZ, giveOpts: []interface{}{kz.WithCur(cur.KZT), kz.WithCurConvMU(true), kz.WithCurIgnoreMU(false)},
		wantAmount: "жүз теңге нөл тиын",
	},
	{
		giveAmountString: "1000", giveAmountFloat64: 1000, giveLang: lang.KZ, giveOpts: []interface{}{kz.WithCur(cur.KZT), kz.WithCurConvMU(true), kz.WithCurIgnoreMU(true)},
		wantAmount: "бір мың теңге",
	},
	{
		giveAmountString: "10000", giveAmountFloat64: 10000, giveLang: lang.KZ, giveOpts: []interface{}{kz.WithCur(cur.KZT), kz.WithCurConvMU(true), kz.WithCurIgnoreMU(true)},
		wantAmount: "он мың теңге",
	},
	{
		giveAmountString: "10000", giveAmountFloat64: 10000, giveLang: lang.KZ, giveOpts: []interface{}{kz.WithCur(cur.USD), kz.WithCurConvMU(true), kz.WithCurIgnoreMU(true)},
		wantAmount: "он мың АҚШ доллары",
	},
	{
		giveAmountString: "5956.5", giveAmountFloat64: 5956.5, giveLang: lang.KZ, giveOpts: []interface{}{kz.WithCur(cur.USD), kz.WithCurConvMU(true), kz.WithCurIgnoreMU(false)},
		wantAmount: "бес мың тоғыз жүз елу алты АҚШ доллары елу цент",
	},
	{
		giveAmountString: "964913.39", giveAmountFloat64: 964913.39, giveLang: lang.KZ, giveOpts: []interface{}{kz.WithCur(cur.EUR), kz.WithCurConvMU(true), kz.WithCurIgnoreMU(false)},
		wantAmount: "тоғыз жүз алпыс төрт мың тоғыз жүз он үш еуро отыз тоғыз евроцент",
	},
	{
		giveAmountString: "4.39", giveAmountFloat64: 4.39, giveLang: lang.KZ, giveOpts: []interface{}{kz.WithCur(cur.RUB), kz.WithCurConvMU(true), kz.WithCurIgnoreMU(false)},
		wantAmount: "төрт рубль отыз тоғыз тиын",
	},
	{
		giveAmountString: "4.39", giveAmountFloat64: 4.39, giveLang: lang.Unknown, giveOpts: []interface{}{kz.WithCur(cur.RUB), kz.WithCurConvMU(true), kz.WithCurIgnoreMU(false)},
		wantAmount: "", wantErr: lang.ErrBadLanguage,
	},
	{
		giveAmountString: "4.39", giveAmountFloat64: 4.39, giveLang: lang.Lang(10), giveOpts: []interface{}{kz.WithCur(cur.RUB), kz.WithCurConvMU(true), kz.WithCurIgnoreMU(false)},
		wantAmount: "", wantErr: lang.ErrBadLanguage,
	},
}

func TestString(t *testing.T) {
	for _, v := range testCases {
		v := v
		t.Run(v.giveAmountString, func(tt *testing.T) {
			gotAmount, gotErr := String(v.giveAmountString, v.giveLang, v.giveOpts...)
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

func TestMustString(t *testing.T) {
	for _, v := range testCases {
		v := v
		t.Run(v.giveAmountString, func(tt *testing.T) {
			gotAmount := MustString(v.giveAmountString, v.giveLang, v.giveOpts...)
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
