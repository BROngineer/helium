package env

type config interface {
	setPrefix(string)
	setCapitalize()
}

type Option interface {
	apply(config)
}

type prefix string

func (o prefix) apply(c config) {
	c.setPrefix(string(o))
}

func Prefix(p string) Option {
	return prefix(p)
}

type capitalized struct{}

func (o capitalized) apply(c config) {
	c.setCapitalize()
}

func Capitalized() Option {
	return capitalized{}
}
