package flagset

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

type flagItem interface {
	Value() any
	Name() string
	Description() string
	Shorthand() string
	Separator() string
	IsShared() bool
	IsSetFromEnv() bool
	IsSetFromCmd() bool
	FromCommandLine(string) error
	FromEnvVariable(string) error
}

type FlagSet struct {
	flags        []flagItem
	envVarBinder *env.VarNameConstructor
}

// Parse iterates over the given args and calls the corresponding parse function
// for long flags and short flags. It returns an error if any parsing fails.
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
		default:
			i++
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// BindEnvVars binds environment variables to the corresponding flags in the FlagSet.
// It constructs a VarNameConstructor using the provided characters 'charOld' and 'charNew'
// and the environment options in the FlagSet. For each flag in the FlagSet, it retrieves
// the value of the environment variable using the VarNameConstructor and calls the
// FromEnvVariable method of the flag. If an error occurs during parsing, it is joined
// with the previous errors using the errors.Join function. The function returns the
// error encountered during parsing, if any.
func (fs *FlagSet) BindEnvVars() error {
	var (
		empty string
		err   error
	)
	for _, f := range fs.flags {
		val := os.Getenv(fs.envVarBinder.VarFromFlagName(f.Name()))
		if val == empty {
			continue
		}
		if e := f.FromEnvVariable(val); e != nil {
			err = errors.Join(e)
		}
	}
	return err
}

// parseLong trims the long flag name prefix from the argument and checks if
// the flag exists in the FlagSet. If the flag exists, it delegates to the parse
// method to parse the flag value. Returns the index of the next flag and any error
// encountered during parsing.
func (fs *FlagSet) parseLong(args []string, i int) (int, error) {
	trimmed := strings.TrimPrefix(args[i], longFlagNamePrefix)
	f := fs.flagByName(trimmed)
	if f == nil {
		return -1, ferrors.UnknownFlag(trimmed)
	}
	return fs.parse(f, i, args)
}

// parseShort trims the short flag name prefix from the argument and checks if
// the flag exists in the FlagSet. If the flag exists, it delegates to the parse
// method to parse the flag value. Returns the index of the next flag and any error
// encountered during parsing.
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

// parse iterates over the given args and calls the corresponding parse function
// for long flags and short flags. It returns the index of the next flag and any error
// encountered during parsing.
func (fs *FlagSet) parse(f flagItem, i int, args []string) (int, error) {
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

// parseStacked iterates over the given stacked flags and checks if each flag exists in the FlagSet.
// If the flag exists, it calls the FromCommandLine method of the flag with an empty value.
// Returns an error if any flag is not found or if an error occurs during parsing.
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

// nextFlagIndex finds the index of the next flag name in the given arguments
// starting from the specified index. It returns the index of the next flag name
// if found, otherwise it returns -1.
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

// flagByName searches for a flagItem in the FlagSet with the given name.
// It returns the found flagItem or nil if no match is found.
func (fs *FlagSet) flagByName(name string) flagItem {
	idx := slices.IndexFunc(fs.flags, func(f flagItem) bool {
		return strings.EqualFold(f.Name(), name)
	})
	if idx == -1 {
		return nil
	}
	return fs.flags[idx]
}

// flagByShorthand returns the flagItem with the given shorthand from the FlagSet.
// If the shorthand is not found, it returns nil.
func (fs *FlagSet) flagByShorthand(shorthand string) flagItem {
	idx := slices.IndexFunc(fs.flags, func(f flagItem) bool {
		return strings.EqualFold(f.Shorthand(), shorthand)
	})
	if idx == -1 {
		return nil
	}
	return fs.flags[idx]
}

// addFlag checks if a flag with the same name or shorthand already exists in the FlagSet.
// If a duplicate flag is found, it prints an error message to stderr and exits the program.
// Otherwise, it adds the flag to the `flags` slice of the FlagSet.
func (fs *FlagSet) addFlag(f flagItem) {
	if fl := fs.flagByName(f.Name()); fl != nil {
		_, _ = fmt.Fprintf(os.Stderr, "flag \"%s\" already defined\n", f.Name())
		os.Exit(1)
	}
	if f.Shorthand() != "" && fs.flagByShorthand(f.Shorthand()) != nil {
		_, _ = fmt.Fprintf(os.Stderr, "flag with shorthand \"%s\" already defined\n", f.Shorthand())
		os.Exit(1)
	}
	set := make([]flagItem, len(fs.flags)+1)
	copy(set, fs.flags)
	set[len(fs.flags)] = f
	fs.flags = set
}

// GetString returns the string value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetString(fs *FlagSet, name string) string {
	f := fs.flagByName(name)
	return flag.DerefOrDie[string](f.Value())
}

// GetStringPtr returns a pointer to a string value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetStringPtr(fs *FlagSet, name string) *string {
	f := fs.flagByName(name)
	return flag.PtrOrDie[string](f.Value())
}

// GetBool returns the bool value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetBool(fs *FlagSet, name string) bool {
	f := fs.flagByName(name)
	return flag.DerefOrDie[bool](f.Value())
}

// GetBoolPtr returns a pointer to a bool value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetBoolPtr(fs *FlagSet, name string) *bool {
	f := fs.flagByName(name)
	return flag.PtrOrDie[bool](f.Value())
}

// GetDuration returns the time.Duration value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetDuration(fs *FlagSet, name string) time.Duration {
	f := fs.flagByName(name)
	return flag.DerefOrDie[time.Duration](f.Value())
}

// GetDurationPtr returns a pointer to a time.Duration value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetDurationPtr(fs *FlagSet, name string) *time.Duration {
	f := fs.flagByName(name)
	return flag.PtrOrDie[time.Duration](f.Value())
}

// GetInt returns the int value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetInt(fs *FlagSet, name string) int {
	f := fs.flagByName(name)
	return flag.DerefOrDie[int](f.Value())
}

// GetIntPtr returns a pointer to an int value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetIntPtr(fs *FlagSet, name string) *int {
	f := fs.flagByName(name)
	return flag.PtrOrDie[int](f.Value())
}

// GetInt8 returns the int8 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetInt8(fs *FlagSet, name string) int8 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[int8](f.Value())
}

// GetInt8Ptr returns a pointer to an int8 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetInt8Ptr(fs *FlagSet, name string) *int8 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[int8](f.Value())
}

// GetInt16 returns the int16 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetInt16(fs *FlagSet, name string) int16 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[int16](f.Value())
}

// GetInt16Ptr returns a pointer to an int16 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetInt16Ptr(fs *FlagSet, name string) *int16 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[int16](f.Value())
}

// GetInt32 returns the int32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetInt32(fs *FlagSet, name string) int32 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[int32](f.Value())
}

// GetInt32Ptr returns a pointer to an int32 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetInt32Ptr(fs *FlagSet, name string) *int32 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[int32](f.Value())
}

// GetInt64 returns the int64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetInt64(fs *FlagSet, name string) int64 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[int64](f.Value())
}

// GetInt64Ptr returns a pointer to an int64 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetInt64Ptr(fs *FlagSet, name string) *int64 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[int64](f.Value())
}

// GetUint returns the uint value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetUint(fs *FlagSet, name string) uint {
	f := fs.flagByName(name)
	return flag.DerefOrDie[uint](f.Value())
}

// GetUintPtr returns a pointer to an uint value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetUintPtr(fs *FlagSet, name string) *uint {
	f := fs.flagByName(name)
	return flag.PtrOrDie[uint](f.Value())
}

// GetUint8 returns the uint8 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetUint8(fs *FlagSet, name string) uint8 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[uint8](f.Value())
}

// GetUint8Ptr returns a pointer to an uint8 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetUint8Ptr(fs *FlagSet, name string) *uint8 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[uint8](f.Value())
}

// GetUint16 returns the uint16 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetUint16(fs *FlagSet, name string) uint16 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[uint16](f.Value())
}

// GetUint16Ptr returns a pointer to an uint16 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetUint16Ptr(fs *FlagSet, name string) *uint16 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[uint16](f.Value())
}

// GetUint32 returns the uint32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetUint32(fs *FlagSet, name string) uint32 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[uint32](f.Value())
}

// GetUint32Ptr returns a pointer to an uint32 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetUint32Ptr(fs *FlagSet, name string) *uint32 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[uint32](f.Value())
}

// GetUint64 returns the uint64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetUint64(fs *FlagSet, name string) uint64 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[uint64](f.Value())
}

// GetUint64Ptr returns a pointer to an uint64 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetUint64Ptr(fs *FlagSet, name string) *uint64 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[uint64](f.Value())
}

// GetFloat32 returns the float32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetFloat32(fs *FlagSet, name string) float32 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[float32](f.Value())
}

// GetFloat32Ptr returns a pointer to a float32 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetFloat32Ptr(fs *FlagSet, name string) *float32 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[float32](f.Value())
}

// GetFloat64 returns the float64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetFloat64(fs *FlagSet, name string) float64 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[float64](f.Value())
}

// GetFloat64Ptr returns a pointer to a float64 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetFloat64Ptr(fs *FlagSet, name string) *float64 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[float64](f.Value())
}

// GetIntSlice returns the []int value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetIntSlice(fs *FlagSet, name string) []int {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]int](f.Value())
}

// GetIntSlicePtr returns a pointer to a []int value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetIntSlicePtr(fs *FlagSet, name string) *[]int {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]int](f.Value())
}

// GetInt8Slice returns the []int8 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetInt8Slice(fs *FlagSet, name string) []int8 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]int8](f.Value())
}

// GetInt8SlicePtr returns a pointer to a []int8 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetInt8SlicePtr(fs *FlagSet, name string) *[]int8 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]int8](f.Value())
}

// GetInt16Slice returns the []int16 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetInt16Slice(fs *FlagSet, name string) []int16 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]int16](f.Value())
}

// GetInt16SlicePtr returns a pointer to a []int16 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetInt16SlicePtr(fs *FlagSet, name string) *[]int16 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]int16](f.Value())
}

// GetInt32Slice returns the []int32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetInt32Slice(fs *FlagSet, name string) []int32 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]int32](f.Value())
}

// GetInt32SlicePtr returns a pointer to a []int32 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetInt32SlicePtr(fs *FlagSet, name string) *[]int32 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]int32](f.Value())
}

// GetInt64Slice returns the []int64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetInt64Slice(fs *FlagSet, name string) []int64 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]int64](f.Value())
}

// GetInt64SlicePtr returns a pointer to a []int64 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetInt64SlicePtr(fs *FlagSet, name string) *[]int64 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]int64](f.Value())
}

// GetUintSlice returns the []uint value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetUintSlice(fs *FlagSet, name string) []uint {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]uint](f.Value())
}

// GetUintSlicePtr returns a pointer to a []uint value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetUintSlicePtr(fs *FlagSet, name string) *[]uint {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]uint](f.Value())
}

// GetUint8Slice returns the []uint8 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetUint8Slice(fs *FlagSet, name string) []uint8 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]uint8](f.Value())
}

// GetUint8SlicePtr returns a pointer to a []uint8 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetUint8SlicePtr(fs *FlagSet, name string) *[]uint8 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]uint8](f.Value())
}

// GetUint16Slice returns the []uint16 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetUint16Slice(fs *FlagSet, name string) []uint16 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]uint16](f.Value())
}

// GetUint16SlicePtr returns a pointer to a []uint16 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetUint16SlicePtr(fs *FlagSet, name string) *[]uint16 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]uint16](f.Value())
}

// GetUint32Slice returns the []uint32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetUint32Slice(fs *FlagSet, name string) []uint32 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]uint32](f.Value())
}

// GetUint32SlicePtr returns a pointer to a []uint32 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetUint32SlicePtr(fs *FlagSet, name string) *[]uint32 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]uint32](f.Value())
}

// GetUint64Slice returns the []uint64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetUint64Slice(fs *FlagSet, name string) []uint64 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]uint64](f.Value())
}

// GetUint64SlicePtr returns a pointer to a []uint64 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetUint64SlicePtr(fs *FlagSet, name string) *[]uint64 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]uint64](f.Value())
}

// GetFloat32Slice returns the []float32 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetFloat32Slice(fs *FlagSet, name string) []float32 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]float32](f.Value())
}

// GetFloat32SlicePtr returns a pointer to a float64 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetFloat32SlicePtr(fs *FlagSet, name string) *[]float32 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]float32](f.Value())
}

// GetFloat64Slice returns the []float64 value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetFloat64Slice(fs *FlagSet, name string) []float64 {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]float64](f.Value())
}

// GetFloat64SlicePtr returns a pointer to a []float64 value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetFloat64SlicePtr(fs *FlagSet, name string) *[]float64 {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]float64](f.Value())
}

// GetStringSlice returns the []string value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetStringSlice(fs *FlagSet, name string) []string {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]string](f.Value())
}

// GetStringSlicePtr returns a pointer to a []string value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetStringSlicePtr(fs *FlagSet, name string) *[]string {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]string](f.Value())
}

// GetDurationSlice returns the []time.Duration value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetDurationSlice(fs *FlagSet, name string) []time.Duration {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]time.Duration](f.Value())
}

// GetDurationSlicePtr returns a pointer to a []time.Duration value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetDurationSlicePtr(fs *FlagSet, name string) *[]time.Duration {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]time.Duration](f.Value())
}

// GetBoolSlice returns the []bool value associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetBoolSlice(fs *FlagSet, name string) []bool {
	f := fs.flagByName(name)
	return flag.DerefOrDie[[]bool](f.Value())
}

// GetBoolSlicePtr returns a pointer to a []bool value associated with the given name from the FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetBoolSlicePtr(fs *FlagSet, name string) *[]bool {
	f := fs.flagByName(name)
	return flag.PtrOrDie[[]bool](f.Value())
}

// GetCounter returns the uint64 value, reflecting the counter, associated with the given name from the FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetCounter(fs *FlagSet, name string) int {
	f := fs.flagByName(name)
	return flag.DerefOrDie[int](f.Value())
}

// GetTypedFlag returns the value of defined type, associated with the given name from the given FlagSet.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value is nil
//   - flag value has a different type
func GetTypedFlag[T any](fs *FlagSet, name string) T {
	f := fs.flagByName(name)
	return flag.DerefOrDie[T](f.Value())
}

// GetTypedFlagPtr returns a pointer to the value of defined type,
// associated with the given name from the given FlagSet.
// If the flag value is not set, it returns nil.
// It will exit with code 1 if:
//   - flag does not exist
//   - flag value has a different type
func GetTypedFlagPtr[T any](fs *FlagSet, name string) *T {
	f := fs.flagByName(name)
	return flag.PtrOrDie[T](f.Value())
}
