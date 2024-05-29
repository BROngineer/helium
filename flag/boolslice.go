package flag

import (
	"fmt"
	"strconv"
	"strings"
)

type boolSlice = flag[[]bool]

type BoolSlice struct {
	*boolSlice
}

func (f *BoolSlice) Parse(input string) error {
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
	}
	s := strings.Split(input, f.Separator())
	parsed := make([]bool, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseBool(el)
		if err != nil {
			return err
		}
		parsed = append(parsed, v)
	}
	if f.IsVisited() {
		stored := *f.Value().(*[]bool)
		parsed = append(stored, parsed...)
	}
	f.value = &parsed
	f.visited = true
	return nil
}

func NewBoolSlice(name string, opts ...Option) *BoolSlice {
	f := newFlag[[]bool](name)
	applyForFlag(f, opts...)
	return &BoolSlice{f}
}
