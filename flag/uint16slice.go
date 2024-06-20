package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type uint16Slice = flag[[]uint16]

type Uint16Slice struct {
	*uint16Slice
}

type uint16SliceParser struct {
	*embeddedParser
}

func defaultUint16SliceParser() *uint16SliceParser {
	return &uint16SliceParser{&embeddedParser{}}
}

func (p *uint16SliceParser) ParseCmd(input string) (any, error) {
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	s := strings.Split(input, p.Separator())
	parsed := make([]uint16, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 16)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, uint16(v))
	}
	if p.IsVisited() {
		stored := DerefOrDie[[]uint16](p.CurrentValue())
		parsed = append(stored, parsed...)
	}
	return &parsed, nil
}

func (p *uint16SliceParser) ParseEnv(input string) (any, error) {
	s := strings.Split(input, p.Separator())
	parsed := make([]uint16, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 16)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, uint16(v))
	}
	return &parsed, nil
}

func NewUint16Slice(name string, opts ...Option) *Uint16Slice {
	f := newFlag[[]uint16](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultUint16SliceParser())
	}
	return &Uint16Slice{f}
}
