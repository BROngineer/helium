package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type intSlice = flag[[]int]

type IntSlice struct {
	*intSlice
}

type intSliceParser struct {
	*embeddedParser
}

func defaultIntSliceParser() *intSliceParser {
	return &intSliceParser{&embeddedParser{}}
}

func (p *intSliceParser) ParseCmd(input string) (any, error) {
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	s := strings.Split(input, p.Separator())
	parsed := make([]int, 0, len(s))
	for _, el := range s {
		v, err := strconv.Atoi(el)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, v)
	}
	if p.IsVisited() {
		stored := DerefOrDie[[]int](p.CurrentValue())
		parsed = append(stored, parsed...)
	}
	return &parsed, nil
}

func (p *intSliceParser) ParseEnv(input string) (any, error) {
	p.visited = true
	s := strings.Split(input, p.Separator())
	parsed := make([]int, 0, len(s))
	for _, el := range s {
		v, err := strconv.Atoi(el)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, v)
	}
	return &parsed, nil
}

func NewIntSlice(name string, opts ...Option) *IntSlice {
	f := newFlag[[]int](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultIntSliceParser())
	}
	return &IntSlice{f}
}
