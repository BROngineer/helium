package env

type config interface {
	setPrefix(string)
	setCapitalize()
	setReplacement(string, string)
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

type varNameReplacement struct {
	old string
	new string
}

func (o varNameReplacement) apply(c config) {
	c.setReplacement(o.old, o.new)
}

func VarNameReplace(oldChar, newChar string) Option {
	return varNameReplacement{oldChar, newChar}
}
