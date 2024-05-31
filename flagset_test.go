package helium

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/brongineer/helium/env"
	ferrors "github.com/brongineer/helium/errors"
	"github.com/brongineer/helium/flag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertParsedValue(t *testing.T, fs *FlagSet, r result) {
	switch r.flagType {
	case "string":
		val := fs.GetString(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetStringPtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "stringSlice":
		val := fs.GetStringSlice(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetStringSlicePtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "bool":
		val := fs.GetBool(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetBoolPtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "boolSlice":
		val := fs.GetBoolSlice(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetBoolSlicePtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "duration":
		val := fs.GetDuration(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetDurationPtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "durationSlice":
		val := fs.GetDurationSlice(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetDurationSlicePtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int":
		val := fs.GetInt(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetIntPtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "intSlice":
		val := fs.GetIntSlice(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetIntSlicePtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int8":
		val := fs.GetInt8(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetInt8Ptr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int8slice":
		val := fs.GetInt8Slice(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetInt8SlicePtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int16":
		val := fs.GetInt16(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetInt16Ptr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int16slice":
		val := fs.GetInt16Slice(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetInt16SlicePtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int32":
		val := fs.GetInt32(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetInt32Ptr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int32slice":
		val := fs.GetInt32Slice(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetInt32SlicePtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int64":
		val := fs.GetInt64(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetInt64Ptr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int64slice":
		val := fs.GetInt64Slice(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetInt64SlicePtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint":
		val := fs.GetUint(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetUintPtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uintSlice":
		val := fs.GetUintSlice(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetUintSlicePtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint8":
		val := fs.GetUint8(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetUint8Ptr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint8slice":
		val := fs.GetUint8Slice(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetUint8SlicePtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint16":
		val := fs.GetUint16(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetUint16Ptr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint16slice":
		val := fs.GetUint16Slice(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetUint16SlicePtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint32":
		val := fs.GetUint32(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetUint32Ptr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint32slice":
		val := fs.GetUint32Slice(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetUint32SlicePtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint64":
		val := fs.GetUint64(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetUint64Ptr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint64slice":
		val := fs.GetUint64Slice(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetUint64SlicePtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "float32":
		val := fs.GetFloat32(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetFloat32Ptr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "float32slice":
		val := fs.GetFloat32Slice(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetFloat32SlicePtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "float64":
		val := fs.GetFloat64(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetFloat64Ptr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "float64slice":
		val := fs.GetFloat64Slice(r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := fs.GetFloat64SlicePtr(r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "counter":
		val := fs.GetCounter(r.flagName)
		assert.Equal(t, r.flagValue, val)
	default:
		t.Fatalf("unknown flag type: %s", r.flagType)
	}
}

type result struct {
	flagName  string
	flagValue any
	flagType  string
}

type expected struct {
	parsed      []result
	err         bool
	expectedErr error
}

type flagSetTest struct {
	name     string
	flagSet  func() *FlagSet
	expected expected
	input    []string
}

func TestFlagSet_Parse(t *testing.T) {
	t.Parallel()
	tests := []flagSetTest{
		{
			name: "parse",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.String("sample-string")
				fs.Bool("sample-bool")
				return fs
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-string", flagValue: "foo", flagType: "string"},
					{flagName: "sample-bool", flagValue: true, flagType: "bool"},
				},
				err: false,
			},
			input: []string{"--sample-string", "foo", "--sample-bool"},
		},
		{
			name: "parse error",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.String("sample-string")
				fs.Bool("sample-bool")
				return fs
			},
			expected: expected{
				parsed:      []result{},
				err:         true,
				expectedErr: ferrors.ErrNoValueProvided,
			},
			input: []string{"--sample-string", "--sample-bool"},
		},
		{
			name: "parse shorthand",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.String("sample-string")
				fs.Bool("sample-bool")
				fs.Counter("sample-counter", flag.Shorthand("c"))
				return fs
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-string", flagValue: "foo", flagType: "string"},
					{flagName: "sample-bool", flagValue: true, flagType: "bool"},
					{flagName: "sample-counter", flagValue: 1, flagType: "counter"},
				},
				err: false,
			},
			input: []string{"--sample-string", "foo", "-c", "--sample-bool"},
		},
		{
			name: "parse stacked",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.Bool("sample-bool-one", flag.Shorthand("o"))
				fs.Bool("sample-bool-two", flag.Shorthand("t"))
				fs.Counter("sample-counter", flag.Shorthand("c"))
				return fs
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-bool-one", flagValue: true, flagType: "bool"},
					{flagName: "sample-bool-two", flagValue: true, flagType: "bool"},
					{flagName: "sample-counter", flagValue: 3, flagType: "counter"},
				},
				err: false,
			},
			input: []string{"-octcc"},
		},
		{
			name: "parse stacked with value",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.Bool("sample-bool", flag.Shorthand("b"))
				fs.Counter("sample-counter", flag.Shorthand("c"))
				fs.Int("sample-int", flag.Shorthand("i"))
				return fs
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-bool", flagValue: true, flagType: "bool"},
					{flagName: "sample-counter", flagValue: 2, flagType: "counter"},
					{flagName: "sample-int", flagValue: 10, flagType: "int"},
				},
				err: false,
			},
			input: []string{"-bcci", "10"},
		},
		{
			name: "parse error duplicate",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.String("sample-string")
				fs.Bool("sample-bool")
				return fs
			},
			expected: expected{
				parsed:      []result{},
				err:         true,
				expectedErr: ferrors.ErrFlagVisited,
			},
			input: []string{"--sample-string", "foo", "--sample-bool", "--sample-string", "bar"},
		},
		{
			name: "parse slice duplicate",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.String("sample-string")
				fs.BoolSlice("sample-bool")
				return fs
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-bool", flagValue: []bool{true, false}, flagType: "boolSlice"},
					{flagName: "sample-string", flagValue: "foo", flagType: "string"},
				},
				err: false,
			},
			input: []string{"--sample-bool", "true", "--sample-string", "foo", "--sample-bool", "false"},
		},
		{
			name: "parse unknown flag error",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.String("sample-string")
				return fs
			},
			expected: expected{
				parsed:      []result{},
				err:         true,
				expectedErr: ferrors.ErrUnknownFlag,
			},
			input: []string{"--sample"},
		},
		{
			name: "parse unknown shorthand error",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.String("sample-string", flag.Shorthand("s"))
				return fs
			},
			expected: expected{
				parsed:      []result{},
				err:         true,
				expectedErr: ferrors.ErrUnknownShorthand,
			},
			input: []string{"-v"},
		},
		{
			name: "parse unknown shorthand stacked error",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.String("sample-string", flag.Shorthand("s"))
				return fs
			},
			expected: expected{
				parsed:      []result{},
				err:         true,
				expectedErr: ferrors.ErrUnknownShorthand,
			},
			input: []string{"-vvs"},
		},
		{
			name: "parse stacked no value error",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.String("sample-string", flag.Shorthand("s"))
				return fs
			},
			expected: expected{
				parsed:      []result{},
				err:         true,
				expectedErr: ferrors.ErrNoValueProvided,
			},
			input: []string{"-sss"},
		},
		{
			name: "parse no value error",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.String("sample-string", flag.Shorthand("s"))
				return fs
			},
			expected: expected{
				parsed:      []result{},
				err:         true,
				expectedErr: ferrors.ErrNoValueProvided,
			},
			input: []string{"--sample-string"},
		},
		{
			name: "parse integers",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.Int("sample-int", flag.Shorthand("a"))
				fs.Int8("sample-int8", flag.Shorthand("b"))
				fs.Int16("sample-int16", flag.Shorthand("c"))
				fs.Int32("sample-int32", flag.Shorthand("d"))
				fs.Int64("sample-int64", flag.Shorthand("e"))
				return fs
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-int", flagValue: 1, flagType: "int"},
					{flagName: "sample-int8", flagValue: int8(2), flagType: "int8"},
					{flagName: "sample-int16", flagValue: int16(3), flagType: "int16"},
					{flagName: "sample-int32", flagValue: int32(4), flagType: "int32"},
					{flagName: "sample-int64", flagValue: int64(5), flagType: "int64"},
				},
				err: false,
			},
			input: []string{"--sample-int", "1", "--sample-int8", "2", "--sample-int16", "3", "--sample-int32", "4", "--sample-int64", "5"},
		},
		{
			name: "parse unsigned integers",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.Uint("sample-uint", flag.Shorthand("a"))
				fs.Uint8("sample-uint8", flag.Shorthand("b"))
				fs.Uint16("sample-uint16", flag.Shorthand("c"))
				fs.Uint32("sample-uint32", flag.Shorthand("d"))
				fs.Uint64("sample-uint64", flag.Shorthand("e"))
				return fs
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-uint", flagValue: uint(1), flagType: "uint"},
					{flagName: "sample-uint8", flagValue: uint8(2), flagType: "uint8"},
					{flagName: "sample-uint16", flagValue: uint16(3), flagType: "uint16"},
					{flagName: "sample-uint32", flagValue: uint32(4), flagType: "uint32"},
					{flagName: "sample-uint64", flagValue: uint64(5), flagType: "uint64"},
				},
				err: false,
			},
			input: []string{"--sample-uint", "1", "--sample-uint8", "2", "--sample-uint16", "3", "--sample-uint32", "4", "--sample-uint64", "5"},
		},
		{
			name: "parse integer slices",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.IntSlice("sample-int", flag.Shorthand("a"))
				fs.Int8Slice("sample-int8", flag.Shorthand("b"))
				fs.Int16Slice("sample-int16", flag.Shorthand("c"))
				fs.Int32Slice("sample-int32", flag.Shorthand("d"), flag.Separator("|"))
				fs.Int64Slice("sample-int64", flag.Shorthand("e"), flag.Separator(";"))
				return fs
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-int", flagValue: []int{1, 2}, flagType: "intSlice"},
					{flagName: "sample-int8", flagValue: []int8{3, 4}, flagType: "int8slice"},
					{flagName: "sample-int16", flagValue: []int16{5, 6}, flagType: "int16slice"},
					{flagName: "sample-int32", flagValue: []int32{7, 8}, flagType: "int32slice"},
					{flagName: "sample-int64", flagValue: []int64{9, 10}, flagType: "int64slice"},
				},
				err: false,
			},
			input: []string{
				"--sample-int", "1,2",
				"--sample-int8", "3,4",
				"--sample-int16", "5",
				"-c", "6",
				"--sample-int32", "7|8",
				"-e", "9;10",
			},
		},
		{
			name: "parse unsigned integers",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.UintSlice("sample-uint", flag.Shorthand("a"))
				fs.Uint8Slice("sample-uint8", flag.Shorthand("b"))
				fs.Uint16Slice("sample-uint16", flag.Shorthand("c"))
				fs.Uint32Slice("sample-uint32", flag.Shorthand("d"))
				fs.Uint64Slice("sample-uint64", flag.Shorthand("e"))
				return fs
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-uint", flagValue: []uint{1}, flagType: "uintSlice"},
					{flagName: "sample-uint8", flagValue: []uint8{2}, flagType: "uint8slice"},
					{flagName: "sample-uint16", flagValue: []uint16{3}, flagType: "uint16slice"},
					{flagName: "sample-uint32", flagValue: []uint32{4}, flagType: "uint32slice"},
					{flagName: "sample-uint64", flagValue: []uint64{5}, flagType: "uint64slice"},
				},
				err: false,
			},
			input: []string{"--sample-uint", "1", "--sample-uint8", "2", "--sample-uint16", "3", "--sample-uint32", "4", "--sample-uint64", "5"},
		},
		{
			name: "parse floats",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.Float32("sample-float32", flag.Shorthand("a"))
				fs.Float64("sample-float64", flag.Shorthand("b"))
				fs.Float32Slice("sample-float32-slice", flag.Shorthand("c"))
				fs.Float64Slice("sample-float64-slice", flag.Shorthand("d"))
				return fs
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-float32", flagValue: float32(3.14), flagType: "float32"},
					{flagName: "sample-float64", flagValue: 3.14, flagType: "float64"},
					{flagName: "sample-float32-slice", flagValue: []float32{2.7, 3.14}, flagType: "float32slice"},
					{flagName: "sample-float64-slice", flagValue: []float64{2.7, 3.14}, flagType: "float64slice"},
				},
				err: false,
			},
			input: []string{
				"--sample-float32", "3.14",
				"-b", "3.14",
				"-c", "2.7,3.14",
				"--sample-float64-slice", "2.7",
				"--sample-float64-slice", "3.14",
			},
		},
		{
			name: "parse string slice and duration",
			flagSet: func() *FlagSet {
				fs := &FlagSet{}
				fs.Duration("sample-duration", flag.Shorthand("a"))
				fs.StringSlice("sample-string", flag.Shorthand("b"), flag.Separator(";"))
				fs.DurationSlice("sample-duration-slice", flag.Shorthand("c"), flag.Separator(" "))
				return fs
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-duration", flagValue: time.Minute * 15, flagType: "duration"},
					{flagName: "sample-string", flagValue: []string{"foo", "bar"}, flagType: "stringSlice"},
					{flagName: "sample-duration-slice", flagValue: []time.Duration{time.Second, time.Hour * 2}, flagType: "durationSlice"},
				},
				err: false,
			},
			input: []string{
				"--sample-duration", "15m",
				"-b", "foo;bar",
				"--sample-duration-slice", "1s", "2h",
			},
		},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fs := tt.flagSet()
			err := fs.Parse(tt.input)
			if tt.expected.err {
				require.Error(t, err)
				assert.True(t, errors.Is(err, tt.expected.expectedErr))
				return
			}
			assert.NoError(t, err)
			for _, r := range tt.expected.parsed {
				require.NotNil(t, fs.flagByName(r.flagName).Value())
				assertParsedValue(t, fs, r)
			}
		})
	}
}

type envTestCase struct {
	name      string
	flagSet   func() *FlagSet
	variables map[string]string
	expected  expected
}

func TestFlagSet_BindEnvVars(t *testing.T) {
	t.Parallel()

	tests := []envTestCase{
		{
			name: "env vars success",
			flagSet: func() *FlagSet {
				fs := &FlagSet{envOptions: []env.Option{env.Prefix("test1"), env.Capitalized()}}
				fs.String("sample-string")
				fs.Bool("sample-bool")
				fs.Duration("sample-duration")
				return fs
			},
			variables: map[string]string{
				"TEST1_SAMPLE_STRING":   "foo",
				"TEST1_SAMPLE_BOOL":     "true",
				"TEST1_SAMPLE_DURATION": "15m",
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-duration", flagValue: time.Minute * 15, flagType: "duration"},
					{flagName: "sample-string", flagValue: "foo", flagType: "string"},
					{flagName: "sample-bool", flagValue: true, flagType: "bool"},
				},
				err: false,
			},
		},
		{
			name: "env vars failed",
			flagSet: func() *FlagSet {
				fs := &FlagSet{envOptions: []env.Option{env.Prefix("test2"), env.Capitalized()}}
				fs.String("sample-string")
				fs.Bool("sample-bool")
				fs.Duration("sample-duration")
				return fs
			},
			variables: map[string]string{
				"TEST2_SAMPLE_STRING":   "foo",
				"TEST2_SAMPLE_BOOL":     "yes",
				"TEST2_SAMPLE_DURATION": "15m",
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-duration", flagValue: time.Minute * 15, flagType: "duration"},
					{flagName: "sample-string", flagValue: "foo", flagType: "string"},
					{flagName: "sample-bool", flagValue: true, flagType: "bool"},
				},
				err:         true,
				expectedErr: ferrors.ErrParseFailed,
			},
		},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for k, v := range tt.variables {
				err := os.Setenv(k, v)
				require.NoError(t, err)
			}
			defer func() {
				for k := range tt.variables {
					err := os.Unsetenv(k)
					require.NoError(t, err)
				}
			}()
			fs := tt.flagSet()
			err := fs.BindEnvVars("-", "_")
			if tt.expected.err {
				require.Error(t, err)
				assert.True(t, errors.Is(err, tt.expected.expectedErr))
				return
			}
			require.NoError(t, err)
			for _, r := range tt.expected.parsed {
				require.NotNil(t, fs.flagByName(r.flagName).Value())
				assertParsedValue(t, fs, r)
			}
		})
	}
}
