package flag

import (
	"strconv"
	"strings"
)

type float64Slice = flag[[]float64]

type Float64Slice struct {
	*float64Slice
}

func (f *Float64Slice) Parse(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]float64, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseFloat(el, 64)
		if err != nil {
			return err
		}
		parsed = append(parsed, v)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewFloat64Slice(name string, opts ...Option) *Float64Slice {
	f := newFlag[[]float64](name)
	applyForFlag(f, opts...)
	return &Float64Slice{f}
}
