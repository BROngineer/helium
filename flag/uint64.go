package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fuint64 = flag[uint64]

type Uint64 struct {
	*fuint64
}

type uint64Parser struct {
	*embeddedParser
}

func defaultUint64Parser() *uint64Parser {
	return &uint64Parser{&embeddedParser{}}
}

func (p *uint64Parser) ParseCmd(input string) (any, error) {
	if p.IsVisited() {
		return nil, errors.ErrFlagVisited
	}
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	v, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return nil, err
	}
	parsed := v
	return &parsed, nil
}

func (p *uint64Parser) ParseEnv(input string) (any, error) {
	var (
		v   uint64
		err error
	)
	if v, err = strconv.ParseUint(input, 10, 64); err != nil {
		return nil, err
	}
	parsed := v
	return &parsed, nil
}

func NewUint64(name string, opts ...Option) *Uint64 {
	f := newFlag[uint64](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultUint64Parser())
	}
	return &Uint64{f}
}
