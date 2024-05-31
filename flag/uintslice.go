package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type uintSlice = flag[[]uint]

type UintSlice struct {
	*uintSlice
}

func (f *UintSlice) FromCommandLine(input string) error {
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]uint, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 32)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, uint(v))
	}
	if f.IsVisited() {
		stored := DerefOrDie[[]uint](f.Value())
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func (f *UintSlice) FromEnvVariable(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]uint, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 32)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, uint(v))
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewUintSlice(name string, opts ...Option) *UintSlice {
	f := newFlag[[]uint](name)
	applyForFlag(f, opts...)
	return &UintSlice{f}
}
