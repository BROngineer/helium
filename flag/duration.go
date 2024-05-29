package flag

import (
	"fmt"
	"time"
)

type duration = flag[time.Duration]

type Duration struct {
	*duration
}

func (f *Duration) Parse(input string) error {
	if f.IsVisited() {
		return fmt.Errorf("flag already parsed")
	}
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
	}
	v, err := time.ParseDuration(input)
	if err != nil {
		return err
	}
	f.value = &v
	f.visited = true
	return nil
}

func NewDuration(name string, opts ...Option) *Duration {
	f := newFlag[time.Duration](name)
	applyForFlag(f, opts...)
	return &Duration{f}
}
