package flag

import "strconv"

type fuint8 = flag[uint8]

type Uint8 struct {
	*fuint8
}

func (f *Uint8) Parse(input string) error {
	parsed, err := strconv.ParseUint(input, 10, 8)
	if err != nil {
		return err
	}
	v := uint8(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewUint8(name string, opts ...Option) *Uint8 {
	f := newFlag[uint8](name)
	applyForFlag(f, opts...)
	return &Uint8{f}
}
