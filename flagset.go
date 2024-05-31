package helium

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/brongineer/helium/env"
	ferrors "github.com/brongineer/helium/errors"
	"github.com/brongineer/helium/flag"
)

const (
	longFlagNamePrefix  = "--"
	shortFlagNamePrefix = "-"
)

type flagPropertyGetter interface {
	Value() any
	Name() string
	Description() string
	Shorthand() string
	Separator() string
	IsShared() bool
	IsVisited() bool
	FromCommandLine(string) error
	FromEnvVariable(string) error
}

type FlagSet struct {
	flags      []flagPropertyGetter
	envOptions []env.Option
}

func (fs *FlagSet) Parse(args []string) error {
	var (
		i   int
		err error
	)

	for i = 0; i < len(args); {
		switch {
		case strings.HasPrefix(args[i], longFlagNamePrefix):
			i, err = fs.parseLong(args, i)
		case strings.HasPrefix(args[i], shortFlagNamePrefix):
			i, err = fs.parseShort(args, i)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (fs *FlagSet) BindEnvVars(charOld, charNew string) error {
	var (
		empty string
		err   error
	)
	c := env.Constructor(charOld, charNew, fs.envOptions...)
	for _, f := range fs.flags {
		val := os.Getenv(c.VarFromFlagName(f.Name()))
		if val == empty {
			continue
		}
		if e := f.FromEnvVariable(val); e != nil {
			err = errors.Join(e)
		}
	}
	return err
}

func (fs *FlagSet) parseLong(args []string, i int) (int, error) {
	trimmed := strings.TrimPrefix(args[i], longFlagNamePrefix)
	f := fs.flagByName(trimmed)
	if f == nil {
		return -1, ferrors.UnknownFlag(trimmed)
	}
	return fs.parse(f, i, args)
}

func (fs *FlagSet) parseShort(args []string, i int) (int, error) {
	trimmed := strings.TrimPrefix(args[i], shortFlagNamePrefix)
	if len(trimmed) > 1 {
		stacked := strings.Split(trimmed, "")
		if err := fs.parseStacked(stacked[:len(stacked)-1]); err != nil {
			return -1, err
		}
		trimmed = stacked[len(stacked)-1]
	}
	f := fs.flagByShorthand(trimmed)
	if f == nil {
		return -1, ferrors.UnknownShorthand(trimmed)
	}
	return fs.parse(f, i, args)
}

func (fs *FlagSet) parse(f flagPropertyGetter, i int, args []string) (int, error) {
	idx := nextFlagIndex(i+1, args)
	if idx > -1 {
		v := strings.Join(args[i+1:idx], f.Separator())
		if err := f.FromCommandLine(v); err != nil {
			return -1, err
		}
		return idx, nil
	}
	if i == len(args)-1 {
		if err := f.FromCommandLine(""); err != nil {
			return -1, err
		}
		return len(args), nil
	}
	v := strings.Join(args[i+1:], f.Separator())
	return len(args), f.FromCommandLine(v)
}

func (fs *FlagSet) parseStacked(stacked []string) error {
	for _, s := range stacked {
		f := fs.flagByShorthand(s)
		if f == nil {
			return ferrors.UnknownShorthand(s)
		}
		if err := f.FromCommandLine(""); err != nil {
			return err
		}
	}
	return nil
}

func nextFlagIndex(start int, args []string) int {
	for i := start; i < len(args); i++ {
		if strings.HasPrefix(args[i], longFlagNamePrefix) {
			return i
		}
		if strings.HasPrefix(args[i], shortFlagNamePrefix) {
			return i
		}
	}
	return -1
}

func (fs *FlagSet) flagByName(name string) flagPropertyGetter {
	idx := slices.IndexFunc(fs.flags, func(f flagPropertyGetter) bool {
		return strings.EqualFold(f.Name(), name)
	})
	if idx == -1 {
		return nil
	}
	return fs.flags[idx]
}

func (fs *FlagSet) flagByShorthand(shorthand string) flagPropertyGetter {
	idx := slices.IndexFunc(fs.flags, func(f flagPropertyGetter) bool {
		return strings.EqualFold(f.Shorthand(), shorthand)
	})
	if idx == -1 {
		return nil
	}
	return fs.flags[idx]
}

func (fs *FlagSet) addFlag(f flagPropertyGetter) {
	if fl := fs.flagByName(f.Name()); fl != nil {
		_, _ = fmt.Fprintf(os.Stderr, "flag \"%s\" already defined\n", f.Name())
		os.Exit(1)
	}
	if f.Shorthand() != "" && fs.flagByShorthand(f.Shorthand()) != nil {
		_, _ = fmt.Fprintf(os.Stderr, "flag with shorthand \"%s\" already defined\n", f.Shorthand())
		os.Exit(1)
	}
	set := make([]flagPropertyGetter, len(fs.flags)+1)
	copy(set, fs.flags)
	set[len(fs.flags)] = f
	fs.flags = set
}

// String creates a new string flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve string values from the command line.
func (fs *FlagSet) String(name string, opts ...flag.Option) {
	f := flag.NewString(name, opts...)
	fs.addFlag(f)
}

// Bool creates a new bool flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve bool values from the command line.
func (fs *FlagSet) Bool(name string, opts ...flag.Option) {
	f := flag.NewBool(name, opts...)
	fs.addFlag(f)
}

// Duration creates a new time.Duration flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve time.Duration values from the command line.
func (fs *FlagSet) Duration(name string, opts ...flag.Option) {
	f := flag.NewDuration(name, opts...)
	fs.addFlag(f)
}

// Int creates a new int flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve int values from the command line.
func (fs *FlagSet) Int(name string, opts ...flag.Option) {
	f := flag.NewInt(name, opts...)
	fs.addFlag(f)
}

// Int8 creates a new int8 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve int8 values from the command line.
func (fs *FlagSet) Int8(name string, opts ...flag.Option) {
	f := flag.NewInt8(name, opts...)
	fs.addFlag(f)
}

// Int16 creates a new int16 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve int16 values from the command line.
func (fs *FlagSet) Int16(name string, opts ...flag.Option) {
	f := flag.NewInt16(name, opts...)
	fs.addFlag(f)
}

// Int32 creates a new int32 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve int32 values from the command line.
func (fs *FlagSet) Int32(name string, opts ...flag.Option) {
	f := flag.NewInt32(name, opts...)
	fs.addFlag(f)
}

// Int64 creates a new int64 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve int64 values from the command line.
func (fs *FlagSet) Int64(name string, opts ...flag.Option) {
	f := flag.NewInt64(name, opts...)
	fs.addFlag(f)
}

// Uint creates a new uint flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve uint values from the command line.
func (fs *FlagSet) Uint(name string, opts ...flag.Option) {
	f := flag.NewUint(name, opts...)
	fs.addFlag(f)
}

// Uint8 creates a new uint8 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve uint8 values from the command line.
func (fs *FlagSet) Uint8(name string, opts ...flag.Option) {
	f := flag.NewUint8(name, opts...)
	fs.addFlag(f)
}

// Uint16 creates a new uint16 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve uint16 values from the command line.
func (fs *FlagSet) Uint16(name string, opts ...flag.Option) {
	f := flag.NewUint16(name, opts...)
	fs.addFlag(f)
}

// Uint32 creates a new uint32 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve uint32 values from the command line.
func (fs *FlagSet) Uint32(name string, opts ...flag.Option) {
	f := flag.NewUint32(name, opts...)
	fs.addFlag(f)
}

// Uint64 creates a new uint64 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve uint64 values from the command line.
func (fs *FlagSet) Uint64(name string, opts ...flag.Option) {
	f := flag.NewUint64(name, opts...)
	fs.addFlag(f)
}

// Float32 creates a new float32 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve float32 values from the command line.
func (fs *FlagSet) Float32(name string, opts ...flag.Option) {
	f := flag.NewFloat32(name, opts...)
	fs.addFlag(f)
}

// Float64 creates a new float64 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve float64 values from the command line.
func (fs *FlagSet) Float64(name string, opts ...flag.Option) {
	f := flag.NewFloat64(name, opts...)
	fs.addFlag(f)
}

// IntSlice creates a new []int flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []int values from the command line.
func (fs *FlagSet) IntSlice(name string, opts ...flag.Option) {
	f := flag.NewIntSlice(name, opts...)
	fs.addFlag(f)
}

// Int8Slice creates a new []int8 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []int8 values from the command line.
func (fs *FlagSet) Int8Slice(name string, opts ...flag.Option) {
	f := flag.NewInt8Slice(name, opts...)
	fs.addFlag(f)
}

// Int16Slice creates a new []int16 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []int16 values from the command line.
func (fs *FlagSet) Int16Slice(name string, opts ...flag.Option) {
	f := flag.NewInt16Slice(name, opts...)
	fs.addFlag(f)
}

// Int32Slice creates a new []int32 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []int32 values from the command line.
func (fs *FlagSet) Int32Slice(name string, opts ...flag.Option) {
	f := flag.NewInt32Slice(name, opts...)
	fs.addFlag(f)
}

// Int64Slice creates a new []int64 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []int64 values from the command line.
func (fs *FlagSet) Int64Slice(name string, opts ...flag.Option) {
	f := flag.NewInt64Slice(name, opts...)
	fs.addFlag(f)
}

// UintSlice creates a new []uint flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []uint values from the command line.
func (fs *FlagSet) UintSlice(name string, opts ...flag.Option) {
	f := flag.NewUintSlice(name, opts...)
	fs.addFlag(f)
}

// Uint8Slice creates a new []uint8 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []uint8 values from the command line.
func (fs *FlagSet) Uint8Slice(name string, opts ...flag.Option) {
	f := flag.NewUint8Slice(name, opts...)
	fs.addFlag(f)
}

// Uint16Slice creates a new []uint16 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []uint16 values from the command line.
func (fs *FlagSet) Uint16Slice(name string, opts ...flag.Option) {
	f := flag.NewUint16Slice(name, opts...)
	fs.addFlag(f)
}

// Uint32Slice creates a new []uint32 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []uint32 values from the command line.
func (fs *FlagSet) Uint32Slice(name string, opts ...flag.Option) {
	f := flag.NewUint32Slice(name, opts...)
	fs.addFlag(f)
}

// Uint64Slice creates a new []uint64 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []uint64 values from the command line.
func (fs *FlagSet) Uint64Slice(name string, opts ...flag.Option) {
	f := flag.NewUint64Slice(name, opts...)
	fs.addFlag(f)
}

// Float32Slice creates a new []float32 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []float32 values from the command line.
func (fs *FlagSet) Float32Slice(name string, opts ...flag.Option) {
	f := flag.NewFloat32Slice(name, opts...)
	fs.addFlag(f)
}

// Float64Slice creates a new []float64 flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []float64 values from the command line.
func (fs *FlagSet) Float64Slice(name string, opts ...flag.Option) {
	f := flag.NewFloat64Slice(name, opts...)
	fs.addFlag(f)
}

// StringSlice creates a new []string flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []string values from the command line.
func (fs *FlagSet) StringSlice(name string, opts ...flag.Option) {
	f := flag.NewStringSlice(name, opts...)
	fs.addFlag(f)
}

// DurationSlice creates a new []time.Duration flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []time.Duration values from the command line.
func (fs *FlagSet) DurationSlice(name string, opts ...flag.Option) {
	f := flag.NewDurationSlice(name, opts...)
	fs.addFlag(f)
}

// BoolSlice creates a new []bool flag with the specified name and options.
// It appends the created flag to the list of flags which belongs to FlagSet.
//
// The created flag implements the Flag interface and can be used to set
// and retrieve []bool values from the command line.
func (fs *FlagSet) BoolSlice(name string, opts ...flag.Option) {
	f := flag.NewBoolSlice(name, opts...)
	fs.addFlag(f)
}

func (fs *FlagSet) Counter(name string, opts ...flag.Option) {
	f := flag.NewCounter(name, opts...)
	fs.addFlag(f)
}

// GetString returns the string value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetString(name string) string {
	f := fs.flagByName(name)
	return flag.DerefOrDie[string](f.Value())
}

// GetStringPtr returns a pointer to a string value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetStringPtr(name string) *string {
	f := fs.flagByName(name)
	return flag.PtrOrDie[string](f.Value())
}

// GetBool returns the bool value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetBool(name string) bool {
	f := fs.flagByName(name)
	return flag.DerefOrDie[bool](f.Value())
}

// GetBoolPtr returns a pointer to a bool value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetBoolPtr(name string) *bool {
	f := fs.flagByName(name)
	return flag.PtrOrDie[bool](f.Value())
}

// GetDuration returns the time.Duration value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetDuration(name string) time.Duration {
	f := fs.flagByName(name)
	return flag.DerefOrDie[time.Duration](f.Value())
}

// GetDurationPtr returns a pointer to a time.Duration value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetDurationPtr(name string) *time.Duration {
	f := fs.flagByName(name)
	return flag.PtrOrDie[time.Duration](f.Value())
}

// GetInt returns the int value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt(name string) int {
	f := fs.flagByName(name)
	return flag.DerefOrDie[int](f.Value())
}

// GetIntPtr returns a pointer to an int value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetIntPtr(name string) *int {
	f := fs.flagByName(name)
	return flag.PtrOrDie[int](f.Value())
}

// GetInt8 returns the int8 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt8(name string) int8 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[int8](f.Value())
}

// GetInt8Ptr returns a pointer to an int8 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt8Ptr(name string) *int8 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[int8](f.Value())
}

// GetInt16 returns the int16 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt16(name string) int16 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[int16](f.Value())
}

// GetInt16Ptr returns a pointer to an int16 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt16Ptr(name string) *int16 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[int16](f.Value())
}

// GetInt32 returns the int32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt32(name string) int32 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[int32](f.Value())
}

// GetInt32Ptr returns a pointer to an int32 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt32Ptr(name string) *int32 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[int32](f.Value())
}

// GetInt64 returns the int64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt64(name string) int64 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[int64](f.Value())
}

// GetInt64Ptr returns a pointer to an int64 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt64Ptr(name string) *int64 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[int64](f.Value())
}

// GetUint returns the uint value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint(name string) uint {
	f := fs.flagByName(name)
	return flag.DerefOrDie[uint](f.Value())
}

// GetUintPtr returns a pointer to an uint value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUintPtr(name string) *uint {
	f := fs.flagByName(name)
	return flag.PtrOrDie[uint](f.Value())
}

// GetUint8 returns the uint8 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint8(name string) uint8 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[uint8](f.Value())
}

// GetUint8Ptr returns a pointer to an uint8 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint8Ptr(name string) *uint8 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[uint8](f.Value())
}

// GetUint16 returns the uint16 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint16(name string) uint16 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[uint16](f.Value())
}

// GetUint16Ptr returns a pointer to an uint16 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint16Ptr(name string) *uint16 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[uint16](f.Value())
}

// GetUint32 returns the uint32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint32(name string) uint32 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[uint32](f.Value())
}

// GetUint32Ptr returns a pointer to an uint32 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint32Ptr(name string) *uint32 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[uint32](f.Value())
}

// GetUint64 returns the uint64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint64(name string) uint64 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[uint64](f.Value())
}

// GetUint64Ptr returns a pointer to an uint64 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint64Ptr(name string) *uint64 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[uint64](f.Value())
}

// GetFloat32 returns the float32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetFloat32(name string) float32 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[float32](f.Value())
}

// GetFloat32Ptr returns a pointer to a float32 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetFloat32Ptr(name string) *float32 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[float32](f.Value())
}

// GetFloat64 returns the float64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetFloat64(name string) float64 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[float64](f.Value())
}

// GetFloat64Ptr returns a pointer to a float64 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetFloat64Ptr(name string) *float64 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[float64](f.Value())
}

// GetIntSlice returns the []int value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetIntSlice(name string) []int {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]int](f.Value())
}

// GetIntSlicePtr returns a pointer to a []int value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetIntSlicePtr(name string) *[]int {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]int](f.Value())
}

// GetInt8Slice returns the []int8 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt8Slice(name string) []int8 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]int8](f.Value())
}

// GetInt8SlicePtr returns a pointer to a []int8 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt8SlicePtr(name string) *[]int8 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]int8](f.Value())
}

// GetInt16Slice returns the []int16 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt16Slice(name string) []int16 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]int16](f.Value())
}

// GetInt16SlicePtr returns a pointer to a []int16 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt16SlicePtr(name string) *[]int16 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]int16](f.Value())
}

// GetInt32Slice returns the []int32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt32Slice(name string) []int32 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]int32](f.Value())
}

// GetInt32SlicePtr returns a pointer to a []int32 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt32SlicePtr(name string) *[]int32 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]int32](f.Value())
}

// GetInt64Slice returns the []int64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetInt64Slice(name string) []int64 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]int64](f.Value())
}

// GetInt64SlicePtr returns a pointer to a []int64 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetInt64SlicePtr(name string) *[]int64 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]int64](f.Value())
}

// GetUintSlice returns the []uint value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUintSlice(name string) []uint {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]uint](f.Value())
}

// GetUintSlicePtr returns a pointer to a []uint value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUintSlicePtr(name string) *[]uint {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]uint](f.Value())
}

// GetUint8Slice returns the []uint8 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint8Slice(name string) []uint8 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]uint8](f.Value())
}

// GetUint8SlicePtr returns a pointer to a []uint8 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint8SlicePtr(name string) *[]uint8 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]uint8](f.Value())
}

// GetUint16Slice returns the []uint16 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint16Slice(name string) []uint16 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]uint16](f.Value())
}

// GetUint16SlicePtr returns a pointer to a []uint16 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint16SlicePtr(name string) *[]uint16 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]uint16](f.Value())
}

// GetUint32Slice returns the []uint32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint32Slice(name string) []uint32 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]uint32](f.Value())
}

// GetUint32SlicePtr returns a pointer to a []uint32 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint32SlicePtr(name string) *[]uint32 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]uint32](f.Value())
}

// GetUint64Slice returns the []uint64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetUint64Slice(name string) []uint64 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]uint64](f.Value())
}

// GetUint64SlicePtr returns a pointer to a []uint64 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetUint64SlicePtr(name string) *[]uint64 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]uint64](f.Value())
}

// GetFloat32Slice returns the []float32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetFloat32Slice(name string) []float32 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]float32](f.Value())
}

// GetFloat32SlicePtr returns a pointer to a float64 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetFloat32SlicePtr(name string) *[]float32 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]float32](f.Value())
}

// GetFloat64Slice returns the []float64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetFloat64Slice(name string) []float64 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]float64](f.Value())
}

// GetFloat64SlicePtr returns a pointer to a []float64 value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetFloat64SlicePtr(name string) *[]float64 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]float64](f.Value())
}

// GetStringSlice returns the []string value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetStringSlice(name string) []string {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]string](f.Value())
}

// GetStringSlicePtr returns a pointer to a []string value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetStringSlicePtr(name string) *[]string {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]string](f.Value())
}

// GetDurationSlice returns the []time.Duration value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetDurationSlice(name string) []time.Duration {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]time.Duration](f.Value())
}

// GetDurationSlicePtr returns a pointer to a []time.Duration value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetDurationSlicePtr(name string) *[]time.Duration {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]time.Duration](f.Value())
}

// GetBoolSlice returns the []bool value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetBoolSlice(name string) []bool {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]bool](f.Value())
}

// GetBoolSlicePtr returns a pointer to a []bool value associated with the given name from the FlagSet.
// If flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has different type
func (fs *FlagSet) GetBoolSlicePtr(name string) *[]bool {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]bool](f.Value())
}

// GetCounter returns the uint64 value, reflecting the counter, associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has different type
func (fs *FlagSet) GetCounter(name string) int {
	f := fs.flagByName(name)
	return flag.DerefOrDie[int](f.Value())
}
