package flag

import (
	"strconv"
	"strings"
)

type uint32Slice = flag[[]uint32]

type Uint32Slice struct {
	*uint32Slice
}

func (f *Uint32Slice) Parse(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]uint32, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 32)
		if err != nil {
			return err
		}
		parsed = append(parsed, uint32(v))
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewUint32Slice(name string, opts ...Option) *Uint32Slice {
	f := newFlag[[]uint32](name)
	applyForFlag(f, opts...)
	return &Uint32Slice{f}
}
