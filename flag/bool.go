package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fbool = flag[bool]

type Bool struct {
	*fbool
}

func (f *Bool) FromCommandLine(input string) error {
	var (
		v   bool
		err error
	)
	if f.IsVisited() {
		return errors.FlagVisited(f.Name())
	}
	switch input {
	case "":
		v = true
	default:
		v, err = strconv.ParseBool(input)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
	}
	f.value = &v
	f.visited = true
	return nil
}

func (f *Bool) FromEnvVariable(input string) error {
	var (
		parsed bool
		err    error
	)
	if parsed, err = strconv.ParseBool(input); err != nil {
		return errors.ParseError(f.Name(), err)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewBool(name string, opts ...Option) *Bool {
	f := newFlag[bool](name)
	applyForFlag(f, opts...)
	return &Bool{f}
}
