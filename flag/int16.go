package flag

import (
	"fmt"
	"strconv"
)

type fint16 = flag[int16]

type Int16 struct {
	*fint16
}

func (f *Int16) Parse(input string) error {
	if f.IsVisited() {
		return fmt.Errorf("flag already parsed")
	}
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
	}
	parsed, err := strconv.ParseInt(input, 10, 16)
	if err != nil {
		return err
	}
	v := int16(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewInt16(name string, opts ...Option) *Int16 {
	f := newFlag[int16](name)
	applyForFlag(f, opts...)
	return &Int16{f}
}
