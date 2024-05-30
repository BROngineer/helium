package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fuint = flag[uint]

type Uint struct {
	*fuint
}

func (f *Uint) Parse(input string) error {
	if f.IsVisited() {
		return errors.FlagVisited(f.Name())
	}
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	parsed, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return errors.ParseError(f.Name(), err)
	}
	v := uint(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewUint(name string, opts ...Option) *Uint {
	f := newFlag[uint](name)
	applyForFlag(f, opts...)
	return &Uint{f}
}
