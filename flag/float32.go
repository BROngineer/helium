package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type ffloat32 = flag[float32]

type Float32Flag struct {
	*ffloat32
}

type float32Parser struct {
	*embeddedParser
}

func defaultFloat32Parser() *float32Parser {
	return &float32Parser{&embeddedParser{}}
}

func (p *float32Parser) ParseCmd(input string) (any, error) {
	if p.IsSetFromCmd() {
		return nil, errors.ErrFlagVisited
	}
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	v, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return nil, err
	}
	parsed := float32(v)
	return &parsed, nil
}

func (p *float32Parser) ParseEnv(input string) (any, error) {
	var (
		v   float64
		err error
	)
	if v, err = strconv.ParseFloat(input, 32); err != nil {
		return nil, err
	}
	parsed := float32(v)
	return &parsed, nil
}

func Float32(name string, opts ...Option) *Float32Flag {
	f := newFlag[float32](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultFloat32Parser())
	}
	return &Float32Flag{f}
}
