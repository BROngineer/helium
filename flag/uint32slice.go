package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type uint32Slice = flag[[]uint32]

type Uint32Slice struct {
	*uint32Slice
}

type uint32SliceParser struct {
	*embeddedParser
}

func defaultUint32SliceParser() *uint32SliceParser {
	return &uint32SliceParser{&embeddedParser{}}
}

func (p *uint32SliceParser) ParseCmd(input string) (any, error) {
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	s := strings.Split(input, p.Separator())
	parsed := make([]uint32, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 32)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, uint32(v))
	}
	if p.IsVisited() {
		stored := DerefOrDie[[]uint32](p.CurrentValue())
		parsed = append(stored, parsed...)
	}
	return &parsed, nil
}

func (p *uint32SliceParser) ParseEnv(input string) (any, error) {
	s := strings.Split(input, p.Separator())
	parsed := make([]uint32, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 32)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, uint32(v))
	}
	return &parsed, nil
}

func NewUint32Slice(name string, opts ...Option) *Uint32Slice {
	f := newFlag[[]uint32](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultUint32SliceParser())
	}
	return &Uint32Slice{f}
}
