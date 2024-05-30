package flag

import (
	"github.com/brongineer/helium/errors"
)

type fstring = flag[string]

type String struct {
	*fstring
}

func (f *String) Parse(input string) error {
	if f.IsVisited() {
		return errors.FlagVisited(f.Name())
	}
	var empty string
	if input == empty {
		return errors.NoValueProvided(f.Name())
	}
	f.value = &input
	f.visited = true
	return nil
}

func NewString(name string, opts ...Option) *String {
	f := newFlag[string](name)
	applyForFlag(f, opts...)
	return &String{f}
}
