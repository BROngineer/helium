package flag

import (
	"fmt"
	"strings"
)

type stringSlice = flag[[]string]

type StringSlice struct {
	*stringSlice
}

func (f *StringSlice) Parse(input string) error {
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
	}
	v := strings.Split(input, f.Separator())
	if f.IsVisited() {
		stored := *f.Value().(*[]string)
		v = append(stored, v...)
	}
	f.value = &v
	f.visited = true
	return nil
}

func NewStringSlice(name string, opts ...Option) *StringSlice {
	f := newFlag[[]string](name)
	applyForFlag(f, opts...)
	return &StringSlice{f}
}
