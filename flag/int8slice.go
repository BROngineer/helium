package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type int8Slice = flag[[]int8]

type Int8SliceFlag struct {
	*int8Slice
}

type int8SliceParser struct {
	*embeddedParser
}

func defaultInt8SliceParser() *int8SliceParser {
	return &int8SliceParser{&embeddedParser{}}
}

func (p *int8SliceParser) ParseCmd(input string) (any, error) {
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	s := strings.Split(input, p.Separator())
	parsed := make([]int8, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 8)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, int8(v))
	}
	if p.IsSetFromCmd() {
		stored := DerefOrDie[[]int8](p.CurrentValue())
		parsed = append(stored, parsed...)
	}
	return &parsed, nil
}

func (p *int8SliceParser) ParseEnv(input string) (any, error) {
	s := strings.Split(input, p.Separator())
	parsed := make([]int8, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 8)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, int8(v))
	}
	return &parsed, nil
}

func Int8Slice(name string, opts ...Option) *Int8SliceFlag {
	f := newFlag[[]int8](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultInt8SliceParser())
	}
	return &Int8SliceFlag{f}
}
