package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fint64 = flag[int64]

type Int64Flag struct {
	*fint64
}

type int64Parser struct {
	*embeddedParser
}

func defaultInt64Parser() *int64Parser {
	return &int64Parser{&embeddedParser{}}
}

func (p *int64Parser) ParseCmd(input string) (any, error) {
	if p.IsSetFromCmd() {
		return nil, errors.ErrFlagVisited
	}
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	v, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return nil, err
	}
	parsed := v
	return &parsed, nil
}

func (p *int64Parser) ParseEnv(value string) (any, error) {
	var (
		parsed int64
		err    error
	)
	if parsed, err = strconv.ParseInt(value, 10, 64); err != nil {
		return nil, err
	}
	return &parsed, nil
}

func Int64(name string, opts ...Option) *Int64Flag {
	f := newFlag[int64](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultInt64Parser())
	}
	return &Int64Flag{f}
}
