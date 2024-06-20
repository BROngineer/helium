package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type int16Slice = flag[[]int16]

type Int16Slice struct {
	*int16Slice
}

type int16SliceParser struct {
	*embeddedParser
}

func defaultInt16SliceParser() *int16SliceParser {
	return &int16SliceParser{&embeddedParser{}}
}

func (p *int16SliceParser) ParseCmd(input string) (any, error) {
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	s := strings.Split(input, p.Separator())
	parsed := make([]int16, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 16)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, int16(v))
	}
	if p.IsVisited() {
		stored := DerefOrDie[[]int16](p.CurrentValue())
		parsed = append(stored, parsed...)
	}
	return &parsed, nil
}

func (p *int16SliceParser) ParseEnv(input string) (any, error) {
	s := strings.Split(input, p.Separator())
	parsed := make([]int16, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 16)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, int16(v))
	}
	return &parsed, nil
}

func NewInt16Slice(name string, opts ...Option) *Int16Slice {
	f := newFlag[[]int16](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultInt16SliceParser())
	}
	return &Int16Slice{f}
}
