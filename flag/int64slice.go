package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type int64Slice = flag[[]int64]

type Int64Slice struct {
	*int64Slice
}

type int64SliceParser struct {
	*embeddedParser
}

func defaultInt64SliceParser() *int64SliceParser {
	return &int64SliceParser{&embeddedParser{}}
}

func (p *int64SliceParser) ParseCmd(input string) (any, error) {
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	s := strings.Split(input, p.Separator())
	parsed := make([]int64, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 64)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, v)
	}
	if p.IsVisited() {
		stored := DerefOrDie[[]int64](p.CurrentValue())
		parsed = append(stored, parsed...)
	}
	return &parsed, nil
}

func (p *int64SliceParser) ParseEnv(input string) (any, error) {
	s := strings.Split(input, p.Separator())
	parsed := make([]int64, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 64)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, v)
	}
	return &parsed, nil
}

func NewInt64Slice(name string, opts ...Option) *Int64Slice {
	f := newFlag[[]int64](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultInt64SliceParser())
	}
	return &Int64Slice{f}
}
