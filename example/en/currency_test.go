package en

import (
	"fmt"
	"testing"

	"github.com/gammban/numtow/curtow"
	"github.com/gammban/numtow/curtow/cur"
	"github.com/gammban/numtow/lang"
	"github.com/gammban/numtow/lang/en"
	"github.com/gammban/numtow/lang/ru"
)

func TestCurrencyEN(_ *testing.T) {
	// convert currency to english words using curtow package
	fmt.Println(curtow.MustString("12", lang.EN, en.WithCur(cur.USD)))                           // twelve dollars and 00 cents
	fmt.Println(curtow.MustString("12", lang.EN, en.WithCur(cur.USD), en.WithCurIgnoreMU(true))) // twelve dollars
	fmt.Println(curtow.MustString("12", lang.EN, en.WithCur(cur.USD), en.WithCurConvMU(true)))   // twelve dollars and zero cents
	fmt.Println(curtow.MustString("12.12", lang.EN, en.WithCur(cur.GBP)))
	fmt.Println(curtow.MustString("12.12", lang.RU, ru.WithCur(cur.GBP)))
	fmt.Println(curtow.MustString("12.03", lang.EN, en.WithCur(cur.GBP)))
	fmt.Println(curtow.MustString("12.03", lang.RU, ru.WithCur(cur.GBP)))
	fmt.Println(curtow.MustString(fmt.Sprintf("%.2f", 12.031), lang.EN, en.WithCur(cur.GBP)))
	fmt.Println(curtow.MustString(fmt.Sprintf("%.2f", 12.001), lang.EN, en.WithCur(cur.GBP)))

	res, err := en.CurrencyString("53241", en.WithCur(cur.EUR))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res) // fifty-three thousand, two hundred and forty-one euros and 00 eurocents
}
