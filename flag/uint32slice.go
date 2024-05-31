package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type uint32Slice = flag[[]uint32]

type Uint32Slice struct {
	*uint32Slice
}

func (f *Uint32Slice) FromCommandLine(input string) error {
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]uint32, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 32)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, uint32(v))
	}
	if f.IsVisited() {
		stored := DerefOrDie[[]uint32](f.Value())
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func (f *Uint32Slice) FromEnvVariable(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]uint32, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 8)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, uint32(v))
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewUint32Slice(name string, opts ...Option) *Uint32Slice {
	f := newFlag[[]uint32](name)
	applyForFlag(f, opts...)
	return &Uint32Slice{f}
}
