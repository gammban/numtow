package kz

const (
	defaultSeparator = '.'
)

type Options struct {
	ParseFracLen     uint
	ParseSeparator   rune
	FmtFracIgnore    bool
	FmtFracUseDigits bool
}

//nolint:gochecknoglobals // Default options for kazakh language.
var (
	FormatDefault = []OptFunc{
		WithParseSep(defaultSeparator),
		WithParseFracLen(0),
		WithFmtFracIgnore(false),
		WithFmtFracUseDigits(false),
	}
)

type OptFunc func(o *Options)

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
func WithFmtFracUseDigits(showNumber bool) OptFunc {
	return func(o *Options) {
		o.FmtFracUseDigits = showNumber
	}
}

func prepareOptions(o ...OptFunc) *Options {
	e := &Options{
		FmtFracIgnore:    false,
		FmtFracUseDigits: false,
		ParseFracLen:     0,
		ParseSeparator:   defaultSeparator,
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
