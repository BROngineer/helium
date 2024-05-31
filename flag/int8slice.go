package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type int8Slice = flag[[]int8]

type Int8Slice struct {
	*int8Slice
}

func (f *Int8Slice) FromCommandLine(input string) error {
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]int8, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 8)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, int8(v))
	}
	if f.IsVisited() {
		stored := DerefOrDie[[]int8](f.Value())
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func (f *Int8Slice) FromEnvVariable(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]int8, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 8)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, int8(v))
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewInt8Slice(name string, opts ...Option) *Int8Slice {
	f := newFlag[[]int8](name)
	applyForFlag(f, opts...)
	return &Int8Slice{f}
}
