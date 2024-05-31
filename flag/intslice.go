package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type intSlice = flag[[]int]

type IntSlice struct {
	*intSlice
}

func (f *IntSlice) FromCommandLine(input string) error {
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]int, 0, len(s))
	for _, el := range s {
		v, err := strconv.Atoi(el)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, v)
	}
	if f.IsVisited() {
		stored := DerefOrDie[[]int](f.Value())
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func (f *IntSlice) FromEnvVariable(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]int, 0, len(s))
	for _, el := range s {
		v, err := strconv.Atoi(el)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, v)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewIntSlice(name string, opts ...Option) *IntSlice {
	f := newFlag[[]int](name)
	applyForFlag(f, opts...)
	return &IntSlice{f}
}
