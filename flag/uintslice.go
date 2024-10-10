package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type uintSlice = flag[[]uint]

type UintSliceFlag struct {
	*uintSlice
}

type uintSliceParser struct {
	*embeddedParser
}

func defaultUintSliceParser() *uintSliceParser {
	return &uintSliceParser{&embeddedParser{}}
}

func (p *uintSliceParser) ParseCmd(input string) (any, error) {
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	s := strings.Split(input, p.Separator())
	parsed := make([]uint, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 32)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, uint(v))
	}
	if p.IsSetFromCmd() {
		stored := DerefOrDie[[]uint](p.CurrentValue())
		parsed = append(stored, parsed...)
	}
	return &parsed, nil
}

func (p *uintSliceParser) ParseEnv(input string) (any, error) {
	s := strings.Split(input, p.Separator())
	parsed := make([]uint, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 32)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, uint(v))
	}
	return &parsed, nil
}

func UintSlice(name string, opts ...Option) *UintSliceFlag {
	f := newFlag[[]uint](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultUintSliceParser())
	}
	return &UintSliceFlag{f}
}
