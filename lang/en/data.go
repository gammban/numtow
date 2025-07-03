package en

import "github.com/gammban/numtow/curtow/cur"

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
			UnitSingular: "eurocent",
			UnitPlural:   "eurocents",
		},
		cur.GBP: {
			Singular:     "pound",
			Plural:       "pounds",
			UnitSingular: "penny",
			UnitPlural:   "pence",
		},
		cur.CHF: {
			Singular:     "franc",
			Plural:       "francs",
			UnitSingular: "rappen",
			UnitPlural:   "rappen",
		},
		cur.CAD: {
			Singular:     "Canadian dollar",
			Plural:       "Canadian dollars",
			UnitSingular: "cent",
			UnitPlural:   "cents",
		},
		cur.JPY: {
			Singular:     "yen",
			Plural:       "yen",
			UnitSingular: "sen",
			UnitPlural:   "sen",
		},
		cur.CNY: {
			Singular:     "yuan",
			Plural:       "yuan",
			UnitSingular: "jiao",
			UnitPlural:   "jiao",
		},
		cur.AUD: {
			Singular:     "Australian dollar",
			Plural:       "Australian dollars",
			UnitSingular: "cent",
			UnitPlural:   "cents",
		},
		cur.AED: {
			Singular:     "dirham",
			Plural:       "dirhams",
			UnitSingular: "fils",
			UnitPlural:   "fils",
		},
		cur.TRY: {
			Singular:     "lira",
			Plural:       "lira",
			UnitSingular: "kurush",
			UnitPlural:   "kurush",
		},
		cur.KGS: {
			Singular:     "som",
			Plural:       "som",
			UnitSingular: "tiyin",
			UnitPlural:   "tiyin",
		},
		cur.ZAR: {
			Singular:     "rand",
			Plural:       "rand",
			UnitSingular: "cent",
			UnitPlural:   "cents",
		},
		cur.NOK: {
			Singular:     "krone",
			Plural:       "krone",
			UnitSingular: "øre",
			UnitPlural:   "øre",
		},
		cur.SEK: {
			Singular:     "krone",
			Plural:       "krone",
			UnitSingular: "øre",
			UnitPlural:   "øre",
		},
	}
)
