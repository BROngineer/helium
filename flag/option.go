package flag

type flagPropertySetter interface {
	setDescription(string)
	setShorthand(string)
	setShared()
	setDefaultValue(any)
	setSeparator(string)
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

func applyForFlag(f flagPropertySetter, opts ...Option) {
	for _, opt := range opts {
		opt.apply(f)
	}
}
