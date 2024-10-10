package flag

import (
	"github.com/brongineer/helium/errors"
)

type fstring = flag[string]

type StringFlag struct {
	*fstring
}

type stringParser struct {
	*embeddedParser
}

func defaultStringParser() *stringParser {
	return &stringParser{&embeddedParser{}}
}

func (p *stringParser) ParseCmd(input string) (any, error) {
	if p.IsSetFromCmd() {
		return nil, errors.ErrFlagVisited
	}
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	parsed := input
	return &parsed, nil
}

func (p *stringParser) ParseEnv(input string) (any, error) {
	parsed := input
	return &parsed, nil
}

func String(name string, opts ...Option) *StringFlag {
	f := newFlag[string](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultStringParser())
	}
	return &StringFlag{f}
}
