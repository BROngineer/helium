package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fint16 = flag[int16]

type Int16 struct {
	*fint16
}

type int16Parser struct {
	*embeddedParser
}

func defaultInt16Parser() *int16Parser {
	return &int16Parser{&embeddedParser{}}
}

func (p *int16Parser) ParseCmd(input string) (any, error) {
	if p.IsVisited() {
		return nil, errors.ErrFlagVisited
	}
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	v, err := strconv.ParseInt(input, 10, 16)
	if err != nil {
		return nil, err
	}
	parsed := int16(v)
	return &parsed, nil
}

func (p *int16Parser) ParseEnv(value string) (any, error) {
	var (
		v   int64
		err error
	)
	if v, err = strconv.ParseInt(value, 10, 16); err != nil {
		return nil, err
	}
	parsed := int16(v)
	return &parsed, nil
}

func NewInt16(name string, opts ...Option) *Int16 {
	f := newFlag[int16](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultInt16Parser())
	}
	return &Int16{f}
}
