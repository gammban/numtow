package ru

import (
	"github.com/gammban/numtow/internal/digit"
	"github.com/gammban/numtow/lang/ru/gender"
)

const (
	zero  = "ноль"
	minus = "минус"
	sep   = " "
)

//nolint
var (
	// единицы
	unitsMale   = [10]string{"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"} // мужской род
	unitsFemale = [10]string{"", "одна", "две", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"} // женский род
	unitsNeuter = [10]string{"", "одно", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"} // средний род
	// от 11 до 19
	teens = [10]string{"", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
	// десятки
	tens = [10]string{"", "десять", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
	// сотни:
	hundreds = [10]string{"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}
	// mega
	megaPlural   = [14]string{"", "тысяч", "миллионов", "миллиардов", "триллионов", "квадриллионов", "квинтиллионов", "секстиллионов", "септиллионов", "октиллионов", "нониллионов", "дециллионов", "ундециллионов", "дуодециллионов"}
	megaSingular = [14]string{"", "тысяча", "миллион", "миллиард", "триллион", "квадриллион", "квинтиллион", "секстиллион", "септиллион", "октиллион", "нониллион", "дециллион", "ундециллион", "дуодециллион"}
	mega234      = [14]string{"", "тысячи", "миллиона", "миллиарда", "триллиона", "квадриллиона", "квинтиллиона", "секстиллиона", "септиллиона", "октиллиона", "нониллиона", "дециллиона", "ундециллиона", "дуодециллиона"}
)

func getMegaByDeclination(d Declination, idx int) string {
	switch d {
	case DeclinationSingular:
		return megaSingular[idx]
	case Declination234:
		return mega234[idx]
	case DeclinationPlural:
		return megaPlural[idx]
	default:
		return megaPlural[idx]
	}
}

func getUnitsByGender(g gender.Gender, unit digit.Digit) (result string, err error) {
	switch g {
	case gender.Male:
		return unitsMale[unit], nil
	case gender.Female:
		return unitsFemale[unit], nil
	case gender.Neuter:
		return unitsNeuter[unit], nil
	case gender.Unknown:
		return "", gender.ErrBadGender
	default:
		return "", gender.ErrBadGender
	}
}
