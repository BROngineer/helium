package flag

import "strconv"

type fint = flag[int]

type Int struct {
	*fint
}

func (f *Int) Parse(input string) error {
	v, err := strconv.Atoi(input)
	if err != nil {
		return err
	}
	f.value = &v
	f.visited = true
	return nil
}

func NewInt(name string, opts ...Option) *Int {
	f := newFlag[int](name)
	applyForFlag(f, opts...)
	return &Int{f}
}
