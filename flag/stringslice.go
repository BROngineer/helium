package flag

import "strings"

type stringSlice = flag[[]string]

type StringSlice struct {
	*stringSlice
}

func (f *StringSlice) Parse(input string) error {
	v := strings.Split(input, f.Separator())
	f.value = &v
	f.visited = true
	return nil
}

func NewStringSlice(name string, opts ...Option) *StringSlice {
	f := newFlag[[]string](name)
	applyForFlag(f, opts...)
	return &StringSlice{f}
}
