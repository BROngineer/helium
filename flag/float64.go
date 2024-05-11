package flag

import "strconv"

type ffloat64 = flag[float64]

type Float64 struct {
	*ffloat64
}

func (f *Float64) Parse(input string) error {
	v, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return err
	}
	f.value = &v
	f.visited = true
	return nil
}

func NewFloat64(name string, opts ...Option) *Float64 {
	f := newFlag[float64](name)
	applyForFlag(f, opts...)
	return &Float64{f}
}
