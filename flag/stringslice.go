package flag

import (
	"strings"

	"github.com/brongineer/helium/errors"
)

type stringSlice = flag[[]string]

type StringSlice struct {
	*stringSlice
}

type stringSliceParser struct {
	*embeddedParser
}

func defaultStringSliceParser() *stringSliceParser {
	return &stringSliceParser{&embeddedParser{}}
}

func (p *stringSliceParser) ParseCmd(input string) (any, error) {
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	parsed := strings.Split(input, p.Separator())
	if p.IsVisited() {
		stored := DerefOrDie[[]string](p.CurrentValue())
		parsed = append(stored, parsed...)
	}
	return &parsed, nil
}

func (p *stringSliceParser) ParseEnv(input string) (any, error) {
	parsed := strings.Split(input, p.Separator())
	return &parsed, nil
}

func NewStringSlice(name string, opts ...Option) *StringSlice {
	f := newFlag[[]string](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultStringSliceParser())
	}
	return &StringSlice{f}
}
