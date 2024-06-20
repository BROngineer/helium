package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fuint16 = flag[uint16]

type Uint16 struct {
	*fuint16
}

type uint16Parser struct {
	*embeddedParser
}

func defaultUint16Parser() *uint16Parser {
	return &uint16Parser{&embeddedParser{}}
}

func (p *uint16Parser) ParseCmd(input string) (any, error) {
	if p.IsVisited() {
		return nil, errors.ErrFlagVisited
	}
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	v, err := strconv.ParseUint(input, 10, 16)
	if err != nil {
		return nil, err
	}
	parsed := uint16(v)
	return &parsed, nil
}

func (p *uint16Parser) ParseEnv(input string) (any, error) {
	var (
		v   uint64
		err error
	)
	if v, err = strconv.ParseUint(input, 10, 16); err != nil {
		return nil, err
	}
	parsed := uint16(v)
	return &parsed, nil
}

func NewUint16(name string, opts ...Option) *Uint16 {
	f := newFlag[uint16](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultUint16Parser())
	}
	return &Uint16{f}
}
