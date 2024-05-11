package flag

import "strconv"

type ffloat32 = flag[float32]

type Float32 struct {
	*ffloat32
}

func (f *Float32) Parse(input string) error {
	parsed, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return err
	}
	v := float32(parsed)
	f.value = &v
	f.visited = true
	return nil
}

func NewFloat32(name string, opts ...Option) *Float32 {
	f := newFlag[float32](name)
	applyForFlag(f, opts...)
	return &Float32{f}
}
