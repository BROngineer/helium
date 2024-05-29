package flag

import (
	"fmt"
	"strconv"
	"strings"
)

type int64Slice = flag[[]int64]

type Int64Slice struct {
	*int64Slice
}

func (f *Int64Slice) Parse(input string) error {
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]int64, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 64)
		if err != nil {
			return err
		}
		parsed = append(parsed, v)
	}
	if f.IsVisited() {
		stored := *f.Value().(*[]int64)
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewInt64Slice(name string, opts ...Option) *Int64Slice {
	f := newFlag[[]int64](name)
	applyForFlag(f, opts...)
	return &Int64Slice{f}
}
