package cur

type ISO4217 struct {
	Name           string
	AlphabeticCode string
	NumericCode    uint16
	MinorUnits     MinorUnits
}

//nolint:gochecknoglobals,mnd,exhaustive // Currency details
var details = map[Currency]ISO4217{
	KZT: {
		Name:           "Tenge",
		AlphabeticCode: "KZT",
		NumericCode:    398,
		MinorUnits:     MinorUnits2,
	},
	RUB: {
		Name:           "Russian Ruble",
		AlphabeticCode: "RUB",
		NumericCode:    643,
		MinorUnits:     MinorUnits2,
	},
	USD: {
		Name:           "US Dollar",
		AlphabeticCode: "USD",
		NumericCode:    840,
		MinorUnits:     MinorUnits2,
	},
	EUR: {
		Name:           "Euro",
		AlphabeticCode: "EUR",
		NumericCode:    978,
		MinorUnits:     MinorUnits2,
	},
	GBP: {
		Name:           "Great Britain Pound",
		AlphabeticCode: "GBP",
		NumericCode:    826,
		MinorUnits:     MinorUnits2,
	},
	CHF: {
		Name:           "Swiss Franc",
		AlphabeticCode: "CHF",
		NumericCode:    756,
		MinorUnits:     MinorUnits2,
	},
	CAD: {
		Name:           "Canadian Dollar",
		AlphabeticCode: "CAD",
		NumericCode:    124,
		MinorUnits:     MinorUnits2,
	},
	JPY: {
		Name:           "Japanese Yen",
		AlphabeticCode: "JPY",
		NumericCode:    392,
		MinorUnits:     MinorUnits2,
	},
	CNY: {
		Name:           "Chinese Yuan",
		AlphabeticCode: "CNY",
		NumericCode:    156,
		MinorUnits:     MinorUnits2,
	},
	AUD: {
		Name:           "Australian dollar",
		AlphabeticCode: "AUD",
		NumericCode:    036,
		MinorUnits:     MinorUnits2,
	},
	AED: {
		Name:           "United Arab Emirates Dirham",
		AlphabeticCode: "AED",
		NumericCode:    784,
		MinorUnits:     MinorUnits2,
	},
	TRY: {
		Name:           "Turkish lira",
		AlphabeticCode: "TRY",
		NumericCode:    949,
		MinorUnits:     MinorUnits2,
	},
	KGS: {
		Name:           "Kyrgyz som",
		AlphabeticCode: "KGS",
		NumericCode:    417,
		MinorUnits:     MinorUnits2,
	},
	ZAR: {
		Name:           "South African rand",
		AlphabeticCode: "ZAR",
		NumericCode:    710,
		MinorUnits:     MinorUnits2,
	},
	NOK: {
		Name:           "Norsk krone",
		AlphabeticCode: "NOK",
		NumericCode:    578,
		MinorUnits:     MinorUnits2,
	},
	SEK: {
		Name:           "Swedish Krona",
		AlphabeticCode: "SEK",
		NumericCode:    752,
		MinorUnits:     MinorUnits2,
	},
}

//nolint:gochecknoglobals // ISO 4217 numeric codes
var detailsIso = map[uint16]Currency{
	398: KZT,
	643: RUB,
	840: USD,
	978: EUR,
	826: GBP,
	756: CHF,
	124: CAD,
	392: JPY,
	156: CNY,
	036: AUD,
	784: AED,
	949: TRY,
	417: KGS,
	710: ZAR,
	578: NOK,
	752: SEK,
}
