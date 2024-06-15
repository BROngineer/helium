package flag

type flagPropertySetter interface {
	setDescription(string)
	setShorthand(string)
	setShared()
	setDefaultValue(any)
	setSeparator(string)
	setCommandLineParser(func(string) (any, error))
	setEnvVariableParser(func(string) (any, error))
}

type Option interface {
	apply(flagPropertySetter)
}

type description struct {
	value string
}

func (d description) apply(f flagPropertySetter) {
	f.setDescription(d.value)
}

func Description(value string) Option {
	return description{value}
}

type shorthand struct {
	value string
}

func (s shorthand) apply(f flagPropertySetter) {
	f.setShorthand(s.value)
}

func Shorthand(value string) Option {
	return shorthand{value}
}

type shared struct{}

func (s shared) apply(f flagPropertySetter) {
	f.setShared()
}

func Shared() Option {
	return shared{}
}

type defaultValue struct {
	value any
}

func (d defaultValue) apply(f flagPropertySetter) {
	f.setDefaultValue(d.value)
}

func DefaultValue(value any) Option {
	return defaultValue{value}
}

type separator struct {
	value string
}

func (s separator) apply(f flagPropertySetter) {
	f.setSeparator(s.value)
}

func Separator(value string) Option {
	return separator{value}
}

type commandLineParser func(string) (any, error)

func (o commandLineParser) apply(f flagPropertySetter) {
	f.setCommandLineParser(o)
}

func CommandLineParser(parser func(string) (any, error)) Option {
	return commandLineParser(parser)
}

type envVariableParser func(string) (any, error)

func (o envVariableParser) apply(f flagPropertySetter) {
	f.setEnvVariableParser(o)
}

func EnvVariableParser(parser func(string) (any, error)) Option {
	return envVariableParser(parser)
}

func applyForFlag(f flagPropertySetter, opts ...Option) {
	for _, opt := range opts {
		opt.apply(f)
	}
}
