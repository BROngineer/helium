package flag

import (
	"fmt"
	"strconv"
)

type fuint16 = flag[uint16]

type Uint16 struct {
	*fuint16
}

func (f *Uint16) Parse(input string) error {
	if f.IsVisited() {
		return fmt.Errorf("flag already parsed")
	}
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
	}
	parsed, err := strconv.ParseUint(input, 10, 16)
	if err != nil {
		return err
	}
	v := uint16(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewUint16(name string, opts ...Option) *Uint16 {
	f := newFlag[uint16](name)
	applyForFlag(f, opts...)
	return &Uint16{f}
}
