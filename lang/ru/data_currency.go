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
			NameGender: gender.Female,
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
