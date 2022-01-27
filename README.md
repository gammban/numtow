# numtow

[![GoDoc](https://godoc.org/github.com/gammban/numtow?status.svg)](https://godoc.org/github.com/gammban/numtow)
[![Go Report Card](https://goreportcard.com/badge/github.com/gammban/numtow)](https://goreportcard.com/report/github.com/gammban/numtow)
[![codecov](https://codecov.io/gh/gammban/numtow/branch/main/graph/badge.svg)](https://codecov.io/gh/gammban/numtow)

golang library to convert number to words, currencies to words. 

## Import

```shell
go get github.com/gammban/numtow
```

## Usage

```go
import "github.com/gammban/numtow" 

...

// convert examples using numtow package
fmt.Println(numtow.MustString("12", lang.EN))                           // twelve
fmt.Println(numtow.MustString("8691705", lang.EN, en.FormatDefault))    // eight million, six hundred and ninety-one thousand, seven hundred and five
fmt.Println(numtow.MustString("8691705", lang.EN, en.FormatWithoutAnd)) // eight million six hundred ninety-one thousand seven hundred five
fmt.Println(numtow.MustString("1", lang.KZ)) // бір
fmt.Println(numtow.MustString("2", lang.KZ)) // екі
fmt.Println(numtow.MustString("1", lang.RU)) // одна
fmt.Println(numtow.MustString("2", lang.RU)) // две

// convert examples using directly language package
fmt.Println(en.MustString("1"))                                                     // one
fmt.Println(en.MustString("22"))                                                    // twenty-two
fmt.Println(en.MustString("-1234567.89", en.FormatDefault...))                      // minus one million, two hundred and thirty-four thousand, five hundred and sixty-seven point eighty-nine
fmt.Println(en.MustString("-1234567.89", en.FormatWithoutAnd...))                   // minus one million two hundred thirty-four thousand five hundred sixty-seven point eighty-nine
fmt.Println(en.MustString("1000000"))                                               // one million
fmt.Println(en.MustString("2000000"))                                               // two million
fmt.Println(en.MustString("1450926"))                                               // one million, four hundred and fifty thousand, nine hundred and twenty-six
fmt.Println(en.MustString("1450926", en.WithFmtGroupSep("")))                       // one million four hundred and fifty thousand nine hundred and twenty-six
fmt.Println(en.MustString("1450926", en.WithFmtAndSep("")))                         // one million, four hundred fifty thousand, nine hundred twenty-six
fmt.Println(en.MustString("1450926", en.WithFmtAndSep(""), en.WithFmtGroupSep(""))) // one million four hundred fifty thousand nine hundred twenty-six
fmt.Println(kz.MustString("4515.753"))                                              // төрт мың бес жүз он бес бүтін мыңнан жеті жүз елу үш
fmt.Println(ru.MustString("123.457", ru.FormatDefault...))                          // сто двадцать три целых четыреста пятьдесят семь тысячных

_, err := en.String("bad")
if err != nil {
    fmt.Println(err) // parse number error
}

res, err := kz.String("12")
fmt.Println(res) // он екi

res, err = ru.Float64(12)
fmt.Println(res) // двенадцать

res, err = ru.Float64(1, ru.WithGender(gender.Male))
fmt.Println(res) // один

res, err = ru.Int64(1, ru.WithGender(gender.Female))
fmt.Println(res) // одна

res, err = ru.String("1", ru.WithGender(gender.Neuter))
fmt.Println(res) // одно

```

more examples in example directory

## Default and customized formats

each language package has own default formats

```go
fmt.Println(numtow.MustString("8691705", lang.EN, en.FormatDefault))    // eight million, six hundred and ninety-one thousand, seven hundred and five
fmt.Println(numtow.MustString("8691705", lang.EN, en.FormatWithoutAnd)) // eight million six hundred ninety-one thousand seven hundred five

fmt.Println(en.MustString("-1234567.89", en.FormatDefault...))                      // minus one million, two hundred and thirty-four thousand, five hundred and sixty-seven point eighty-nine
fmt.Println(en.MustString("-1234567.89", en.FormatWithoutAnd...))                   // minus one million two hundred thirty-four thousand five hundred sixty-seven point eighty-nine

```

build your own format
```go
// using your own format
var MyFormat = []en.OptFunc{
    en.WithParseSep(','),
    en.WithFmtFracUseDigits(true),
    en.WithFmtAndSep(""),
    en.WithFmtGroupSep(""),
}
fmt.Println(en.MustString("12,54", MyFormat...)) // twelve point 54
```

## Convert currencies

```go
import "github.com/gammban/numtow/curtow"
...
// convert currency to english words using curtow package
fmt.Println(curtow.MustString("12", lang.EN, en.WithCur(cur.USD)))                           // twelve dollars and 00 cents
fmt.Println(curtow.MustString("12", lang.EN, en.WithCur(cur.USD), en.WithCurIgnoreMU(true))) // twelve dollars
fmt.Println(curtow.MustString("12", lang.EN, en.WithCur(cur.USD), en.WithCurConvMU(true)))   // twelve dollars and zero cents

res, err := en.CurrencyString("53241", en.WithCur(cur.EUR))
if err != nil {
	fmt.Println(err)
}
fmt.Println(res) // fifty-three thousand, two hundred and forty-one euros and 00 cents

fmt.Println(curtow.MustString("12", lang.KZ, kz.WithCur(cur.KZT)))                           // он екі теңге 00 тиын
fmt.Println(curtow.MustString("12", lang.KZ, kz.WithCur(cur.KZT), kz.WithCurIgnoreMU(true))) // он екі теңге
fmt.Println(curtow.MustString("12", lang.KZ, kz.WithCur(cur.KZT), kz.WithCurConvMU(true)))   // он екі теңге нөл тиын
fmt.Println(curtow.MustFloat64(25.79, lang.KZ, kz.WithCur(cur.USD), kz.WithCurConvMU(true))) // жиырма бес АҚШ доллары жетпіс тоғыз цент

res, err := kz.CurrencyString("53241", kz.WithCur(cur.KZT))
fmt.Println(res) // елу үш мың екі жүз қырық бір теңге 00 тиын

res, err = kz.CurrencyFloat64(125.53, kz.WithCur(cur.KZT))
fmt.Println(res) // жүз жиырма бес теңге 53 тиын
```

## Supported currencies

- USD
- EUR
- KZT
- RUB

## Supported languages

- Kazakh
- English
- Russian

## License

MIT