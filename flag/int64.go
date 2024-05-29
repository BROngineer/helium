package flag

import (
	"fmt"
	"strconv"
)

type fint64 = flag[int64]

type Int64 struct {
	*fint64
}

func (f *Int64) Parse(input string) error {
	if f.IsVisited() {
		return fmt.Errorf("flag already parsed")
	}
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
	}
	v, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return err
	}
	f.value = &v
	f.visited = true
	return nil
}

func NewInt64(name string, opts ...Option) *Int64 {
	f := newFlag[int64](name)
	applyForFlag(f, opts...)
	return &Int64{f}
}
