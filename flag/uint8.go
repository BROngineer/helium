package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fuint8 = flag[uint8]

type Uint8 struct {
	*fuint8
}

type uint8Parser struct {
	*embeddedParser
}

func defaultUint8Parser() *uint8Parser {
	return &uint8Parser{&embeddedParser{}}
}

func (p *uint8Parser) ParseCmd(input string) (any, error) {
	if p.IsVisited() {
		return nil, errors.ErrFlagVisited
	}
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	v, err := strconv.ParseUint(input, 10, 8)
	if err != nil {
		return nil, err
	}
	parsed := uint8(v)
	return &parsed, nil
}

func (p *uint8Parser) ParseEnv(input string) (any, error) {
	var (
		v   uint64
		err error
	)
	if v, err = strconv.ParseUint(input, 10, 8); err != nil {
		return nil, err
	}
	parsed := uint8(v)
	return &parsed, nil
}

func NewUint8(name string, opts ...Option) *Uint8 {
	f := newFlag[uint8](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultUint8Parser())
	}
	return &Uint8{f}
}
