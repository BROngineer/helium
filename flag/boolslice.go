package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type boolSlice = flag[[]bool]

type BoolSlice struct {
	*boolSlice
}

func (f *BoolSlice) Parse(input string) error {
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]bool, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseBool(el)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, v)
	}
	if f.IsVisited() {
		stored := DerefOrDie[[]bool](f.Value())
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewBoolSlice(name string, opts ...Option) *BoolSlice {
	f := newFlag[[]bool](name)
	applyForFlag(f, opts...)
	return &BoolSlice{f}
}
