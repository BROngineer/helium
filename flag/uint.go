package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fuint = flag[uint]

type UintFlag struct {
	*fuint
}

type uintParser struct {
	*embeddedParser
}

func defaultUintParser() *uintParser {
	return &uintParser{&embeddedParser{}}
}

func (p *uintParser) ParseCmd(input string) (any, error) {
	if p.IsSetFromCmd() {
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
	parsed := uint(v)
	return &parsed, nil
}

func (p *uintParser) ParseEnv(input string) (any, error) {
	var (
		v   uint64
		err error
	)
	if v, err = strconv.ParseUint(input, 10, 32); err != nil {
		return nil, err
	}
	parsed := uint(v)
	return &parsed, nil
}

func Uint(name string, opts ...Option) *UintFlag {
	f := newFlag[uint](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultUintParser())
	}
	return &UintFlag{f}
}
