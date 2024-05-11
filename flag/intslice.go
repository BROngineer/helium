package flag

import (
	"strconv"
	"strings"
)

type intSlice = flag[[]int]

type IntSlice struct {
	*intSlice
}

func (f *IntSlice) Parse(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]int, 0, len(s))
	for _, el := range s {
		v, err := strconv.Atoi(el)
		if err != nil {
			return err
		}
		parsed = append(parsed, v)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewIntSlice(name string, opts ...Option) *IntSlice {
	f := newFlag[[]int](name)
	applyForFlag(f, opts...)
	return &IntSlice{f}
}
