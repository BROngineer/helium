package flag

import (
	"strings"
	"time"

	"github.com/brongineer/helium/errors"
)

type durationSlice = flag[[]time.Duration]

type DurationSlice struct {
	*durationSlice
}

func (f *DurationSlice) Parse(input string) error {
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]time.Duration, 0, len(s))
	for _, el := range s {
		v, err := time.ParseDuration(el)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, v)
	}
	if f.IsVisited() {
		stored := DerefOrDie[[]time.Duration](f.Value())
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewDurationSlice(name string, opts ...Option) *DurationSlice {
	f := newFlag[[]time.Duration](name)
	applyForFlag(f, opts...)
	return &DurationSlice{f}
}
