package flag

import (
	"strconv"
	"strings"
)

type int16Slice = flag[[]int16]

type Int16Slice struct {
	*int16Slice
}

func (f *Int16Slice) Parse(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]int16, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseInt(el, 10, 16)
		if err != nil {
			return err
		}
		parsed = append(parsed, int16(v))
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewInt16Slice(name string, opts ...Option) *Int16Slice {
	f := newFlag[[]int16](name)
	applyForFlag(f, opts...)
	return &Int16Slice{f}
}
