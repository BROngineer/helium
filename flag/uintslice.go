package flag

import (
	"strconv"
	"strings"
)

type uintSlice = flag[[]uint]

type UintSlice struct {
	*uintSlice
}

func (f *UintSlice) Parse(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]uint, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 32)
		if err != nil {
			return err
		}
		parsed = append(parsed, uint(v))
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewUintSlice(name string, opts ...Option) *UintSlice {
	f := newFlag[[]uint](name)
	applyForFlag(f, opts...)
	return &UintSlice{f}
}
