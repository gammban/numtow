package kz

import (
	"fmt"
	"math"
	"testing"

	"github.com/dantedenis/numtow"
	"github.com/dantedenis/numtow/lang"
	"github.com/dantedenis/numtow/lang/kz"
)

func TestKz(t *testing.T) {
	// convert number to kazakh using numtow package
	fmt.Println(numtow.MustString("1", lang.KZ)) // бір
	fmt.Println(numtow.MustString("2", lang.KZ)) // екі

	// convert number to kazakh words using kz package
	fmt.Println(kz.MustString("1"))                                                          // бір
	fmt.Println(kz.MustString("2"))                                                          // екі
	fmt.Println(kz.MustString("3", kz.FormatDefault...))                                     // үш
	fmt.Println(kz.MustString("4515.753"))                                                   // төрт мың бес жүз он бес бүтін мыңнан жеті жүз елу үш
	fmt.Println(kz.MustString("4515,753", kz.WithParseSep(',')))                             // төрт мың бес жүз он бес бүтін мыңнан жеті жүз елу үш
	fmt.Println(kz.MustString("4515,753", kz.WithParseSep(','), kz.WithFmtFracIgnore(true))) // төрт мың бес жүз он бес
	fmt.Println(kz.MustString("4515.753", kz.WithFmtFracUseDigits(true)))                    // төрт мың бес жүз он бес бүтін мыңнан 753

	fmt.Println(kz.MustString("1000000"))       // бір миллион
	fmt.Println(kz.MustString("2000000"))       // екі миллион
	fmt.Println(kz.MustString("9999999"))       // тоғыз миллион тоғыз жүз тоқсан тоғыз мың тоғыз жүз тоқсан тоғыз
	fmt.Println(kz.MustString("10000000"))      // он миллион
	fmt.Println(kz.MustString("100000000"))     // жүз миллион
	fmt.Println(kz.MustString("1000000000"))    // бір миллиард
	fmt.Println(kz.MustString("1000000000000")) // бір триллион

	fmt.Println(kz.MustString("0.1"))      // нөл бүтін оннан бір
	fmt.Println(kz.MustString(".9"))       // нөл бүтін оннан тоғыз
	fmt.Println(kz.MustString("-5.47"))    // минус бес бүтін жүзден қырық жеті
	fmt.Println(kz.MustString("2.125"))    // екі бүтін мыңнан жүз жиырма бес
	fmt.Println(kz.MustString("9.1256"))   // тоғыз бүтін он мыңнан бір мың екі жүз елу алты
	fmt.Println(kz.MustString("15678.54")) // он бес мың алты жүз жетпіс сегіз бүтін жүзден елу төрт

	_, err := kz.Float64(math.NaN())
	if err != nil {
		fmt.Println(err) // parse number error: float64 is NaN
	}

	_, err = kz.String("bad")
	if err != nil {
		fmt.Println(err) // parse number error
	}

	result, err := kz.String("456")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result) // төрт жүз елу алты

	var MyFormat = []kz.OptFunc{
		kz.WithParseSep(','),
		kz.WithFmtFracUseDigits(true),
	}

	fmt.Println(kz.MustString("15,98", MyFormat...)) // он бес бүтін жүзден 98
	fmt.Println(kz.MustString("100", MyFormat...))   // жүз
}
