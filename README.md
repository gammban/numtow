# numtow

[![Go Report Card](https://goreportcard.com/badge/github.com/gammban/numtow)](https://goreportcard.com/report/github.com/gammban/numtow)

golang library to convert number to words. Supported languages: kazakh, english, russian.

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

## Supported languages

- Kazakh
- English
- Russian

## License

MIT