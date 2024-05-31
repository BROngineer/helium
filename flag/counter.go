package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type counter = flag[int]

type Counter struct {
	*counter
}

func (f *Counter) FromCommandLine(_ string) error {
	current := DerefOrDie[int](f.Value())
	v := current + 1
	f.value = &v
	f.visited = true
	return nil
}

func (f *Counter) FromEnvVariable(input string) error {
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

func NewCounter(name string, opts ...Option) *Counter {
	f := newFlag[int](name)
	applyForFlag(f, opts...)
	if f.defaultValue == nil {
		f.setDefaultValue(0)
	}
	return &Counter{f}
}
