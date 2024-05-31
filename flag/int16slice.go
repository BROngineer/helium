package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type int16Slice = flag[[]int16]

type Int16Slice struct {
	*int16Slice
}

func (f *Int16Slice) FromCommandLine(input string) error {
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]int16, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 16)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, int16(v))
	}
	if f.IsVisited() {
		stored := DerefOrDie[[]int16](f.Value())
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func (f *Int16Slice) FromEnvVariable(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]int16, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 8)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, int16(v))
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewInt16Slice(name string, opts ...Option) *Int16Slice {
	f := newFlag[[]int16](name)
	applyForFlag(f, opts...)
	return &Int16Slice{f}
}
