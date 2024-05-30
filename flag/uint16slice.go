package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type uint16Slice = flag[[]uint16]

type Uint16Slice struct {
	*uint16Slice
}

func (f *Uint16Slice) Parse(input string) error {
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]uint16, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 16)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, uint16(v))
	}
	if f.IsVisited() {
		stored := DerefOrDie[[]uint16](f.Value())
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewUint16Slice(name string, opts ...Option) *Uint16Slice {
	f := newFlag[[]uint16](name)
	applyForFlag(f, opts...)
	return &Uint16Slice{f}
}
