package kz

import (
	"fmt"
	"github.com/dantedenis/numtow/curtow"
	"github.com/dantedenis/numtow/curtow/cur"
	"github.com/dantedenis/numtow/lang"
	"github.com/dantedenis/numtow/lang/kz"
	"testing"
)

func TestCurrencyEN(t *testing.T) {
	// convert currency to kazakh words using curtow package
	fmt.Println(curtow.MustString("12", lang.KZ, kz.WithCur(cur.KZT)))                           // он екі теңге 00 тиын
	fmt.Println(curtow.MustString("12", lang.KZ, kz.WithCur(cur.KZT), kz.WithCurIgnoreMU(true))) // он екі теңге
	fmt.Println(curtow.MustString("12", lang.KZ, kz.WithCur(cur.KZT), kz.WithCurConvMU(true)))   // он екі теңге нөл тиын
	fmt.Println(curtow.MustFloat64(25.79, lang.KZ, kz.WithCur(cur.USD), kz.WithCurConvMU(true))) // жиырма бес АҚШ доллары жетпіс тоғыз цент

	res, err := kz.CurrencyString("53241", kz.WithCur(cur.KZT))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res) // елу үш мың екі жүз қырық бір теңге 00 тиын

	res, err = kz.CurrencyFloat64(125.53, kz.WithCur(cur.KZT))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res) // жүз жиырма бес теңге 53 тиын
}
