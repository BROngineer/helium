package flag

import (
	"fmt"
	"os"
	"time"
)

const defaultSliceSeparator = ","

type intFlag interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type uintFlag interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type floatFlag interface {
	~float32 | ~float64
}

type intSliceFlag interface {
	~[]int | ~[]int8 | ~[]int16 | ~[]int32 | ~[]int64
}

type uintSliceFlag interface {
	~[]uint | ~[]uint8 | ~[]uint16 | ~[]uint32 | ~[]uint64
}

type floatSliceFlag interface {
	~[]float32 | ~[]float64
}

type sliceFlag interface {
	~[]string | ~[]bool | ~[]time.Duration | intSliceFlag | uintSliceFlag | floatSliceFlag
}

type allowed interface {
	~string | ~bool | intFlag | uintFlag | floatFlag | sliceFlag | time.Duration // | counter
}

type flag[T allowed] struct {
	name         string
	description  string
	shorthand    string
	required     bool
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

func (f *flag[T]) IsRequired() bool {
	return f.required
}

func (f *flag[T]) IsShared() bool {
	return f.shared
}

func (f *flag[T]) IsVisited() bool {
	return f.visited
}

func newFlag[T allowed](name string) *flag[T] {
	return &flag[T]{name: name}
}

func (f *flag[T]) setDescription(description string) {
	f.description = description
}

func (f *flag[T]) setShorthand(shorthand string) {
	f.shorthand = shorthand
}

func (f *flag[T]) setRequired() {
	f.required = true
}

func (f *flag[T]) setShared() {
	f.shared = true
}

func (f *flag[T]) setDefaultValue(value any) {
	var v T
	switch any(v).(type) {
	case T:
		v, _ = value.(T)
	default:
		_, _ = fmt.Fprintf(os.Stderr, "Error: wrong type for default value: got %T, wanted %T\n", value, v)
		os.Exit(1)
	}
	f.defaultValue = &v
}

func (f *flag[T]) setSeparator(separator string) {
	f.separator = separator
}
