package flag

import (
	"fmt"
	"strconv"
)

type fuint = flag[uint]

type Uint struct {
	*fuint
}

func (f *Uint) Parse(input string) error {
	if f.IsVisited() {
		return fmt.Errorf("flag already parsed")
	}
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
	}
	parsed, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return err
	}
	v := uint(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewUint(name string, opts ...Option) *Uint {
	f := newFlag[uint](name)
	applyForFlag(f, opts...)
	return &Uint{f}
}
