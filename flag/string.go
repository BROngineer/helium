package flag

type fstring = flag[string]

type String struct {
	*fstring
}

func (f *String) Parse(input string) error {
	f.value = &input
	f.visited = true
	return nil
}

func NewString(name string, opts ...Option) *String {
	f := newFlag[string](name)
	applyForFlag(f, opts...)
	return &String{f}
}
