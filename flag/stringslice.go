package flag

import (
	"strings"

	"github.com/brongineer/helium/errors"
)

type stringSlice = flag[[]string]

type StringSlice struct {
	*stringSlice
}

func (f *StringSlice) FromCommandLine(input string) error {
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	v := strings.Split(input, f.Separator())
	if f.IsVisited() {
		stored := DerefOrDie[[]string](f.Value())
		v = append(stored, v...)
	}
	f.value = &v
	f.visited = true
	return nil
}

func (f *StringSlice) FromEnvVariable(input string) error {
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
