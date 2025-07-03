package ru

import "github.com/gammban/numtow/lang/ru/gender"

const (
	intPartMale   = "целых"
	intPartFemale = "целая"
	intPartNeuter = "целое"
	fracPartLen   = 15
)

//nolint:gochecknoglobals // Fractional part names for russian language.
var (
	fracPartMale = [fracPartLen]string{
		"", "десятых", "сотых",
		"тысячных", "десятитысячных", "стотысячных",
		"миллионных", "десятимиллионных", "стомиллионных",
		"миллиардных", "десятимиллиардных", "стомиллиардных",
		"триллионных", "десятитриллионных", "стотриллионных",
	}
	fracPartFemaleSingular = [fracPartLen]string{
		"", "десятая", "сотая",
		"тысячная", "десятитысячная", "стотысячная",
		"миллионная", "десятимиллионная", "стомиллионная",
		"миллиардная", "десятимиллиардная", "стомиллиардная",
		"триллионная", "десятитриллионная", "стотриллионная",
	}
	fracPartFemale234 = [fracPartLen]string{
		"", "десятых", "сотых",
		"тысячных", "десятитысячных", "стотысячных",
		"миллионных", "десятимиллионных", "стомиллионных",
		"миллиардных", "десятимиллиардных", "стомиллиардных",
		"триллионных", "десятитриллионных", "стотриллионных",
	}
	fracPartFemalePlural = [fracPartLen]string{
		"", "десятых", "сотых",
		"тысячных", "десятитысячных", "стотысячных",
		"миллионных", "десятимиллионных", "стомиллионных",
		"миллиардных", "десятимиллиардных", "стомиллиардных",
		"триллионных", "десятитриллионных", "стотриллионных",
	}
	fracPartNeuterSingular = [fracPartLen]string{
		"", "десятое", "сотое",
		"тысячное", "десятитысячное", "стотысячное",
		"миллионное", "десятимиллионное", "стомиллионное",
		"миллиардное", "десятимиллиардное", "стомиллиардное",
		"триллионное", "десятитриллионное", "стотриллионное",
	}
	fracPartNeuter = [fracPartLen]string{
		"", "десятых", "сотых",
		"тысячных", "десятитысячных", "стотысячных",
		"миллионных", "десятимиллионных", "стомиллионных",
		"миллиардных", "десятимиллиардных", "стомиллиардных",
		"триллионных", "десятитриллионных", "стотриллионных",
	}
)

func getFracPart(d Declination, g gender.Gender, idx int) string {
	if idx > fracPartLen {
		return ""
	}

	switch g {
	case gender.Male:
		switch d {
		case DeclinationSingular, Declination234, DeclinationPlural:
			return fracPartMale[idx]
		default:
			return ""
		}
	case gender.Female:
		switch d {
		case DeclinationSingular:
			return fracPartFemaleSingular[idx]
		case Declination234:
			return fracPartFemale234[idx]
		case DeclinationPlural:
			return fracPartFemalePlural[idx]
		default:
			return ""
		}
	case gender.Neuter:
		switch d {
		case DeclinationSingular:
			return fracPartNeuterSingular[idx]
		case Declination234, DeclinationPlural:
			return fracPartNeuter[idx]
		default:
			return ""
		}
	case gender.Unknown:
		return ""
	default:
		return ""
	}
}

func getIntPart(d Declination, g gender.Gender) string {
	switch g {
	case gender.Female:
		switch d {
		case DeclinationSingular:
			return intPartFemale
		case Declination234, DeclinationPlural:
			return intPartMale
		default:
			return ""
		}
	case gender.Male:
		switch d {
		case DeclinationSingular, Declination234, DeclinationPlural:
			return intPartMale
		default:
			return ""
		}
	case gender.Neuter:
		switch d {
		case DeclinationSingular:
			return intPartNeuter
		case Declination234, DeclinationPlural:
			return intPartMale
		default:
			return ""
		}
	case gender.Unknown:
		return ""
	default:
		return ""
	}
}
