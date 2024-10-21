package numtow

import (
	"fmt"

	"github.com/gammban/numtow/lang"
	"github.com/gammban/numtow/lang/en"
	"github.com/gammban/numtow/lang/ru"
	"github.com/gammban/numtow/lang/ru/gender"
)

func ExampleMustString_default() {
	fmt.Println(MustString("8691705", lang.EN, en.FormatDefault))
	// Output: eight million, six hundred and ninety-one thousand, seven hundred and five
}

func ExampleMustString_without_and() {
	fmt.Println(MustString("8691705", lang.EN, en.FormatWithoutAnd))
	// Output: eight million six hundred ninety-one thousand seven hundred five
}

func ExampleMustFloat64_default() {
	fmt.Println(MustFloat64(1, lang.RU))
	// Output: одна
}

func ExampleMustFloat64_female() {
	fmt.Println(MustFloat64(2, lang.RU))
	// Output: две
}

func ExampleMustFloat64_male() {
	fmt.Println(MustFloat64(2, lang.RU, ru.WithFmtGender(gender.Male)))
	// Output: два
}
