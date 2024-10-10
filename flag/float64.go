package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type ffloat64 = flag[float64]

type Float64Flag struct {
	*ffloat64
}

type float64Parser struct {
	*embeddedParser
}

func defaultFloat64Parser() *float64Parser {
	return &float64Parser{&embeddedParser{}}
}

func (p *float64Parser) ParseCmd(input string) (any, error) {
	if p.IsSetFromCmd() {
		return nil, errors.ErrFlagVisited
	}
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	parsed, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

func (p *float64Parser) ParseEnv(input string) (any, error) {
	var (
		parsed float64
		err    error
	)
	if parsed, err = strconv.ParseFloat(input, 64); err != nil {
		return nil, err
	}
	return &parsed, nil
}

func Float64(name string, opts ...Option) *Float64Flag {
	f := newFlag[float64](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultFloat64Parser())
	}
	return &Float64Flag{f}
}
