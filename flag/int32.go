package flag

import (
	"fmt"
	"strconv"
)

type fint32 = flag[int32]

type Int32 struct {
	*fint32
}

func (f *Int32) Parse(input string) error {
	if f.IsVisited() {
		return fmt.Errorf("flag already parsed")
	}
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
	}
	parsed, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return err
	}
	v := int32(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewInt32(name string, opts ...Option) *Int32 {
	f := newFlag[int32](name)
	applyForFlag(f, opts...)
	return &Int32{f}
}
