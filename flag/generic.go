package flag

import (
	"fmt"
	"os"
)

const defaultSliceSeparator = ","

type flag[T any] struct {
	name         string
	description  string
	shorthand    string
	shared       bool
	visited      bool
	defaultValue *T
	value        *T
	separator    string
}

func (f *flag[T]) Value() any {
	if f.value == nil {
		return f.defaultValue
	}
	return f.value
}

func (f *flag[T]) Name() string {
	return f.name
}

func (f *flag[T]) Description() string {
	return f.description
}

func (f *flag[T]) Shorthand() string {
	return f.shorthand
}

func (f *flag[T]) Separator() string {
	if f.separator == "" {
		return defaultSliceSeparator
	}
	return f.separator
}

func (f *flag[T]) IsShared() bool {
	return f.shared
}

func (f *flag[T]) IsVisited() bool {
	return f.visited
}

func newFlag[T any](name string) *flag[T] {
	return &flag[T]{name: name}
}

func (f *flag[T]) setDescription(description string) {
	f.description = description
}

func (f *flag[T]) setShorthand(shorthand string) {
	f.shorthand = shorthand
}

func (f *flag[T]) setShared() {
	f.shared = true
}

func (f *flag[T]) setDefaultValue(value any) {
	var v T
	switch val := value.(type) {
	case T:
		v = val
	default:
		_, _ = fmt.Fprintf(os.Stderr, "Error: wrong type for default value: wanted %T, got %T\n", v, value)
		os.Exit(1)
	}
	f.defaultValue = &v
}

func (f *flag[T]) setSeparator(separator string) {
	f.separator = separator
}
