package flag

import (
	"strconv"

	"github.com/brongineer/helium/errors"
)

type fint32 = flag[int32]

type Int32 struct {
	*fint32
}

func (f *Int32) Parse(input string) error {
	if f.IsVisited() {
		return errors.FlagVisited(f.Name())
	}
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	parsed, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return errors.ParseError(f.Name(), err)
	}
	v := int32(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewInt32(name string, opts ...Option) *Int32 {
	f := newFlag[int32](name)
	applyForFlag(f, opts...)
	return &Int32{f}
}
