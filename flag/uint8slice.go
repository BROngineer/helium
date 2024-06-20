package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type uint8Slice = flag[[]uint8]

type Uint8Slice struct {
	*uint8Slice
}

type uint8SliceParser struct {
	*embeddedParser
}

func defaultUint8SliceParser() *uint8SliceParser {
	return &uint8SliceParser{&embeddedParser{}}
}

func (p *uint8SliceParser) ParseCmd(input string) (any, error) {
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	s := strings.Split(input, p.Separator())
	parsed := make([]uint8, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 8)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, uint8(v))
	}
	if p.IsVisited() {
		stored := DerefOrDie[[]uint8](p.CurrentValue())
		parsed = append(stored, parsed...)
	}
	return &parsed, nil
}

func (p *uint8SliceParser) ParseEnv(input string) (any, error) {
	s := strings.Split(input, p.Separator())
	parsed := make([]uint8, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 8)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, uint8(v))
	}
	return &parsed, nil
}

func NewUint8Slice(name string, opts ...Option) *Uint8Slice {
	f := newFlag[[]uint8](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultUint8SliceParser())
	}
	return &Uint8Slice{f}
}
