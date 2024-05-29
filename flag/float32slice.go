package flag

import (
	"fmt"
	"strconv"
	"strings"
)

type float32Slice = flag[[]float32]

type Float32Slice struct {
	*float32Slice
}

func (f *Float32Slice) Parse(input string) error {
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]float32, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseFloat(el, 32)
		if err != nil {
			return err
		}
		parsed = append(parsed, float32(v))
	}
	if f.IsVisited() {
		stored := *f.Value().(*[]float32)
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewFloat32Slice(name string, opts ...Option) *Float32Slice {
	f := newFlag[[]float32](name)
	applyForFlag(f, opts...)
	return &Float32Slice{f}
}
