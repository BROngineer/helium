package flag

import "fmt"

type fstring = flag[string]

type String struct {
	*fstring
}

func (f *String) Parse(input string) error {
	if f.IsVisited() {
		return fmt.Errorf("flag already parsed")
	}
	var empty string
	if input == empty {
		return fmt.Errorf("no value provided")
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
