package flag

import "strconv"

type fuint = flag[uint]

type Uint struct {
	*fuint
}

func (f *Uint) Parse(input string) error {
	parsed, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return err
	}
	v := uint(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewUint(name string, opts ...Option) *Uint {
	f := newFlag[uint](name)
	applyForFlag(f, opts...)
	return &Uint{f}
}
