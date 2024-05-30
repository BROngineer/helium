package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fint16 = flag[int16]

type Int16 struct {
	*fint16
}

func (f *Int16) Parse(input string) error {
	if f.IsVisited() {
		return errors.FlagVisited(f.Name())
	}
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	parsed, err := strconv.ParseInt(input, 10, 16)
	if err != nil {
		return errors.ParseError(f.Name(), err)
	}
	v := int16(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewInt16(name string, opts ...Option) *Int16 {
	f := newFlag[int16](name)
	applyForFlag(f, opts...)
	return &Int16{f}
}
