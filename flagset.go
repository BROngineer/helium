package helium

import (
	"fmt"
	"github.com/brongineer/helium/internal/generic"
	"github.com/brongineer/helium/options"
	"os"
	"slices"
	"strings"
	"time"
)

type flag interface {
	Value() any
	Name() string
	Shorthand() string
	Separator() string
	IsRequired() bool
	IsShared() bool
	IsVisited() bool
	Parse(string) error
}

type FlagSet struct {
	flags []flag
}

func (fs *FlagSet) flagByName(name string) flag {
	idx := slices.IndexFunc(fs.flags, func(f flag) bool {
		return strings.EqualFold(f.Name(), name)
	})
	if idx == -1 {
		return nil
	}
	return fs.flags[idx]
}

func (fs *FlagSet) flagByShorthand(shorthand string) flag {
	idx := slices.IndexFunc(fs.flags, func(f flag) bool {
		return strings.EqualFold(f.Shorthand(), shorthand)
	})
	if idx == -1 {
		return nil
	}
	return fs.flags[idx]
}

func (fs *FlagSet) addFlag(f flag) {
	var fl flag
	fl = fs.flagByName(f.Name())
	if fl != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Flag \"%s\" already defined\n", f.Name())
		os.Exit(1)
	}
	switch {
	case f.Shorthand() != "":
		fl = fs.flagByShorthand(f.Shorthand())
		if fl != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Flag with shorthand \"%s\" already defined\n", f.Name())
			os.Exit(1)
		}
	default:
		set := make([]flag, len(fs.flags)+1)
		copy(set, fs.flags)
		set[len(fs.flags)] = f
		fs.flags = set
	}
}

// String creates a new string flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve string values from the command line.
func (fs *FlagSet) String(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[string](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Bool creates a new bool flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve bool values from the command line.
func (fs *FlagSet) Bool(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[bool](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Duration creates a new time.Duration flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve time.Duration values from the command line.
func (fs *FlagSet) Duration(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[time.Duration](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Int creates a new int flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve int values from the command line.
func (fs *FlagSet) Int(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[int](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Int8 creates a new int8 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve int8 values from the command line.
func (fs *FlagSet) Int8(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[int8](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Int16 creates a new int16 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve int16 values from the command line.
func (fs *FlagSet) Int16(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[int16](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Int32 creates a new int32 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve int32 values from the command line.
func (fs *FlagSet) Int32(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[int32](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Int64 creates a new int64 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve int64 values from the command line.
func (fs *FlagSet) Int64(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[int64](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Uint creates a new uint flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve uint values from the command line.
func (fs *FlagSet) Uint(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[uint](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Uint8 creates a new uint8 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve uint8 values from the command line.
func (fs *FlagSet) Uint8(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[uint8](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Uint16 creates a new uint16 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve uint16 values from the command line.
func (fs *FlagSet) Uint16(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[uint16](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Uint32 creates a new uint32 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve uint32 values from the command line.
func (fs *FlagSet) Uint32(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[uint32](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Uint64 creates a new uint64 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve uint64 values from the command line.
func (fs *FlagSet) Uint64(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[uint64](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Float32 creates a new float32 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve float32 values from the command line.
func (fs *FlagSet) Float32(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[float32](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Float64 creates a new float64 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve float64 values from the command line.
func (fs *FlagSet) Float64(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[float64](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// IntSlice creates a new []int flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []int values from the command line.
func (fs *FlagSet) IntSlice(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[[]int](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Int8Slice creates a new []int8 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []int8 values from the command line.
func (fs *FlagSet) Int8Slice(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[[]int8](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Int16Slice creates a new []int16 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []int16 values from the command line.
func (fs *FlagSet) Int16Slice(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[[]int16](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Int32Slice creates a new []int32 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []int32 values from the command line.
func (fs *FlagSet) Int32Slice(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[[]int32](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Int64Slice creates a new []int64 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []int64 values from the command line.
func (fs *FlagSet) Int64Slice(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[[]int64](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// UintSlice creates a new []uint flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []uint values from the command line.
func (fs *FlagSet) UintSlice(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[[]uint](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Uint8Slice creates a new []uint8 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []uint8 values from the command line.
func (fs *FlagSet) Uint8Slice(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[[]uint8](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Uint16Slice creates a new []uint16 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []uint16 values from the command line.
func (fs *FlagSet) Uint16Slice(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[[]uint16](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Uint32Slice creates a new []uint32 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []uint32 values from the command line.
func (fs *FlagSet) Uint32Slice(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[[]uint32](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Uint64Slice creates a new []uint64 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []uint64 values from the command line.
func (fs *FlagSet) Uint64Slice(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[[]uint64](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Float32Slice creates a new []float32 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []float32 values from the command line.
func (fs *FlagSet) Float32Slice(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[[]float32](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// Float64Slice creates a new []float64 flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []float64 values from the command line.
func (fs *FlagSet) Float64Slice(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[[]float64](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// StringSlice creates a new []string flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []string values from the command line.
func (fs *FlagSet) StringSlice(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[[]string](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// DurationSlice creates a new []time.Duration flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []time.Duration values from the command line.
func (fs *FlagSet) DurationSlice(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[[]time.Duration](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// BoolSlice creates a new []bool flag with the specified name and options.
// It returns a pointer to the created flag.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []bool values from the command line.
func (fs *FlagSet) BoolSlice(name string, opts ...options.FlagOption) {
	f := generic.NewFlag[[]bool](name)
	options.ApplyForFlag(f, opts...)
	fs.addFlag(f)
}

// flagValue returns a pointer to the value of the specified flag in the given FlagSet.
// If the flag does not exist, it returns an error with the message "unknown flag".
// If the value is nil, it returns an error with the message "value is nil".
// If the type of the value does not match the specified type, it returns an error with the message "wrong type".
// The function expects the generic.Allowed type to be one of the types defined in the flags.generic.Allowed interface.
func flagValue[T generic.Allowed](name string, fs *FlagSet) (*T, error) {
	var (
		value T
		ok    bool
	)

	f := fs.flagByName(name)
	if f == nil {
		return nil, fmt.Errorf("unknown flag '%s'", name)
	}
	value, ok = f.Value().(T)
	if !ok {
		return nil, fmt.Errorf("getter type '%T' does not match type of the flag '%s'", value, name)
	}
	return &value, nil
}

// derefOrDie dereferences a pointer and checks for errors. If the error is not nil,
// it prints the error message to stderr and exits the program with code 1. If the pointer is nil,
// it prints an error message to stderr and exits the program with code 1. It returns the dereferenced value.
func derefOrDie[V generic.Allowed](p *V, err error) V {
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

// ptrOrDie returns the pointer value `p` and exits the program if there is an error `err`.
// If `err` is not nil, an error message is printed to stderr and the program exits with code 1.
// The function is used to simplify error handling in flag retrieval functions.
func ptrOrDie[V generic.Allowed](p *V, err error) *V {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err.Error())
		os.Exit(1)
	}
	return p
}

// GetString returns the string value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetString(name string) string {
	return derefOrDie(flagValue[string](name, fs))
}

// GetStringPtr returns a pointer to a string value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetStringPtr(name string) *string {
	return ptrOrDie(flagValue[string](name, fs))
}

// GetBool returns the bool value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetBool(name string) bool {
	return derefOrDie(flagValue[bool](name, fs))
}

// GetBoolPtr returns a pointer to a bool value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetBoolPtr(name string) *bool {
	return ptrOrDie(flagValue[bool](name, fs))
}

// GetDuration returns the time.Duration value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetDuration(name string) time.Duration {
	return derefOrDie(flagValue[time.Duration](name, fs))
}

// GetDurationPtr returns a pointer to a time.Duration value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetDurationPtr(name string) *time.Duration {
	return ptrOrDie(flagValue[time.Duration](name, fs))
}

// GetInt returns the int value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt(name string) int {
	return derefOrDie(flagValue[int](name, fs))
}

// GetIntPtr returns a pointer to an int value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetIntPtr(name string) *int {
	return ptrOrDie(flagValue[int](name, fs))
}

// GetInt8 returns the int8 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt8(name string) int8 {
	return derefOrDie(flagValue[int8](name, fs))
}

// GetInt8Ptr returns a pointer to an int8 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt8Ptr(name string) *int8 {
	return ptrOrDie(flagValue[int8](name, fs))
}

// GetInt16 returns the int16 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt16(name string) int16 {
	return derefOrDie(flagValue[int16](name, fs))
}

// GetInt16Ptr returns a pointer to an int16 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt16Ptr(name string) *int16 {
	return ptrOrDie(flagValue[int16](name, fs))
}

// GetInt32 returns the int32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt32(name string) int32 {
	return derefOrDie(flagValue[int32](name, fs))
}

// GetInt32Ptr returns a pointer to an int32 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt32Ptr(name string) *int32 {
	return ptrOrDie(flagValue[int32](name, fs))
}

// GetInt64 returns the int64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt64(name string) int64 {
	return derefOrDie(flagValue[int64](name, fs))
}

// GetInt64Ptr returns a pointer to an int64 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt64Ptr(name string) *int64 {
	return ptrOrDie(flagValue[int64](name, fs))
}

// GetUint returns the uint value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint(name string) uint {
	return derefOrDie(flagValue[uint](name, fs))
}

// GetUintPtr returns a pointer to an uint value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUintPtr(name string) *uint {
	return ptrOrDie(flagValue[uint](name, fs))
}

// GetUint8 returns the uint8 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint8(name string) uint8 {
	return derefOrDie(flagValue[uint8](name, fs))
}

// GetUint8Ptr returns a pointer to an uint8 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint8Ptr(name string) *uint8 {
	return ptrOrDie(flagValue[uint8](name, fs))
}

// GetUint16 returns the uint16 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint16(name string) uint16 {
	return derefOrDie(flagValue[uint16](name, fs))
}

// GetUint16Ptr returns a pointer to an uint16 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint16Ptr(name string) *uint16 {
	return ptrOrDie(flagValue[uint16](name, fs))
}

// GetUint32 returns the uint32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint32(name string) uint32 {
	return derefOrDie(flagValue[uint32](name, fs))
}

// GetUint32Ptr returns a pointer to an uint32 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint32Ptr(name string) *uint32 {
	return ptrOrDie(flagValue[uint32](name, fs))
}

// GetUint64 returns the uint64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint64(name string) uint64 {
	return derefOrDie(flagValue[uint64](name, fs))
}

// GetUint64Ptr returns a pointer to an uint64 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint64Ptr(name string) *uint64 {
	return ptrOrDie(flagValue[uint64](name, fs))
}

// GetFloat32 returns the float32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetFloat32(name string) float32 {
	return derefOrDie(flagValue[float32](name, fs))
}

// GetFloat32Ptr returns a pointer to a float32 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetFloat32Ptr(name string) *float32 {
	return ptrOrDie(flagValue[float32](name, fs))
}

// GetFloat64 returns the float64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetFloat64(name string) float64 {
	return derefOrDie(flagValue[float64](name, fs))
}

// GetFloat64Ptr returns a pointer to a float64 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetFloat64Ptr(name string) *float64 {
	return ptrOrDie(flagValue[float64](name, fs))
}

// GetIntSlice returns the []int value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetIntSlice(name string) []int {
	return derefOrDie(flagValue[[]int](name, fs))
}

// GetIntSlicePtr returns a pointer to a []int value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetIntSlicePtr(name string) *[]int {
	return ptrOrDie(flagValue[[]int](name, fs))
}

// GetInt8Slice returns the []int8 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt8Slice(name string) []int8 {
	return derefOrDie(flagValue[[]int8](name, fs))
}

// GetInt8SlicePtr returns a pointer to a []int8 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt8SlicePtr(name string) *[]int8 {
	return ptrOrDie(flagValue[[]int8](name, fs))
}

// GetInt16Slice returns the []int16 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt16Slice(name string) []int16 {
	return derefOrDie(flagValue[[]int16](name, fs))
}

// GetInt16SlicePtr returns a pointer to a []int16 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt16SlicePtr(name string) *[]int16 {
	return ptrOrDie(flagValue[[]int16](name, fs))
}

// GetInt32Slice returns the []int32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt32Slice(name string) []int32 {
	return derefOrDie(flagValue[[]int32](name, fs))
}

// GetInt32SlicePtr returns a pointer to a []int32 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt32SlicePtr(name string) *[]int32 {
	return ptrOrDie(flagValue[[]int32](name, fs))
}

// GetInt64Slice returns the []int64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt64Slice(name string) []int64 {
	return derefOrDie(flagValue[[]int64](name, fs))
}

// GetInt64SlicePtr returns a pointer to a []int64 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt64SlicePtr(name string) *[]int64 {
	return ptrOrDie(flagValue[[]int64](name, fs))
}

// GetUintSlice returns the []uint value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUintSlice(name string) []uint {
	return derefOrDie(flagValue[[]uint](name, fs))
}

// GetUintSlicePtr returns a pointer to a []uint value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUintSlicePtr(name string) *[]uint {
	return ptrOrDie(flagValue[[]uint](name, fs))
}

// GetUint8Slice returns the []uint8 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint8Slice(name string) []uint8 {
	return derefOrDie(flagValue[[]uint8](name, fs))
}

// GetUint8SlicePtr returns a pointer to a []uint8 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint8SlicePtr(name string) *[]uint8 {
	return ptrOrDie(flagValue[[]uint8](name, fs))
}

// GetUint16Slice returns the []uint16 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint16Slice(name string) []uint16 {
	return derefOrDie(flagValue[[]uint16](name, fs))
}

// GetUint16SlicePtr returns a pointer to a []uint16 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint16SlicePtr(name string) *[]uint16 {
	return ptrOrDie(flagValue[[]uint16](name, fs))
}

// GetUint32Slice returns the []uint32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint32Slice(name string) []uint32 {
	return derefOrDie(flagValue[[]uint32](name, fs))
}

// GetUint32SlicePtr returns a pointer to a []uint32 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint32SlicePtr(name string) *[]uint32 {
	return ptrOrDie(flagValue[[]uint32](name, fs))
}

// GetUint64Slice returns the []uint64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint64Slice(name string) []uint64 {
	return derefOrDie(flagValue[[]uint64](name, fs))
}

// GetUint64SlicePtr returns a pointer to a []uint64 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint64SlicePtr(name string) *[]uint64 {
	return ptrOrDie(flagValue[[]uint64](name, fs))
}

// GetFloat32Slice returns the []float32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetFloat32Slice(name string) []float32 {
	return derefOrDie(flagValue[[]float32](name, fs))
}

// GetFloat32SlicePtr returns a pointer to a float64 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetFloat32SlicePtr(name string) *[]float32 {
	return ptrOrDie(flagValue[[]float32](name, fs))
}

// GetFloat64Slice returns the []float64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetFloat64Slice(name string) []float64 {
	return derefOrDie(flagValue[[]float64](name, fs))
}

// GetFloat64SlicePtr returns a pointer to a []float64 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetFloat64SlicePtr(name string) *[]float64 {
	return ptrOrDie(flagValue[[]float64](name, fs))
}

// GetStringSlice returns the []string value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetStringSlice(name string) []string {
	return derefOrDie(flagValue[[]string](name, fs))
}

// GetStringSlicePtr returns a pointer to a []string value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetStringSlicePtr(name string) *[]string {
	return ptrOrDie(flagValue[[]string](name, fs))
}

// GetDurationSlice returns the []time.Duration value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetDurationSlice(name string) []time.Duration {
	return derefOrDie(flagValue[[]time.Duration](name, fs))
}

// GetDurationSlicePtr returns a pointer to a []time.Duration value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetDurationSlicePtr(name string) *[]time.Duration {
	return ptrOrDie(flagValue[[]time.Duration](name, fs))
}

// GetBoolSlice returns the []bool value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetBoolSlice(name string) []bool {
	return derefOrDie(flagValue[[]bool](name, fs))
}

// GetBoolSlicePtr returns a pointer to a []bool value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetBoolSlicePtr(name string) *[]bool {
	return ptrOrDie(flagValue[[]bool](name, fs))
}
