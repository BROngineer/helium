package flag

import (
	"fmt"
	"strconv"
	"strings"
)

type uintSlice = flag[[]uint]

type UintSlice struct {
	*uintSlice
}

func (f *UintSlice) Parse(input string) error {
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]uint, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 32)
		if err != nil {
			return err
		}
		parsed = append(parsed, uint(v))
	}
	if f.IsVisited() {
		stored := *f.Value().(*[]uint)
		parsed = append(stored, parsed...)
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
