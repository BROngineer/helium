package flag

import (
	"fmt"
	"strconv"
	"strings"
)

type int32Slice = flag[[]int32]

type Int32Slice struct {
	*int32Slice
}

func (f *Int32Slice) Parse(input string) error {
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]int32, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 32)
		if err != nil {
			return err
		}
		parsed = append(parsed, int32(v))
	}
	if f.IsVisited() {
		stored := *f.Value().(*[]int32)
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewInt32Slice(name string, opts ...Option) *Int32Slice {
	f := newFlag[[]int32](name)
	applyForFlag(f, opts...)
	return &Int32Slice{f}
}
