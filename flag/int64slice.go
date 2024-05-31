package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type int64Slice = flag[[]int64]

type Int64Slice struct {
	*int64Slice
}

func (f *Int64Slice) FromCommandLine(input string) error {
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]int64, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 64)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, v)
	}
	if f.IsVisited() {
		stored := DerefOrDie[[]int64](f.Value())
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func (f *Int64Slice) FromEnvVariable(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]int64, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 8)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, v)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewInt64Slice(name string, opts ...Option) *Int64Slice {
	f := newFlag[[]int64](name)
	applyForFlag(f, opts...)
	return &Int64Slice{f}
}
