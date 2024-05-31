package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fint8 = flag[int8]

type Int8 struct {
	*fint8
}

func (f *Int8) FromCommandLine(input string) error {
	if f.IsVisited() {
		return errors.FlagVisited(f.Name())
	}
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	parsed, err := strconv.ParseInt(input, 10, 8)
	if err != nil {
		return errors.ParseError(f.Name(), err)
	}
	v := int8(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func (f *Int8) FromEnvVariable(value string) error {
	var (
		parsed int64
		err    error
	)
	if parsed, err = strconv.ParseInt(value, 10, 8); err != nil {
		return errors.ParseError(f.Name(), err)
	}
	v := int8(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewInt8(name string, opts ...Option) *Int8 {
	f := newFlag[int8](name)
	applyForFlag(f, opts...)
	return &Int8{f}
}
