package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type ffloat64 = flag[float64]

type Float64 struct {
	*ffloat64
}

func (f *Float64) FromCommandLine(input string) error {
	if f.IsVisited() {
		return errors.FlagVisited(f.Name())
	}
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	v, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return errors.ParseError(f.Name(), err)
	}
	f.value = &v
	f.visited = true
	return nil
}

func (f *Float64) FromEnvVariable(input string) error {
	var (
		parsed float64
		err    error
	)
	if parsed, err = strconv.ParseFloat(input, 32); err != nil {
		return errors.ParseError(f.Name(), err)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewFloat64(name string, opts ...Option) *Float64 {
	f := newFlag[float64](name)
	applyForFlag(f, opts...)
	return &Float64{f}
}
