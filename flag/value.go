package flag

import (
	"fmt"
	"os"

	"github.com/brongineer/helium/errors"
)

// typedValuePtr returns a pointer to the value of the specified flag in the given FlagSet.
// If the flag does not exist, it returns an error with the message "unknown flag".
// If the value is nil, it returns an error with the message "value is nil".
// If the type of the value does not match the specified type, it returns an error with the message "wrong type".
func typedValuePtr[T any](v any) (*T, error) {
	var (
		valuePtr *T
		ok       bool
	)

	valuePtr, ok = v.(*T)
	if !ok {
		return nil, errors.TypeMismatch(valuePtr, v)
	}
	return valuePtr, nil
}

// DerefOrDie dereferences a pointer and checks for errors. If the error is not nil,
// it prints the error message to stderr and exits the program with code 1. If the pointer is nil,
// it prints an error message to stderr and exits the program with code 1. It returns the dereferenced value.
func DerefOrDie[T any](v any) T {
	p, err := typedValuePtr[T](v)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err.Error())
		os.Exit(1)
	}
	if p == nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: flag value is nil\n")
		os.Exit(1)
	}
	return *p
}

// PtrOrDie returns the pointer value `p` and exits the program if there is an error `err`.
// If `err` is not nil, an error message is printed to stderr and the program exits with code 1.
// The function is used to simplify error handling in flag retrieval functions.
func PtrOrDie[T any](v any) *T {
	p, err := typedValuePtr[T](v)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err.Error())
		os.Exit(1)
	}
	return p
}
