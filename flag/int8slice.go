package flag

import (
	"fmt"
	"strconv"
	"strings"
)

type int8Slice = flag[[]int8]

type Int8Slice struct {
	*int8Slice
}

func (f *Int8Slice) Parse(input string) error {
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]int8, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 8)
		if err != nil {
			return err
		}
		parsed = append(parsed, int8(v))
	}
	if f.IsVisited() {
		stored := *f.Value().(*[]int8)
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewInt8Slice(name string, opts ...Option) *Int8Slice {
	f := newFlag[[]int8](name)
	applyForFlag(f, opts...)
	return &Int8Slice{f}
}
