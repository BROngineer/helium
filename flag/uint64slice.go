package flag

import (
	"strconv"
	"strings"
)

type uint64Slice = flag[[]uint64]

type Uint64Slice struct {
	*uint64Slice
}

func (f *Uint64Slice) Parse(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]uint64, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 64)
		if err != nil {
			return err
		}
		parsed = append(parsed, v)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewUint64Slice(name string, opts ...Option) *Uint64Slice {
	f := newFlag[[]uint64](name)
	applyForFlag(f, opts...)
	return &Uint64Slice{f}
}
