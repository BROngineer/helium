package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fbool = flag[bool]

type Bool struct {
	*fbool
}

type boolParser struct {
	*embeddedParser
}

func defaultBoolParser() *boolParser {
	return &boolParser{&embeddedParser{}}
}

func (p *boolParser) ParseCmd(input string) (any, error) {
	var (
		parsed bool
		err    error
	)
	if p.IsVisited() {
		return nil, errors.ErrFlagVisited
	}
	switch input {
	case "":
		parsed = true
	default:
		parsed, err = strconv.ParseBool(input)
		if err != nil {
			return nil, err
		}
	}
	return &parsed, nil
}

func (p *boolParser) ParseEnv(input string) (any, error) {
	var (
		parsed bool
		err    error
	)
	if parsed, err = strconv.ParseBool(input); err != nil {
		return nil, err
	}
	return &parsed, nil
}

func NewBool(name string, opts ...Option) *Bool {
	f := newFlag[bool](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultBoolParser())
	}
	return &Bool{f}
}
