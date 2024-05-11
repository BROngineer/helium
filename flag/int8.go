package flag

import "strconv"

type fint8 = flag[int8]

type Int8 struct {
	*fint8
}

func (f *Int8) Parse(input string) error {
	parsed, err := strconv.ParseInt(input, 10, 8)
	if err != nil {
		return err
	}
	v := int8(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewInt8(name string, opts ...Option) *Int8 {
	f := newFlag[int8](name)
	applyForFlag(f, opts...)
	return &Int8{f}
}
