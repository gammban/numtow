package ru

import (
	"github.com/gammban/numtow/lang/ru/gender"
)

// Options for russian language.
//
// Опции перевода чисел в слова для русского языка.
type Options struct {
	// Decimal separator for parsing.
	// Десятичный разделитель целой и дробной части. По умолчанию '.'. Используется при парсинге числа из строки.
	ParseSeparator rune
	// Размер дробной части. По умолчанию 0.
	ParseFracLen uint
	// Род. По умолчанию женский.
	FmtGender gender.Gender
	// Не учитывать дробную часть числа. По умолчанию учитывается.
	FmtFracIgnore bool
	// Использовать числа вместо слов в дробной части. По умолчанию используются слова.
	FmtFracUseDigits bool
}

//nolint:gochecknoglobals
var (
	FormatDefault = []OptFunc{
		WithFmtGender(gender.Female),
		WithParseSep('.'),
		WithParseFracLen(0),
		WithFmtFracIgnore(false),
		WithFmtFracUseDigits(false),
	}
)

type OptFunc func(o *Options)

// WithFmtGender set gender. Default value is gender.Female.
func WithFmtGender(g gender.Gender) OptFunc {
	return func(o *Options) {
		o.FmtGender = g
	}
}

// WithParseSep sets decimal separator. Default value is point ('.').
func WithParseSep(separator rune) OptFunc {
	return func(o *Options) {
		o.ParseSeparator = separator
	}
}

// WithParseFracLen sets fractional part parsing length. Default value is 0.
func WithParseFracLen(fracLen uint) OptFunc {
	return func(o *Options) {
		o.ParseFracLen = fracLen
	}
}

// WithFmtFracIgnore sets fractional part hiding flag. Default value is false.
func WithFmtFracIgnore(ignore bool) OptFunc {
	return func(o *Options) {
		o.FmtFracIgnore = ignore
	}
}

// WithFmtFracUseDigits sets using digits in fractional part. Default value is false.
func WithFmtFracUseDigits(useDigits bool) OptFunc {
	return func(o *Options) {
		o.FmtFracUseDigits = useDigits
	}
}

func prepareOptions(o ...OptFunc) *Options {
	e := &Options{
		FmtGender:      gender.Female,
		ParseSeparator: '.',
	}

	for _, opt := range o {
		opt(e)
	}

	return e
}

func ParseOpts(options ...interface{}) []OptFunc {
	if len(options) == 0 {
		return nil
	}

	o := make([]OptFunc, 0, len(options))

	for _, v := range options {
		opt, ok := v.(OptFunc)
		if ok {
			o = append(o, opt)
			continue
		}

		opts, ok := v.([]OptFunc)
		if ok {
			o = opts
			break
		}
	}

	return o
}
