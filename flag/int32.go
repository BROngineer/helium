package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fint32 = flag[int32]

type Int32Flag struct {
	*fint32
}

type int32Parser struct {
	*embeddedParser
}

func defaultInt32Parser() *int32Parser {
	return &int32Parser{&embeddedParser{}}
}

func (p *int32Parser) ParseCmd(input string) (any, error) {
	if p.IsSetFromCmd() {
		return nil, errors.ErrFlagVisited
	}
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	v, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return nil, err
	}
	parsed := int32(v)
	return &parsed, nil
}

func (p *int32Parser) ParseEnv(value string) (any, error) {
	var (
		v   int64
		err error
	)
	if v, err = strconv.ParseInt(value, 10, 32); err != nil {
		return nil, err
	}
	parsed := int32(v)
	return &parsed, nil
}

func Int32(name string, opts ...Option) *Int32Flag {
	f := newFlag[int32](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultInt32Parser())
	}
	return &Int32Flag{f}
}
