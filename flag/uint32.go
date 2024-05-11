package flag

import "strconv"

type fuint32 = flag[uint32]

type Uint32 struct {
	*fuint32
}

func (f *Uint32) Parse(input string) error {
	parsed, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return err
	}
	v := uint32(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewUint32(name string, opts ...Option) *Uint32 {
	f := newFlag[uint32](name)
	applyForFlag(f, opts...)
	return &Uint32{f}
}
