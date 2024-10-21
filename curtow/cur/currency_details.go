package cur

type ISO4217 struct {
	Name           string
	AlphabeticCode string
	NumericCode    uint16
	MinorUnits     MinorUnits
}

//nolint:gochecknoglobals
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
}
