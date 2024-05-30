package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fuint32 = flag[uint32]

type Uint32 struct {
	*fuint32
}

func (f *Uint32) Parse(input string) error {
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
	v := uint32(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewUint32(name string, opts ...Option) *Uint32 {
	f := newFlag[uint32](name)
	applyForFlag(f, opts...)
	return &Uint32{f}
}
