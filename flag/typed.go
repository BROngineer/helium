package flag

type TypedFlag[T any] struct {
	*flag[T]
}

func Typed[T any](name string, opts ...Option) *TypedFlag[T] {
	f := newFlag[T](name)
	applyForFlag(f, opts...)
	return &TypedFlag[T]{f}
}
