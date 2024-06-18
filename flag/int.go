package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fint = flag[int]

type Int struct {
	*fint
}

type intParser struct {
	*embeddedParser
}

func defaultIntParser() *intParser {
	return &intParser{&embeddedParser{}}
}

func (p *intParser) ParseCmd(input string) (any, error) {
	if p.IsVisited() {
		return nil, errors.ErrFlagVisited
	}
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	parsed, err := strconv.Atoi(input)
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

func (p *intParser) ParseEnv(input string) (any, error) {
	var (
		parsed int
		err    error
	)
	if parsed, err = strconv.Atoi(input); err != nil {
		return nil, err
	}
	return &parsed, nil
}

func NewInt(name string, opts ...Option) *Int {
	f := newFlag[int](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultIntParser())
	}
	return &Int{f}
}
