package flag

import (
	"strconv"
	"strings"
)

type uint8Slice = flag[[]uint8]

type Uint8Slice struct {
	*uint8Slice
}

func (f *Uint8Slice) Parse(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]uint8, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseUint(el, 10, 8)
		if err != nil {
			return err
		}
		parsed = append(parsed, uint8(v))
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewUint8Slice(name string, opts ...Option) *Uint8Slice {
	f := newFlag[[]uint8](name)
	applyForFlag(f, opts...)
	return &Uint8Slice{f}
}
