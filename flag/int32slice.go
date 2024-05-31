package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type int32Slice = flag[[]int32]

type Int32Slice struct {
	*int32Slice
}

func (f *Int32Slice) FromCommandLine(input string) error {
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]int32, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 32)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, int32(v))
	}
	if f.IsVisited() {
		stored := DerefOrDie[[]int32](f.Value())
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func (f *Int32Slice) FromEnvVariable(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]int32, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 8)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, int32(v))
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewInt32Slice(name string, opts ...Option) *Int32Slice {
	f := newFlag[[]int32](name)
	applyForFlag(f, opts...)
	return &Int32Slice{f}
}
