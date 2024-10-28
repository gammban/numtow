package ru

import (
	"fmt"

	"github.com/gammban/numtow/curtow/cur"
	"github.com/gammban/numtow/lang/ru/gender"
)

var (
	//nolint
	currencyOpts = map[cur.Currency]currencyInfo{
		cur.KZT: {
			Name: map[Declination]string{
				DeclinationPlural:   "тенге",
				DeclinationSingular: "тенге",
				Declination234:      "тенге",
			},
			NameGender: gender.Male,
			UnitName: map[Declination]string{
				DeclinationPlural:   "тиын",
				DeclinationSingular: "тиын",
				Declination234:      "тиын",
			},
			UnitGender: gender.Female,
		},
		cur.RUB: {
			Name: map[Declination]string{
				DeclinationPlural:   "рублей",
				DeclinationSingular: "рубль",
				Declination234:      "рубля",
			},
			NameGender: gender.Male,
			UnitName: map[Declination]string{
				DeclinationPlural:   "копеек",
				DeclinationSingular: "копейка",
				Declination234:      "копейки",
			},
			UnitGender: gender.Female,
		},
		cur.USD: {
			Name: map[Declination]string{
				DeclinationPlural:   "долларов США",
				DeclinationSingular: "доллар США",
				Declination234:      "доллара США",
			},
			NameGender: gender.Male,
			UnitName: map[Declination]string{
				DeclinationPlural:   "центов",
				DeclinationSingular: "цент",
				Declination234:      "цента",
			},
			UnitGender: gender.Male,
		},
		cur.EUR: {
			Name: map[Declination]string{
				DeclinationPlural:   "евро",
				DeclinationSingular: "евро",
				Declination234:      "евро",
			},
			NameGender: gender.Male,
			UnitName: map[Declination]string{
				DeclinationPlural:   "евроцентов",
				DeclinationSingular: "евроцент",
				Declination234:      "евроцента",
			},
			UnitGender: gender.Male,
		},
		cur.GBP: {
			Name: map[Declination]string{
				DeclinationPlural:   "фунтов",
				DeclinationSingular: "фунт",
				Declination234:      "фунта",
			},
			NameGender: gender.Male,
			UnitName: map[Declination]string{
				DeclinationPlural:   "пенсов",
				DeclinationSingular: "пенс",
				Declination234:      "пенса",
			},
			UnitGender: gender.Male,
		},
		cur.CHF: {
			Name: map[Declination]string{
				DeclinationPlural:   "франков",
				DeclinationSingular: "франк",
				Declination234:      "франка",
			},
			NameGender: gender.Male,
			UnitName: map[Declination]string{
				DeclinationPlural:   "раппенов",
				DeclinationSingular: "раппен",
				Declination234:      "раппена",
			},
			UnitGender: gender.Male,
		},
		cur.CAD: {
			Name: map[Declination]string{
				DeclinationPlural:   "канадский долларов",
				DeclinationSingular: "канадский доллар",
				Declination234:      "канадских доллара",
			},
			NameGender: gender.Male,
			UnitName: map[Declination]string{
				DeclinationPlural:   "центов",
				DeclinationSingular: "цент",
				Declination234:      "цента",
			},
			UnitGender: gender.Male,
		},
		cur.JPY: {
			Name: map[Declination]string{
				DeclinationPlural:   "иен",
				DeclinationSingular: "иена",
				Declination234:      "иены",
			},
			NameGender: gender.Female,
			UnitName: map[Declination]string{
				DeclinationPlural:   "сены",
				DeclinationSingular: "сен",
				Declination234:      "сены",
			},
			UnitGender: gender.Male,
		},
		cur.CNY: {Name: map[Declination]string{
			DeclinationPlural:   "юаней",
			DeclinationSingular: "юань",
			Declination234:      "юаня",
		},
			NameGender: gender.Male,
			UnitName: map[Declination]string{
				DeclinationPlural:   "цзяо",
				DeclinationSingular: "цзяо",
				Declination234:      "цзяо",
			},
			UnitGender: gender.Male,
		},
		cur.AUD: {
			Name: map[Declination]string{
				DeclinationPlural:   "австралийских долларов",
				DeclinationSingular: "австралийский доллар",
				Declination234:      "австралийских доллара",
			},
			NameGender: gender.Male,
			UnitName: map[Declination]string{
				DeclinationPlural:   "центов",
				DeclinationSingular: "цент",
				Declination234:      "цента",
			},
			UnitGender: gender.Male,
		},
		cur.AED: {
			Name: map[Declination]string{
				DeclinationPlural:   "дирхамов",
				DeclinationSingular: "дирхам",
				Declination234:      "дирхама",
			},
			NameGender: gender.Male,
			UnitName: map[Declination]string{
				DeclinationPlural:   "филсов",
				DeclinationSingular: "филс",
				Declination234:      "филса",
			},
			UnitGender: gender.Male,
		},
		cur.TRY: {
			Name: map[Declination]string{
				DeclinationPlural:   "лир",
				DeclinationSingular: "лира",
				Declination234:      "лиры",
			},
			NameGender: gender.Female,
			UnitName: map[Declination]string{
				DeclinationPlural:   "курушей",
				DeclinationSingular: "куруш",
				Declination234:      "куруша",
			},
			UnitGender: gender.Male,
		},
		cur.KGS: {
			Name: map[Declination]string{
				DeclinationPlural:   "сомов",
				DeclinationSingular: "сом",
				Declination234:      "сома",
			},
			NameGender: gender.Male,
			UnitName: map[Declination]string{
				DeclinationPlural:   "тиынов",
				DeclinationSingular: "тиын",
				Declination234:      "тиына",
			},
			UnitGender: gender.Male,
		},
		cur.ZAR: {
			Name: map[Declination]string{
				DeclinationPlural:   "рэндов",
				DeclinationSingular: "рэнд",
				Declination234:      "рэнда",
			},
			NameGender: gender.Male,
			UnitName: map[Declination]string{
				DeclinationPlural:   "центов",
				DeclinationSingular: "цент",
				Declination234:      "цента",
			},
			UnitGender: gender.Male,
		},
		cur.NOK: {
			Name: map[Declination]string{
				DeclinationPlural:   "кронов",
				DeclinationSingular: "крона",
				Declination234:      "кроны",
			},
			NameGender: gender.Female,
			UnitName: map[Declination]string{
				DeclinationPlural:   "эре",
				DeclinationSingular: "эре",
				Declination234:      "эре",
			},
			UnitGender: gender.Male,
		},
		cur.SEK: {
			Name: map[Declination]string{
				DeclinationPlural:   "кронов",
				DeclinationSingular: "крона",
				Declination234:      "кроны",
			},
			NameGender: gender.Female,
			UnitName: map[Declination]string{
				DeclinationPlural:   "эре",
				DeclinationSingular: "эре",
				Declination234:      "эре",
			},
			UnitGender: gender.Male,
		},
	}
)

type currencyInfo struct {
	Name       map[Declination]string
	NameGender gender.Gender
	UnitName   map[Declination]string
	UnitGender gender.Gender
}

func (c currencyInfo) GetCurrencyName(d Declination) string {
	res, ok := c.Name[d]
	if ok {
		return res
	}

	return ""
}

func (c currencyInfo) GetCurrencyGender() gender.Gender {
	return c.NameGender
}

func (c currencyInfo) GetCurrencyUnitName(d Declination) string {
	res, ok := c.UnitName[d]
	if ok {
		return res
	}

	return ""
}

func getCurrencyInfo(c cur.Currency) (info currencyInfo, err error) {
	info, ok := currencyOpts[c]
	if ok {
		return info, nil
	}

	return info, fmt.Errorf("%w: info not found", cur.ErrBadCurrency)
}

func (c currencyInfo) GetCurrencyUnitGender() gender.Gender {
	return c.UnitGender
}
