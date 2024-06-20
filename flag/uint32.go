package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fuint32 = flag[uint32]

type Uint32 struct {
	*fuint32
}

type uint32Parser struct {
	*embeddedParser
}

func defaultUint32Parser() *uint32Parser {
	return &uint32Parser{&embeddedParser{}}
}

func (p *uint32Parser) ParseCmd(input string) (any, error) {
	if p.IsVisited() {
		return nil, errors.ErrFlagVisited
	}
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	v, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return nil, err
	}
	parsed := uint32(v)
	return &parsed, nil
}

func (p *uint32Parser) ParseEnv(input string) (any, error) {
	var (
		v   uint64
		err error
	)
	if v, err = strconv.ParseUint(input, 10, 32); err != nil {
		return nil, err
	}
	parsed := uint32(v)
	return &parsed, nil
}

func NewUint32(name string, opts ...Option) *Uint32 {
	f := newFlag[uint32](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultUint32Parser())
	}
	return &Uint32{f}
}
