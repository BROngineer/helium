package options

type flag interface {
	SetDescription(string)
	SetShorthand(string)
	SetRequired()
	SetShared()
	SetDefaultValue(any)
	SetSeparator(string)
}

type FlagOption interface {
	apply(flag)
}

type description struct {
	value string
}

func (d description) apply(f flag) {
	f.SetDescription(d.value)
}

func Description(value string) FlagOption {
	return description{value}
}

type shorthand struct {
	value string
}

func (s shorthand) apply(f flag) {
	f.SetShorthand(s.value)
}

func Shorthand(value string) FlagOption {
	return shorthand{value}
}

type required struct{}

func (r required) apply(f flag) {
	f.SetRequired()
}

func Required() FlagOption {
	return required{}
}

type shared struct{}

func (s shared) apply(f flag) {
	f.SetShared()
}

func Shared() FlagOption {
	return shared{}
}

type defaultValue struct {
	value any
}

func (d defaultValue) apply(f flag) {
	f.SetDefaultValue(d.value)
}

func DefaultValue(value any) FlagOption {
	return defaultValue{value}
}

type separator struct {
	value string
}

func (s separator) apply(f flag) {
	f.SetSeparator(s.value)
}

func Separator(value string) FlagOption {
	return separator{value}
}

func ApplyForFlag(f flag, opts ...FlagOption) {
	for _, opt := range opts {
		opt.apply(f)
	}
}
