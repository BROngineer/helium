package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type int32Slice = flag[[]int32]

type Int32Slice struct {
	*int32Slice
}

type int32SliceParser struct {
	*embeddedParser
}

func defaultInt32SliceParser() *int32SliceParser {
	return &int32SliceParser{&embeddedParser{}}
}

func (p *int32SliceParser) ParseCmd(input string) (any, error) {
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	s := strings.Split(input, p.Separator())
	parsed := make([]int32, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 32)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, int32(v))
	}
	if p.IsVisited() {
		stored := DerefOrDie[[]int32](p.CurrentValue())
		parsed = append(stored, parsed...)
	}
	return &parsed, nil
}

func (p *int32SliceParser) ParseEnv(input string) (any, error) {
	s := strings.Split(input, p.Separator())
	parsed := make([]int32, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 32)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, int32(v))
	}
	return &parsed, nil
}

func NewInt32Slice(name string, opts ...Option) *Int32Slice {
	f := newFlag[[]int32](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultInt32SliceParser())
	}
	return &Int32Slice{f}
}
