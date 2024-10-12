package flagset

import (
	"errors"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/brongineer/helium/env"
	ferrors "github.com/brongineer/helium/errors"
	"github.com/brongineer/helium/flag"
	"github.com/brongineer/helium/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertParsedValue(t *testing.T, fs *FlagSet, r result) {
	switch r.flagType {
	case "string":
		val := GetString(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetStringPtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "stringSlice":
		val := GetStringSlice(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetStringSlicePtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "bool":
		val := GetBool(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetBoolPtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "boolSlice":
		val := GetBoolSlice(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetBoolSlicePtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "duration":
		val := GetDuration(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetDurationPtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "durationSlice":
		val := GetDurationSlice(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetDurationSlicePtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int":
		val := GetInt(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetIntPtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "intSlice":
		val := GetIntSlice(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetIntSlicePtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int8":
		val := GetInt8(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetInt8Ptr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int8slice":
		val := GetInt8Slice(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetInt8SlicePtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int16":
		val := GetInt16(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetInt16Ptr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int16slice":
		val := GetInt16Slice(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetInt16SlicePtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int32":
		val := GetInt32(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetInt32Ptr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int32slice":
		val := GetInt32Slice(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetInt32SlicePtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int64":
		val := GetInt64(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetInt64Ptr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "int64slice":
		val := GetInt64Slice(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetInt64SlicePtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint":
		val := GetUint(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetUintPtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uintSlice":
		val := GetUintSlice(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetUintSlicePtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint8":
		val := GetUint8(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetUint8Ptr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint8slice":
		val := GetUint8Slice(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetUint8SlicePtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint16":
		val := GetUint16(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetUint16Ptr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint16slice":
		val := GetUint16Slice(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetUint16SlicePtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint32":
		val := GetUint32(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetUint32Ptr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint32slice":
		val := GetUint32Slice(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetUint32SlicePtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint64":
		val := GetUint64(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetUint64Ptr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "uint64slice":
		val := GetUint64Slice(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetUint64SlicePtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "float32":
		val := GetFloat32(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetFloat32Ptr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "float32slice":
		val := GetFloat32Slice(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetFloat32SlicePtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "float64":
		val := GetFloat64(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetFloat64Ptr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "float64slice":
		val := GetFloat64Slice(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetFloat64SlicePtr(fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
	case "counter":
		val := GetCounter(fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
	case "custom":
		val := GetTypedFlag[custom](fs, r.flagName)
		assert.Equal(t, r.flagValue, val)
		ptr := GetTypedFlagPtr[custom](fs, r.flagName)
		require.NotNil(t, ptr)
		assert.Equal(t, r.flagValue, *ptr)
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

type custom struct {
	field int
}

type customParser struct {
	*parser.EmbeddedParser
}

func (p *customParser) ParseCmd(s string) (any, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}
	return &custom{field: v}, nil
}

func (p *customParser) ParseEnv(s string) (any, error) {
	return nil, nil
}

func TestFlagSet_Parse(t *testing.T) {
	t.Parallel()
	tests := []flagSetTest{
		{
			name: "parse",
			flagSet: func() *FlagSet {
				fs := New().
					BindFlag(flag.String("sample-string")).
					BindFlag(flag.Bool("sample-bool")).Build()
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
				fs := New().
					BindFlag(flag.String("sample-string")).
					BindFlag(flag.Bool("sample-bool")).Build()
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
				fs := New().
					BindFlag(flag.String("sample-string")).
					BindFlag(flag.Bool("sample-bool")).
					BindFlag(flag.Counter("sample-counter", flag.Shorthand("c"))).Build()
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
				fs := New().
					BindFlag(flag.Bool("sample-bool-one", flag.Shorthand("o"))).
					BindFlag(flag.Bool("sample-bool-two", flag.Shorthand("t"))).
					BindFlag(flag.Counter("sample-counter", flag.Shorthand("c"))).Build()
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
				fs := New().
					BindFlag(flag.Bool("sample-bool", flag.Shorthand("b"))).
					BindFlag(flag.Counter("sample-counter", flag.Shorthand("c"))).
					BindFlag(flag.Int("sample-int", flag.Shorthand("i"))).Build()
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
				fs := New().
					BindFlag(flag.String("sample-string")).
					BindFlag(flag.Bool("sample-bool")).Build()
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
				fs := New().
					BindFlag(flag.String("sample-string")).
					BindFlag(flag.BoolSlice("sample-bool")).Build()
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
				fs := New().
					BindFlag(flag.String("sample-string")).Build()
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
				fs := New().
					BindFlag(flag.String("sample-string", flag.Shorthand("s"))).Build()
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
				fs := New().
					BindFlag(flag.String("sample-string", flag.Shorthand("s"))).Build()
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
				fs := New().
					BindFlag(flag.String("sample-string", flag.Shorthand("s"))).Build()
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
				fs := New().
					BindFlag(flag.String("sample-string", flag.Shorthand("s"))).Build()
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
				fs := New().
					BindFlag(flag.Int("sample-int", flag.Shorthand("a"))).
					BindFlag(flag.Int8("sample-int8", flag.Shorthand("b"))).
					BindFlag(flag.Int16("sample-int16", flag.Shorthand("c"))).
					BindFlag(flag.Int32("sample-int32", flag.Shorthand("d"))).
					BindFlag(flag.Int64("sample-int64", flag.Shorthand("e"))).Build()
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
				fs := New().
					BindFlag(flag.Uint("sample-uint", flag.Shorthand("a"))).
					BindFlag(flag.Uint8("sample-uint8", flag.Shorthand("b"))).
					BindFlag(flag.Uint16("sample-uint16", flag.Shorthand("c"))).
					BindFlag(flag.Uint32("sample-uint32", flag.Shorthand("d"))).
					BindFlag(flag.Uint64("sample-uint64", flag.Shorthand("e"))).Build()
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
				fs := New().
					BindFlag(flag.IntSlice("sample-int", flag.Shorthand("a"))).
					BindFlag(flag.Int8Slice("sample-int8", flag.Shorthand("b"))).
					BindFlag(flag.Int16Slice("sample-int16", flag.Shorthand("c"))).
					BindFlag(flag.Int32Slice("sample-int32", flag.Shorthand("d"), flag.Separator("|"))).
					BindFlag(flag.Int64Slice("sample-int64", flag.Shorthand("e"), flag.Separator(";"))).Build()
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
				fs := New().
					BindFlag(flag.UintSlice("sample-uint", flag.Shorthand("a"))).
					BindFlag(flag.Uint8Slice("sample-uint8", flag.Shorthand("b"))).
					BindFlag(flag.Uint16Slice("sample-uint16", flag.Shorthand("c"))).
					BindFlag(flag.Uint32Slice("sample-uint32", flag.Shorthand("d"))).
					BindFlag(flag.Uint64Slice("sample-uint64", flag.Shorthand("e"))).Build()
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
				fs := New().
					BindFlag(flag.Float32("sample-float32", flag.Shorthand("a"))).
					BindFlag(flag.Float64("sample-float64", flag.Shorthand("b"))).
					BindFlag(flag.Float32Slice("sample-float32-slice", flag.Shorthand("c"))).
					BindFlag(flag.Float64Slice("sample-float64-slice", flag.Shorthand("d"))).Build()
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
				fs := New().
					BindFlag(flag.Duration("sample-duration", flag.Shorthand("a"))).
					BindFlag(flag.StringSlice("sample-string", flag.Shorthand("b"), flag.Separator(";"))).
					BindFlag(flag.DurationSlice("sample-duration-slice", flag.Shorthand("c"), flag.Separator(" "))).
					Build()
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
		{
			name: "custom flag no parser error",
			flagSet: func() *FlagSet {
				fs := New().BindFlag(flag.Typed[custom]("sample")).Build()
				return fs
			},
			expected: expected{
				parsed:      []result{},
				err:         true,
				expectedErr: ferrors.ErrNoParserDefined,
			},
			input: []string{"--sample"},
		},
		{
			name: "custom flag no error",
			flagSet: func() *FlagSet {
				fs := New().
					BindFlag(flag.Typed[custom]("sample", flag.Parser(&customParser{&parser.EmbeddedParser{}}))).
					Build()
				return fs
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample", flagValue: custom{field: 10}, flagType: "custom"},
				},
				err: false,
			},
			input: []string{"--sample", "10"},
		},
		{
			name: "custom flag parse error",
			flagSet: func() *FlagSet {
				fs := New().
					BindFlag(flag.Typed[custom]("sample", flag.Parser(&customParser{&parser.EmbeddedParser{}}))).
					Build()
				return fs
			},
			expected: expected{
				parsed:      []result{},
				err:         true,
				expectedErr: ferrors.ErrParseFailed,
			},
			input: []string{"--sample", "invalid"},
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
				fs := New(env.Prefix("test1"), env.Capitalized(), env.VarNameReplace("-", "_")).
					BindFlag(flag.String("sample-string")).
					BindFlag(flag.Bool("sample-bool")).
					BindFlag(flag.Counter("sample-counter")).
					BindFlag(flag.Float32("sample-float32")).
					BindFlag(flag.Float64("sample-float64")).
					BindFlag(flag.Duration("sample-duration")).Build()
				return fs
			},
			variables: map[string]string{
				"TEST1_SAMPLE_STRING":   "foo",
				"TEST1_SAMPLE_BOOL":     "true",
				"TEST1_SAMPLE_COUNTER":  "1",
				"TEST1_SAMPLE_FLOAT32":  "3.14",
				"TEST1_SAMPLE_FLOAT64":  "3.14",
				"TEST1_SAMPLE_DURATION": "15m",
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-duration", flagValue: time.Minute * 15, flagType: "duration"},
					{flagName: "sample-string", flagValue: "foo", flagType: "string"},
					{flagName: "sample-counter", flagValue: 1, flagType: "counter"},
					{flagName: "sample-float32", flagValue: float32(3.14), flagType: "float32"},
					{flagName: "sample-float64", flagValue: 3.14, flagType: "float64"},
					{flagName: "sample-bool", flagValue: true, flagType: "bool"},
				},
				err: false,
			},
		},
		{
			name: "env vars failed",
			flagSet: func() *FlagSet {
				fs := New(env.Prefix("test2"), env.Capitalized(), env.VarNameReplace("-", "_")).
					BindFlag(flag.String("sample-string")).
					BindFlag(flag.Bool("sample-bool")).
					BindFlag(flag.Counter("sample-counter")).
					BindFlag(flag.Float32("sample-float32")).
					BindFlag(flag.Float64("sample-float64")).
					BindFlag(flag.Duration("sample-duration")).Build()
				return fs
			},
			variables: map[string]string{
				"TEST2_SAMPLE_STRING":   "foo",
				"TEST2_SAMPLE_BOOL":     "yes",
				"TEST2_SAMPLE_COUNTER":  "one",
				"TEST2_SAMPLE_FLOAT32":  "3.4e+39",
				"TEST2_SAMPLE_FLOAT64":  "1.7e+308",
				"TEST2_SAMPLE_DURATION": "15m",
			},
			expected: expected{
				parsed:      []result{},
				err:         true,
				expectedErr: ferrors.ErrParseFailed,
			},
		},
		{
			name: "env vars slices success",
			flagSet: func() *FlagSet {
				fs := New(env.Prefix("test1-slice"), env.Capitalized(), env.VarNameReplace("-", "_")).
					BindFlag(flag.StringSlice("sample-string")).
					BindFlag(flag.BoolSlice("sample-bool")).
					BindFlag(flag.Float32Slice("sample-float32")).
					BindFlag(flag.Float64Slice("sample-float64")).
					BindFlag(flag.DurationSlice("sample-duration")).Build()
				return fs
			},
			variables: map[string]string{
				"TEST1_SLICE_SAMPLE_STRING":   "foo",
				"TEST1_SLICE_SAMPLE_BOOL":     "true",
				"TEST1_SLICE_SAMPLE_FLOAT32":  "3.14",
				"TEST1_SLICE_SAMPLE_FLOAT64":  "3.14",
				"TEST1_SLICE_SAMPLE_DURATION": "15m",
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-duration", flagValue: []time.Duration{time.Minute * 15}, flagType: "durationSlice"},
					{flagName: "sample-string", flagValue: []string{"foo"}, flagType: "stringSlice"},
					{flagName: "sample-float32", flagValue: []float32{3.14}, flagType: "float32slice"},
					{flagName: "sample-float64", flagValue: []float64{3.14}, flagType: "float64slice"},
					{flagName: "sample-bool", flagValue: []bool{true}, flagType: "boolSlice"},
				},
				err: false,
			},
		},
		{
			name: "env vars slices failed",
			flagSet: func() *FlagSet {
				fs := New(env.Prefix("test2-slice"), env.Capitalized(), env.VarNameReplace("-", "_")).
					BindFlag(flag.BoolSlice("sample-bool")).
					BindFlag(flag.Float32Slice("sample-float32")).
					BindFlag(flag.Float64Slice("sample-float64")).
					BindFlag(flag.DurationSlice("sample-duration")).Build()
				return fs
			},
			variables: map[string]string{
				"TEST2_SLICE_SAMPLE_BOOL":     "yes",
				"TEST2_SLICE_SAMPLE_FLOAT32":  "3.4e+39",
				"TEST2_SLICE_SAMPLE_FLOAT64":  "1.7e+308",
				"TEST2_SLICE_SAMPLE_DURATION": "15m",
			},
			expected: expected{
				parsed:      []result{},
				err:         true,
				expectedErr: ferrors.ErrParseFailed,
			},
		},
		{
			name: "env vars signed integers success",
			flagSet: func() *FlagSet {
				fs := New(env.Prefix("integers-success"), env.Capitalized(), env.VarNameReplace("-", "_")).
					BindFlag(flag.Int("sample-int")).
					BindFlag(flag.Int8("sample-int8")).
					BindFlag(flag.Int16("sample-int16")).
					BindFlag(flag.Int32("sample-int32")).
					BindFlag(flag.Int64("sample-int64")).Build()
				return fs
			},
			variables: map[string]string{
				"INTEGERS_SUCCESS_SAMPLE_INT":   "42",
				"INTEGERS_SUCCESS_SAMPLE_INT8":  "42",
				"INTEGERS_SUCCESS_SAMPLE_INT16": "42",
				"INTEGERS_SUCCESS_SAMPLE_INT32": "42",
				"INTEGERS_SUCCESS_SAMPLE_INT64": "42",
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-int", flagValue: 42, flagType: "int"},
					{flagName: "sample-int8", flagValue: int8(42), flagType: "int8"},
					{flagName: "sample-int16", flagValue: int16(42), flagType: "int16"},
					{flagName: "sample-int32", flagValue: int32(42), flagType: "int32"},
					{flagName: "sample-int64", flagValue: int64(42), flagType: "int64"},
				},
				err: false,
			},
		},
		{
			name: "env vars signed integers failure",
			flagSet: func() *FlagSet {
				fs := New(env.Prefix("integers-fail"), env.Capitalized(), env.VarNameReplace("-", "_")).
					BindFlag(flag.Int("sample-int")).
					BindFlag(flag.Int8("sample-int8")).
					BindFlag(flag.Int16("sample-int16")).
					BindFlag(flag.Int32("sample-int32")).
					BindFlag(flag.Int64("sample-int64")).Build()
				return fs
			},
			variables: map[string]string{
				"INTEGERS_FAIL_SAMPLE_INT":   "42.",
				"INTEGERS_FAIL_SAMPLE_INT8":  "1000",
				"INTEGERS_FAIL_SAMPLE_INT16": "-32800",
				"INTEGERS_FAIL_SAMPLE_INT32": "65656",
				"INTEGERS_FAIL_SAMPLE_INT64": "9223372036854775809",
			},
			expected: expected{
				parsed:      []result{},
				err:         true,
				expectedErr: ferrors.ErrParseFailed,
			},
		},
		{
			name: "env vars unsigned integers success",
			flagSet: func() *FlagSet {
				fs := New(env.Prefix("uintegers-success"), env.Capitalized(), env.VarNameReplace("-", "_")).
					BindFlag(flag.Uint("sample-uint")).
					BindFlag(flag.Uint8("sample-uint8")).
					BindFlag(flag.Uint16("sample-uint16")).
					BindFlag(flag.Uint32("sample-uint32")).
					BindFlag(flag.Uint64("sample-uint64")).Build()
				return fs
			},
			variables: map[string]string{
				"UINTEGERS_SUCCESS_SAMPLE_UINT":   "42",
				"UINTEGERS_SUCCESS_SAMPLE_UINT8":  "42",
				"UINTEGERS_SUCCESS_SAMPLE_UINT16": "42",
				"UINTEGERS_SUCCESS_SAMPLE_UINT32": "42",
				"UINTEGERS_SUCCESS_SAMPLE_UINT64": "42",
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-uint", flagValue: uint(42), flagType: "uint"},
					{flagName: "sample-uint8", flagValue: uint8(42), flagType: "uint8"},
					{flagName: "sample-uint16", flagValue: uint16(42), flagType: "uint16"},
					{flagName: "sample-uint32", flagValue: uint32(42), flagType: "uint32"},
					{flagName: "sample-uint64", flagValue: uint64(42), flagType: "uint64"},
				},
				err: false,
			},
		},
		{
			name: "env vars signed uintegers failure",
			flagSet: func() *FlagSet {
				fs := New(env.Prefix("uintegers-fail"), env.Capitalized(), env.VarNameReplace("-", "_")).
					BindFlag(flag.Uint("sample-uint")).
					BindFlag(flag.Uint8("sample-uint8")).
					BindFlag(flag.Uint16("sample-uint16")).
					BindFlag(flag.Uint32("sample-uint32")).
					BindFlag(flag.Uint64("sample-uint64")).Build()
				return fs
			},
			variables: map[string]string{
				"UINTEGERS_FAIL_SAMPLE_UINT":   "42.",
				"UINTEGERS_FAIL_SAMPLE_UINT8":  "1000",
				"UINTEGERS_FAIL_SAMPLE_UINT16": "-32800",
				"UINTEGERS_FAIL_SAMPLE_UINT32": "-165656",
				"UINTEGERS_FAIL_SAMPLE_UINT64": "184467440737095516150",
			},
			expected: expected{
				parsed:      []result{},
				err:         true,
				expectedErr: ferrors.ErrParseFailed,
			},
		},
		{
			name: "env vars signed integers slices success",
			flagSet: func() *FlagSet {
				fs := New(env.Prefix("integers-slice-success"), env.Capitalized(), env.VarNameReplace("-", "_")).
					BindFlag(flag.IntSlice("sample-int")).
					BindFlag(flag.Int8Slice("sample-int8")).
					BindFlag(flag.Int16Slice("sample-int16")).
					BindFlag(flag.Int32Slice("sample-int32")).
					BindFlag(flag.Int64Slice("sample-int64")).Build()
				return fs
			},
			variables: map[string]string{
				"INTEGERS_SLICE_SUCCESS_SAMPLE_INT":   "42",
				"INTEGERS_SLICE_SUCCESS_SAMPLE_INT8":  "42",
				"INTEGERS_SLICE_SUCCESS_SAMPLE_INT16": "42",
				"INTEGERS_SLICE_SUCCESS_SAMPLE_INT32": "42",
				"INTEGERS_SLICE_SUCCESS_SAMPLE_INT64": "42",
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-int", flagValue: []int{42}, flagType: "intSlice"},
					{flagName: "sample-int8", flagValue: []int8{42}, flagType: "int8slice"},
					{flagName: "sample-int16", flagValue: []int16{42}, flagType: "int16slice"},
					{flagName: "sample-int32", flagValue: []int32{42}, flagType: "int32slice"},
					{flagName: "sample-int64", flagValue: []int64{42}, flagType: "int64slice"},
				},
				err: false,
			},
		},
		{
			name: "env vars signed integers slices failure",
			flagSet: func() *FlagSet {
				fs := New(env.Prefix("integers-slice-failure"), env.Capitalized(), env.VarNameReplace("-", "_")).
					BindFlag(flag.IntSlice("sample-int")).
					BindFlag(flag.Int8Slice("sample-int8")).
					BindFlag(flag.Int16Slice("sample-int16")).
					BindFlag(flag.Int32Slice("sample-int32")).
					BindFlag(flag.Int64Slice("sample-int64")).Build()
				return fs
			},
			variables: map[string]string{
				"INTEGERS_SLICE_FAILURE_SAMPLE_INT":   "42.",
				"INTEGERS_SLICE_FAILURE_SAMPLE_INT8":  "1000",
				"INTEGERS_SLICE_FAILURE_SAMPLE_INT16": "-32800",
				"INTEGERS_SLICE_FAILURE_SAMPLE_INT32": "65656",
				"INTEGERS_SLICE_FAILURE_SAMPLE_INT64": "18446744073709551615",
			},
			expected: expected{
				parsed:      []result{},
				err:         true,
				expectedErr: ferrors.ErrParseFailed,
			},
		},
		{
			name: "env vars unsigned integers slices success",
			flagSet: func() *FlagSet {
				fs := New(env.Prefix("uintegers-slice-success"), env.Capitalized(), env.VarNameReplace("-", "_")).
					BindFlag(flag.UintSlice("sample-uint")).
					BindFlag(flag.Uint8Slice("sample-uint8")).
					BindFlag(flag.Uint16Slice("sample-uint16")).
					BindFlag(flag.Uint32Slice("sample-uint32")).
					BindFlag(flag.Uint64Slice("sample-uint64")).Build()
				return fs
			},
			variables: map[string]string{
				"UINTEGERS_SLICE_SUCCESS_SAMPLE_UINT":   "42",
				"UINTEGERS_SLICE_SUCCESS_SAMPLE_UINT8":  "42",
				"UINTEGERS_SLICE_SUCCESS_SAMPLE_UINT16": "42",
				"UINTEGERS_SLICE_SUCCESS_SAMPLE_UINT32": "42",
				"UINTEGERS_SLICE_SUCCESS_SAMPLE_UINT64": "42",
			},
			expected: expected{
				parsed: []result{
					{flagName: "sample-uint", flagValue: []uint{42}, flagType: "uintSlice"},
					{flagName: "sample-uint8", flagValue: []uint8{42}, flagType: "uint8slice"},
					{flagName: "sample-uint16", flagValue: []uint16{42}, flagType: "uint16slice"},
					{flagName: "sample-uint32", flagValue: []uint32{42}, flagType: "uint32slice"},
					{flagName: "sample-uint64", flagValue: []uint64{42}, flagType: "uint64slice"},
				},
				err: false,
			},
		},
		{
			name: "env vars unsigned integers slices failure",
			flagSet: func() *FlagSet {
				fs := New(env.Prefix("uintegers-slice-failure"), env.Capitalized(), env.VarNameReplace("-", "_")).
					BindFlag(flag.IntSlice("sample-uint")).
					BindFlag(flag.Int8Slice("sample-uint8")).
					BindFlag(flag.Int16Slice("sample-uint16")).
					BindFlag(flag.Int32Slice("sample-uint32")).
					BindFlag(flag.Int64Slice("sample-uint64")).Build()
				return fs
			},
			variables: map[string]string{
				"UINTEGERS_SLICE_FAILURE_SAMPLE_UINT":   "42.",
				"UINTEGERS_SLICE_FAILURE_SAMPLE_UINT8":  "1000",
				"UINTEGERS_SLICE_FAILURE_SAMPLE_UINT16": "-32800",
				"UINTEGERS_SLICE_FAILURE_SAMPLE_UINT32": "-65656",
				"UINTEGERS_SLICE_FAILURE_SAMPLE_UINT64": "1018446744073709551615",
			},
			expected: expected{
				parsed:      []result{},
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
			err := fs.BindEnvVars()
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
