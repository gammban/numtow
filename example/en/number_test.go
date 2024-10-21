package en

import (
	"fmt"
	"testing"

	"github.com/dantedenis/numtow"
	"github.com/dantedenis/numtow/lang"
	"github.com/dantedenis/numtow/lang/en"
)

func TestEn(t *testing.T) {
	// convert number to english using numtow package
	fmt.Println(numtow.MustString("12", lang.EN))                           // twelve
	fmt.Println(numtow.MustString("8691705", lang.EN, en.FormatDefault))    // eight million, six hundred and ninety-one thousand, seven hundred and five
	fmt.Println(numtow.MustString("8691705", lang.EN, en.FormatWithoutAnd)) // eight million six hundred ninety-one thousand seven hundred five

	// convert number to english words using en package
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

	_, err := en.String("bad")
	if err != nil {
		fmt.Println(err) // parse number error
	}

	var MyFormat = []en.OptFunc{
		en.WithParseSep(','),
		en.WithFmtFracUseDigits(true),
		en.WithFmtAndSep(""),
		en.WithFmtGroupSep(""),
	}

	fmt.Println(en.MustString("12,54", MyFormat...)) // twelve point 54
}
