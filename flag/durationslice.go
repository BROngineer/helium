package flag

import (
	"fmt"
	"strings"
	"time"
)

type durationSlice = flag[[]time.Duration]

type DurationSlice struct {
	*durationSlice
}

func (f *DurationSlice) Parse(input string) error {
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]time.Duration, 0, len(s))
	for _, el := range s {
		v, err := time.ParseDuration(el)
		if err != nil {
			return err
		}
		parsed = append(parsed, v)
	}
	if f.IsVisited() {
		stored := *f.Value().(*[]time.Duration)
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
