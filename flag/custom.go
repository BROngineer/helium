package flag

type Custom[T any] struct {
	*flag[T]
}

func NewCustom[T any](name string, opts ...Option) *Custom[T] {
	f := newFlag[T](name)
	applyForFlag(f, opts...)
	return &Custom[T]{f}
}
