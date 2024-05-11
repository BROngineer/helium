package flag

type counter = flag[int]

type Counter struct {
	*counter
}

func (f *Counter) Parse(_ string) error {
	current := DerefOrDie[int](f.Value())
	v := current + 1
	f.value = &v
	f.visited = true
	return nil
}

func NewCounter(name string, opts ...Option) *Counter {
	f := newFlag[int](name)
	applyForFlag(f, opts...)
	if f.defaultValue == nil {
		f.setDefaultValue(0)
	}
	return &Counter{f}
}
