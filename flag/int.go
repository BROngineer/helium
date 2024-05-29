package flag

import (
	"fmt"
	"strconv"
)

type fint = flag[int]

type Int struct {
	*fint
}

func (f *Int) Parse(input string) error {
	if f.IsVisited() {
		return fmt.Errorf("flag already parsed")
	}
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
	}
	v, err := strconv.Atoi(input)
	if err != nil {
		return err
	}
	f.value = &v
	f.visited = true
	return nil
}

func NewInt(name string, opts ...Option) *Int {
	f := newFlag[int](name)
	applyForFlag(f, opts...)
	return &Int{f}
}
