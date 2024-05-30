package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fuint16 = flag[uint16]

type Uint16 struct {
	*fuint16
}

func (f *Uint16) Parse(input string) error {
	if f.IsVisited() {
		return errors.FlagVisited(f.Name())
	}
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	parsed, err := strconv.ParseUint(input, 10, 16)
	if err != nil {
		return errors.ParseError(f.Name(), err)
	}
	v := uint16(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewUint16(name string, opts ...Option) *Uint16 {
	f := newFlag[uint16](name)
	applyForFlag(f, opts...)
	return &Uint16{f}
}
