package flag

import "strconv"

type fint64 = flag[int64]

type Int64 struct {
	*fint64
}

func (f *Int64) Parse(input string) error {
	v, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return err
	}
	f.value = &v
	f.visited = true
	return nil
}

func NewInt64(name string, opts ...Option) *Int64 {
	f := newFlag[int64](name)
	applyForFlag(f, opts...)
	return &Int64{f}
}
