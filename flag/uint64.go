package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fuint64 = flag[uint64]

type Uint64 struct {
	*fuint64
}

func (f *Uint64) Parse(input string) error {
	if f.IsVisited() {
		return errors.FlagVisited(f.Name())
	}
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	v, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return errors.ParseError(f.Name(), err)
	}
	f.value = &v
	f.visited = true
	return nil
}

func NewUint64(name string, opts ...Option) *Uint64 {
	f := newFlag[uint64](name)
	applyForFlag(f, opts...)
	return &Uint64{f}
}
