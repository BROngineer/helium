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

func (f *Float64Slice) FromCommandLine(input string) error {
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]float64, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseFloat(el, 64)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, v)
	}
	if f.IsVisited() {
		stored := DerefOrDie[[]float64](f.Value())
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func (f *Float64Slice) FromEnvVariable(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]float64, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseFloat(el, 32)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, v)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewFloat64Slice(name string, opts ...Option) *Float64Slice {
	f := newFlag[[]float64](name)
	applyForFlag(f, opts...)
	return &Float64Slice{f}
}
