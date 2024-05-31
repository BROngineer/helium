package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fint = flag[int]

type Int struct {
	*fint
}

func (f *Int) FromCommandLine(input string) error {
	if f.IsVisited() {
		return errors.FlagVisited(f.Name())
	}
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	v, err := strconv.Atoi(input)
	if err != nil {
		return errors.ParseError(f.Name(), err)
	}
	f.value = &v
	f.visited = true
	return nil
}

func (f *Int) FromEnvVariable(input string) error {
	var (
		parsed int
		err    error
	)
	if parsed, err = strconv.Atoi(input); err != nil {
		return errors.ParseError(f.Name(), err)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewInt(name string, opts ...Option) *Int {
	f := newFlag[int](name)
	applyForFlag(f, opts...)
	return &Int{f}
}
