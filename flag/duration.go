package flag

import (
	"time"

	"github.com/brongineer/helium/errors"
)

type duration = flag[time.Duration]

type Duration struct {
	*duration
}

func (f *Duration) Parse(input string) error {
	if f.IsVisited() {
		return errors.FlagVisited(f.Name())
	}
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	v, err := time.ParseDuration(input)
	if err != nil {
		return errors.ParseError(f.Name(), err)
	}
	f.value = &v
	f.visited = true
	return nil
}

func NewDuration(name string, opts ...Option) *Duration {
	f := newFlag[time.Duration](name)
	applyForFlag(f, opts...)
	return &Duration{f}
}
