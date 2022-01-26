package en

const (
	defaultParseDecimalSeparator = '.'
	defaultFormatMinusSign       = minus
	defaultFormatGroupSeparator  = ","
	defaultFormatAndWord         = and
)

//nolint:gochecknoglobals
var (
	FormatDefault = []OptFunc{
		WithParseSep(defaultParseDecimalSeparator),
		WithParseFracLen(0),
		WithFmtGroupSep(defaultFormatGroupSeparator),
		WithFmtAndSep(defaultFormatAndWord),
		WithFmtMinus(),
		WithFmtFracIgnore(false),
		WithFmtFracUseDigits(false),
	}

	FormatWithoutAnd = []OptFunc{
		WithParseSep(defaultParseDecimalSeparator),
		WithParseFracLen(0),
		WithFmtGroupSep(""),
		WithFmtAndSep(""),
		WithFmtMinus(),
		WithFmtFracIgnore(false),
		WithFmtFracUseDigits(false),
	}
)

type Options struct {
	Parse  ParseOption
	Format FormatOption
}

type ParseOption struct {
	DecimalSeparator rune
	FracLen          uint
}

type FormatOption struct {
	GroupSeparator string // number group separator. Default value is comma ",".
	AndWord        string // number group separator. Default value is "and".
	MinusSignWord  string
	FracIgnore     bool
	FracUseDigits  bool
}

type OptFunc func(o *Options)

// WithFmtGroupSep sets FormatOption number group separator. Default value is comma (',').
func WithFmtGroupSep(groupSeparator string) OptFunc {
	return func(o *Options) {
		o.Format.GroupSeparator = groupSeparator
	}
}

// WithFmtAndSep sets FormatOption number group separator. Default value is "and".
func WithFmtAndSep(andSeparator string) OptFunc {
	return func(o *Options) {
		o.Format.AndWord = andSeparator
	}
}

// WithParseSep sets ParseOption time decimal separator. Default value is '.'.
func WithParseSep(parseDecimalSeparator rune) OptFunc {
	return func(o *Options) {
		o.Parse.DecimalSeparator = parseDecimalSeparator
	}
}

// WithParseFracLen sets ParseOption time length of fractional part to truncate. Default value is 0.
func WithParseFracLen(fracLen uint) OptFunc {
	return func(o *Options) {
		o.Parse.FracLen = fracLen
	}
}

func WithFmtFracIgnore(ignore bool) OptFunc {
	return func(o *Options) {
		o.Format.FracIgnore = ignore
	}
}

func WithFmtFracUseDigits(showNumber bool) OptFunc {
	return func(o *Options) {
		o.Format.FracUseDigits = showNumber
	}
}

func WithFmtMinus() OptFunc {
	return func(o *Options) {
		o.Format.MinusSignWord = minus
	}
}

func WithFmtNegative() OptFunc {
	return func(o *Options) {
		o.Format.MinusSignWord = negative
	}
}

func prepareOptions(o ...OptFunc) *Options {
	e := &Options{
		Format: FormatOption{
			GroupSeparator: defaultFormatGroupSeparator,
			MinusSignWord:  defaultFormatMinusSign,
			AndWord:        defaultFormatAndWord,
			FracIgnore:     false,
			FracUseDigits:  false,
		},
		Parse: ParseOption{
			FracLen:          0,
			DecimalSeparator: defaultParseDecimalSeparator,
		},
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
