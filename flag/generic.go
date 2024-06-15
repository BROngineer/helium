package flag

import (
	"fmt"
	"os"

	"github.com/brongineer/helium/errors"
)

const defaultSliceSeparator = ","

type flag[T any] struct {
	name              string
	description       string
	shorthand         string
	shared            bool
	visited           bool
	defaultValue      *T
	value             *T
	separator         string
	commandLineParser func(string) (any, error)
	envVariableParser func(string) (any, error)
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

func (f *flag[T]) CommandLineParser() func(string) (any, error) {
	return f.commandLineParser
}

func (f *flag[T]) EnvVariableParser() func(string) (any, error) {
	return f.envVariableParser
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

func (f *flag[T]) setCommandLineParser(parser func(string) (any, error)) {
	f.commandLineParser = parser
}

func (f *flag[T]) setEnvVariableParser(parser func(string) (any, error)) {
	f.envVariableParser = parser
}

func customParse[T any](parser func(string) (any, error), input, flag string) (T, error) {
	var (
		val    T
		parsed any
		valPtr *T
		err    error
	)
	parsed, err = parser(input)
	if err != nil {
		return val, errors.ParseError(flag, err)
	}
	valPtr, err = value[T](parsed)
	if err != nil {
		return val, errors.ParseError(flag, err)
	}
	val = *valPtr
	return val, nil
}
