package flag

import "strconv"

type fint16 = flag[int16]

type Int16 struct {
	*fint16
}

func (f *Int16) Parse(input string) error {
	parsed, err := strconv.ParseInt(input, 10, 16)
	if err != nil {
		return err
	}
	v := int16(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewInt16(name string, opts ...Option) *Int16 {
	f := newFlag[int16](name)
	applyForFlag(f, opts...)
	return &Int16{f}
}
