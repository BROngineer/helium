package errors

import (
	"errors"
	"fmt"
)

const (
	unknownFlagMessage      = "unknown flag"
	unknownShorthandMessage = "unknown shorthand"
	alreadyParsedMessage    = "flag already parsed"
	noValueProvidedMessage  = "no value provided for flag"
	parseErrorMessage       = "failed to parse flag"
	typeMismatchMessage     = "value type mismatch"
	noParserDefinedMessage  = "no input parser defined for flag"
)

var (
	ErrUnknownFlag      = errors.New(unknownFlagMessage)
	ErrUnknownShorthand = errors.New(unknownShorthandMessage)
	ErrFlagVisited      = errors.New(alreadyParsedMessage)
	ErrNoValueProvided  = errors.New(noValueProvidedMessage)
	ErrParseFailed      = errors.New(parseErrorMessage)
	ErrTypeMismatch     = errors.New(typeMismatchMessage)
	ErrNoParserDefined  = errors.New(noParserDefinedMessage)
)

func UnknownFlag(flagName string) error {
	return fmt.Errorf("%s: %w", flagName, ErrUnknownFlag)
}

func UnknownShorthand(shorthandName string) error {
	return fmt.Errorf("%s: %w", shorthandName, ErrUnknownShorthand)
}

func FlagVisited(flagName string) error {
	return fmt.Errorf("%s: %w", flagName, ErrFlagVisited)
}

func NoValueProvided(flagName string) error {
	return fmt.Errorf("%s: %w", flagName, ErrNoValueProvided)
}

func ParseError(flagName string, err error) error {
	return errors.Join(fmt.Errorf("%s: %w", flagName, ErrParseFailed), err)
}

func TypeMismatch(actual, expected any) error {
	return fmt.Errorf("expected %T, got %T: %w", expected, actual, ErrTypeMismatch)
}

func NoParserDefined(flagName string) error {
	return fmt.Errorf("%s: %w", flagName, ErrNoParserDefined)
}
