package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type float64Slice = flag[[]float64]

type Float64Slice struct {
	*float64Slice
}

type float64SliceParser struct {
	*embeddedParser
}

func defaultFloat64SliceParser() *float64SliceParser {
	return &float64SliceParser{&embeddedParser{}}
}

func (p *float64SliceParser) ParseCmd(input string) (any, error) {
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	s := strings.Split(input, p.Separator())
	parsed := make([]float64, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseFloat(el, 64)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, v)
	}
	if p.IsVisited() {
		stored := DerefOrDie[[]float64](p.CurrentValue())
		parsed = append(stored, parsed...)
	}
	return &parsed, nil
}

func (p *float64SliceParser) ParseEnv(input string) (any, error) {
	s := strings.Split(input, p.Separator())
	parsed := make([]float64, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseFloat(el, 64)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, v)
	}
	return &parsed, nil
}

func NewFloat64Slice(name string, opts ...Option) *Float64Slice {
	f := newFlag[[]float64](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultFloat64SliceParser())
	}
	return &Float64Slice{f}
}
