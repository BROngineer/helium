package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fuint8 = flag[uint8]

type Uint8 struct {
	*fuint8
}

func (f *Uint8) FromCommandLine(input string) error {
	if f.IsVisited() {
		return errors.FlagVisited(f.Name())
	}
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	parsed, err := strconv.ParseUint(input, 10, 8)
	if err != nil {
		return errors.ParseError(f.Name(), err)
	}
	v := uint8(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func (f *Uint8) FromEnvVariable(input string) error {
	var (
		parsed uint64
		err    error
	)
	if parsed, err = strconv.ParseUint(input, 10, 8); err != nil {
		return errors.ParseError(f.Name(), err)
	}
	v := uint8(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewUint8(name string, opts ...Option) *Uint8 {
	f := newFlag[uint8](name)
	applyForFlag(f, opts...)
	return &Uint8{f}
}
