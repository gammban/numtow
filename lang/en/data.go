package en

import "github.com/dantedenis/numtow/curtow/cur"

const (
	minus       = "minus"
	negative    = "negative"
	zero        = "zero"
	sep         = " "
	hundred     = "hundred"
	integerPart = "point"
	and         = "and"
)

type currencyInfo struct {
	Singular     string
	Plural       string
	UnitSingular string
	UnitPlural   string
}

//nolint:gochecknoglobals
var (
	megs  = [16]string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion", "septillion", "octillion", "nonillion", "decillion", "undecillion", "duodecillion", "tredecillion", "quattuordecillion"}
	units = [10]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	tens  = [10]string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	teens = [10]string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

	curNamesEN = map[cur.Currency]currencyInfo{
		cur.KZT: {
			Singular:     "tenge",
			Plural:       "tenge",
			UnitSingular: "tiyn",
			UnitPlural:   "tiyn",
		},
		cur.RUB: {
			Singular:     "ruble",
			Plural:       "rubles",
			UnitSingular: "kopeck",
			UnitPlural:   "kopecks",
		},
		cur.USD: {
			Singular:     "dollar",
			Plural:       "dollars",
			UnitSingular: "cent",
			UnitPlural:   "cents",
		},
		cur.EUR: {
			Singular:     "euro",
			Plural:       "euros",
			UnitSingular: "cent",
			UnitPlural:   "cents",
		},
		cur.GBP: {
			Singular:     "pound",
			Plural:       "pounds",
			UnitSingular: "penny",
			UnitPlural:   "pence",
		},
	}
)
