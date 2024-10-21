package ru

import (
	"fmt"
	"testing"

	"github.com/gammban/numtow"
	"github.com/gammban/numtow/lang"
	"github.com/gammban/numtow/lang/ru"
	"github.com/gammban/numtow/lang/ru/gender"
)

func TestRu(t *testing.T) {
	// convert number to russian using numtow package
	fmt.Println(numtow.MustString("1", lang.RU))                                 // одна
	fmt.Println(numtow.MustString("2", lang.RU))                                 // две
	fmt.Println(numtow.MustFloat64(2, lang.RU, ru.WithFmtGender(gender.Male)))   // два
	fmt.Println(numtow.MustInt64(731, lang.RU, ru.WithFmtGender(gender.Neuter))) // семьсот тридцать одно
	fmt.Println(numtow.MustString("15345712", lang.RU))                          // пятнадцать миллионов триста сорок пять тысяч семьсот двенадцать
	fmt.Println(numtow.MustString("123.457", lang.RU, ru.FormatDefault))         // сто двадцать три целых четыреста пятьдесят семь тысячных

	// convert number to russian words using ru package
	fmt.Println(ru.MustString("2"))                                             // две
	fmt.Println(ru.MustFloat64(2, ru.WithFmtGender(gender.Male)))               // два
	fmt.Println(ru.MustString("123.457", ru.FormatDefault...))                  // сто двадцать три целых четыреста пятьдесят семь тысячных
	fmt.Println(ru.MustString("123.457", ru.WithFmtFracIgnore(true)))           // сто двадцать три
	fmt.Println(ru.MustString("123.457", ru.WithFmtFracUseDigits(true)))        // сто двадцать три целых 457 тысячных
	fmt.Println(ru.MustString("123.457", ru.WithParseFracLen(2)))               // сто двадцать три целых сорок пять сотых
	fmt.Println(ru.MustString("123,457", ru.WithParseSep(',')))                 // сто двадцать три целых четыреста пятьдесят семь тысячных
	fmt.Println(ru.MustString("0.01"))                                          // ноль целых одна сотая
	fmt.Println(ru.MustString(".01"))                                           // ноль целых одна сотая
	fmt.Println(ru.MustString("10"))                                            // ноль целых одна сотая
	fmt.Println(ru.MustString("1", ru.WithFmtGender(gender.Female)), "рубашка") // одна
	fmt.Println(ru.MustString("1", ru.WithFmtGender(gender.Male)), "карандаш")  // один карандаш
	fmt.Println(ru.MustString("1", ru.WithFmtGender(gender.Neuter)), "окно")    // одно окно
	fmt.Println(ru.MustString("12"))                                            // двенадцать
	fmt.Println(ru.MustString("-57"))                                           // минус пятьдесят семь
	fmt.Println(ru.MustString("148.64"))                                        // сто сорок восемь целых шестьдесят четыре сотых
	fmt.Println(ru.MustString("999999999"))                                     // девятьсот девяносто девять миллионов девятьсот девяносто девять тысяч девятьсот девяносто девять

	_, err := ru.String("bad")
	if err != nil {
		fmt.Println(err) // parse number error
	}

	result, err := ru.String("5")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result) // пять

	result, err = ru.Float64(1.5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result) // одна целая пять десятых

	result, err = ru.Int64(55)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result) // пятьдесят пять

	var MyFormat = []ru.OptFunc{
		ru.WithParseSep(','),
		ru.WithFmtGender(gender.Male),
		ru.WithFmtFracUseDigits(true),
	}

	fmt.Println(ru.MustString("12,57", MyFormat...)) // двенадцать целых 57 сотых
}
