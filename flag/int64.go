package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fint64 = flag[int64]

type Int64 struct {
	*fint64
}

func (f *Int64) FromCommandLine(input string) error {
	if f.IsVisited() {
		return errors.FlagVisited(f.Name())
	}
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	v, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return errors.ParseError(f.Name(), err)
	}
	f.value = &v
	f.visited = true
	return nil
}

func (f *Int64) FromEnvVariable(input string) error {
	var (
		parsed int64
		err    error
	)
	if parsed, err = strconv.ParseInt(input, 10, 64); err != nil {
		return errors.ParseError(f.Name(), err)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewInt64(name string, opts ...Option) *Int64 {
	f := newFlag[int64](name)
	applyForFlag(f, opts...)
	return &Int64{f}
}
