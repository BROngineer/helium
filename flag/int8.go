package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fint8 = flag[int8]

type Int8 struct {
	*fint8
}

type int8Parser struct {
	*embeddedParser
}

func defaultInt8Parser() *int8Parser {
	return &int8Parser{&embeddedParser{}}
}

func (p *int8Parser) ParseCmd(input string) (any, error) {
	if p.IsVisited() {
		return nil, errors.ErrFlagVisited
	}
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	v, err := strconv.ParseInt(input, 10, 8)
	if err != nil {
		return nil, err
	}
	parsed := int8(v)
	return &parsed, nil
}

func (p *int8Parser) ParseEnv(value string) (any, error) {
	var (
		v   int64
		err error
	)
	if v, err = strconv.ParseInt(value, 10, 8); err != nil {
		return nil, err
	}
	parsed := int8(v)
	return &parsed, nil
}

func NewInt8(name string, opts ...Option) *Int8 {
	f := newFlag[int8](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultInt8Parser())
	}
	return &Int8{f}
}
