package flag

import (
	"fmt"
	"os"

	"github.com/brongineer/helium/errors"
)

const defaultSliceSeparator = ","

type flagParser interface {
	SetFromEnv(bool)
	SetFromCmd(bool)
	SetSeparator(string)
	SetCurrentValue(any)
	IsSetFromEnv() bool
	IsSetFromCmd() bool
	Separator() string
	CurrentValue() any
	ParseCmd(string) (any, error)
	ParseEnv(string) (any, error)
}

type flag[T any] struct {
	name         string
	description  string
	shorthand    string
	shared       bool
	defaultValue *T
	value        *T
	separator    string
	parser       flagParser
	setFromEnv   bool
	setFromCmd   bool
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

func (f *flag[T]) Parser() flagParser {
	return f.parser
}

func (f *flag[T]) IsSetFromEnv() bool {
	return f.setFromEnv
}

func (f *flag[T]) IsSetFromCmd() bool {
	return f.setFromCmd
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

func (f *flag[T]) setParser(p flagParser) {
	f.parser = p
}

func (f *flag[T]) reflectStateToParser() {
	f.parser.SetFromEnv(f.IsSetFromEnv())
	f.parser.SetFromCmd(f.IsSetFromCmd())
	f.parser.SetSeparator(f.Separator())
	f.parser.SetCurrentValue(f.value)
}

func (f *flag[T]) parseInput(input string, parseFunc func(string) (any, error)) error {
	var (
		val    any
		parsed *T
		err    error
	)
	if f.parser == nil {
		return errors.ErrNoParserDefined
	}
	f.reflectStateToParser()
	val, err = parseFunc(input)
	if err != nil {
		return errors.ParseError(f.Name(), err)
	}
	parsed, err = typedValuePtr[T](val)
	if err != nil {
		return errors.ParseError(f.Name(), err)
	}
	f.value = parsed
	return nil
}

func (f *flag[T]) FromCommandLine(input string) error {
	if f.parser == nil {
		return errors.NoParserDefined(f.Name())
	}
	err := f.parseInput(input, f.parser.ParseCmd)
	if err == nil {
		f.setFromCmd = true
	}
	return err
}

func (f *flag[T]) FromEnvVariable(input string) error {
	if f.parser == nil {
		return errors.NoParserDefined(f.Name())
	}
	err := f.parseInput(input, f.parser.ParseEnv)
	if err == nil {
		f.setFromEnv = true
	}
	return err
}
