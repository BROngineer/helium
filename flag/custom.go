package flag

import "github.com/brongineer/helium/errors"

type Custom[T any] struct {
	*flag[T]
}

func (f *Custom[T]) FromCommandLine(input string) error {
	if f.CommandLineParser() == nil {
		return errors.NoParserDefined(f.Name())
	}

	val, err := customParse[T](f.CommandLineParser(), input, f.Name())
	if err != nil {
		return err
	}
	f.value = &val
	f.visited = true
	return nil
}

func (f *Custom[T]) FromEnvVariable(input string) error {
	if f.EnvVariableParser() == nil {
		return errors.NoParserDefined(f.Name())
	}

	val, err := customParse[T](f.EnvVariableParser(), input, f.Name())
	if err != nil {
		return err
	}
	f.value = &val
	f.visited = true
	return nil
}

func NewCustom[T any](name string, opts ...Option) *Custom[T] {
	f := newFlag[T](name)
	applyForFlag(f, opts...)
	return &Custom[T]{f}
}
