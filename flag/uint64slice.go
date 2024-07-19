package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type uint64Slice = flag[[]uint64]

type Uint64SliceFlag struct {
	*uint64Slice
}

type uint64SliceParser struct {
	*embeddedParser
}

func defaultUint64SliceParser() *uint64SliceParser {
	return &uint64SliceParser{&embeddedParser{}}
}

func (p *uint64SliceParser) ParseCmd(input string) (any, error) {
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	s := strings.Split(input, p.Separator())
	parsed := make([]uint64, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 64)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, v)
	}
	if p.IsSetFromCmd() {
		stored := DerefOrDie[[]uint64](p.CurrentValue())
		parsed = append(stored, parsed...)
	}
	return &parsed, nil
}

func (p *uint64SliceParser) ParseEnv(input string) (any, error) {
	s := strings.Split(input, p.Separator())
	parsed := make([]uint64, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 64)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, v)
	}
	return &parsed, nil
}

func Uint64Slice(name string, opts ...Option) *Uint64SliceFlag {
	f := newFlag[[]uint64](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultUint64SliceParser())
	}
	return &Uint64SliceFlag{f}
}
