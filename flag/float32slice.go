package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type float32Slice = flag[[]float32]

type Float32Slice struct {
	*float32Slice
}

func (f *Float32Slice) FromCommandLine(input string) error {
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]float32, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseFloat(el, 32)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, float32(v))
	}
	if f.IsVisited() {
		stored := DerefOrDie[[]float32](f.Value())
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func (f *Float32Slice) FromEnvVariable(input string) error {
	s := strings.Split(input, f.Separator())
	parsed := make([]float32, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseFloat(el, 32)
		if err != nil {
			return errors.ParseError(f.Name(), err)
		}
		parsed = append(parsed, float32(v))
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
