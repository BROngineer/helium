package helium

import (
	"github.com/brongineer/helium/internal/generic"
	"github.com/brongineer/helium/options"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type expected struct {
	description  string
	shorthand    string
	separator    string
	defaultValue any
	required     bool
	shared       bool
}

func (e *expected) Description() string {
	return e.description
}

func (e *expected) Shorthand() string {
	return e.shorthand
}

func (e *expected) Separator() string {
	if e.separator == "" {
		return ","
	}
	return e.separator
}

func (e *expected) DefaultValue() any {
	return e.defaultValue
}

func (e *expected) Required() bool {
	return e.required
}

func (e *expected) Shared() bool {
	return e.shared
}

type result[T generic.Allowed] struct {
	some *T
	err  bool
}

func ptrTo[T any](v T) *T {
	return &v
}

type flagTest struct {
	name     string
	opts     []options.FlagOption
	expected expected
}

type getFlagTest[T generic.Allowed] struct {
	name   string
	opts   []options.FlagOption
	input  *string
	wanted result[T]
}

func assertFlag[T generic.Allowed](t *testing.T, fs *FlagSet, tt flagTest) {
	f := fs.flagByName(tt.name)
	assert.NotNil(t, f)
	assert.Equal(t, tt.expected.Description(), f.Description())
	assert.Equal(t, tt.expected.Shorthand(), f.Shorthand())
	assert.Equal(t, tt.expected.Separator(), f.Separator())
	if tt.expected.DefaultValue() != nil {
		assert.Equal(t, tt.expected.DefaultValue(), *f.Value().(*T))
	}
	assert.Equal(t, tt.expected.Required(), f.IsRequired())
	assert.Equal(t, tt.expected.Shared(), f.IsShared())
}

func assertGetFlag[T generic.Allowed](
	t *testing.T,
	fs *FlagSet,
	tt getFlagTest[T],
	getFunc func(string) T,
	getPtrFunc func(string) *T,
) {
	f := fs.flagByName(tt.name)
	assert.NotNil(t, f)
	if tt.input != nil {
		if tt.wanted.err {
			assert.Error(t, f.Parse(*tt.input))
			return
		}
		assert.NoError(t, f.Parse(*tt.input))
		assert.True(t, f.IsVisited())
	}
	assert.Equal(t, *tt.wanted.some, getFunc(tt.name))
	if getPtrFunc != nil {
		assert.Equal(t, *tt.wanted.some, *getPtrFunc(tt.name))
	}
}

func TestFlagSet_String(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue("test")},
			expected{defaultValue: "test"},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
				options.Shorthand("s"),
				options.Required(),
				options.Shared(),
			},
			expected{
				description: "description",
				shorthand:   "s",
				required:    true,
				shared:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.String(tt.name, tt.opts...)
			assertFlag[string](t, fs, tt)
		})
	}
}

func TestFlagSet_GetString(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[string]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("test"),
			result[string]{
				ptrTo("test"),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("1"),
			result[string]{
				ptrTo("1"),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue("default")},
			nil,
			result[string]{
				ptrTo("default"),
				false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.String(tt.name, tt.opts...)
			assertGetFlag[string](t, fs, tt, fs.GetString, fs.GetStringPtr)
		})
	}
}

func TestFlagSet_Duration(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(time.Second)},
			expected{defaultValue: time.Second},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
				options.Shorthand("s"),
				options.Required(),
				options.Shared(),
			},
			expected{
				description: "description",
				shorthand:   "s",
				required:    true,
				shared:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Duration(tt.name, tt.opts...)
			assertFlag[time.Duration](t, fs, tt)
		})
	}
}

func TestFlagSet_GetDuration(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[time.Duration]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("1s"),
			result[time.Duration]{
				ptrTo(time.Second),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("1h"),
			result[time.Duration]{
				ptrTo(time.Hour),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(time.Minute * 5)},
			nil,
			result[time.Duration]{
				ptrTo(time.Minute * 5),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[time.Duration]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Duration(tt.name, tt.opts...)
			assertGetFlag[time.Duration](t, fs, tt, fs.GetDuration, fs.GetDurationPtr)
		})
	}
}

func TestFlagSet_Bool(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(true)},
			expected{defaultValue: true},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
				options.Shorthand("s"),
				options.Required(),
				options.Shared(),
			},
			expected{
				description: "description",
				shorthand:   "s",
				required:    true,
				shared:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Bool(tt.name, tt.opts...)
			assertFlag[bool](t, fs, tt)
		})
	}
}

func TestFlagSet_GetBool(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[bool]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("false"),
			result[bool]{
				ptrTo(false),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("1"),
			result[bool]{
				ptrTo(true),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(true)},
			nil,
			result[bool]{
				ptrTo(true),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(false)},
			ptrTo(""),
			result[bool]{
				ptrTo(true),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("no"),
			result[bool]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Bool(tt.name, tt.opts...)
			assertGetFlag[bool](t, fs, tt, fs.GetBool, fs.GetBoolPtr)
		})
	}
}

func TestFlagSet_Int(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(42)},
			expected{defaultValue: 42},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int(tt.name, tt.opts...)
			assertFlag[int](t, fs, tt)
		})
	}
}

func TestFlagSet_GetInt(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[int]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42"),
			result[int]{
				ptrTo(42),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(42)},
			nil,
			result[int]{
				ptrTo(42),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[int]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int(tt.name, tt.opts...)
			assertGetFlag[int](t, fs, tt, fs.GetInt, fs.GetIntPtr)
		})
	}
}

func TestFlagSet_Int8(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(int8(42))},
			expected{defaultValue: int8(42)},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int8(tt.name, tt.opts...)
			assertFlag[int8](t, fs, tt)
		})
	}
}

func TestFlagSet_GetInt8(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[int8]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42"),
			result[int8]{
				ptrTo(int8(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(int8(42))},
			nil,
			result[int8]{
				ptrTo(int8(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[int8]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int8(tt.name, tt.opts...)
			assertGetFlag[int8](t, fs, tt, fs.GetInt8, fs.GetInt8Ptr)
		})
	}
}

func TestFlagSet_Int16(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(int16(42))},
			expected{defaultValue: int16(42)},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int16(tt.name, tt.opts...)
			assertFlag[int16](t, fs, tt)
		})
	}
}

func TestFlagSet_GetInt16(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[int16]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42"),
			result[int16]{
				ptrTo(int16(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(int16(42))},
			nil,
			result[int16]{
				ptrTo(int16(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[int16]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int16(tt.name, tt.opts...)
			assertGetFlag[int16](t, fs, tt, fs.GetInt16, fs.GetInt16Ptr)
		})
	}
}

func TestFlagSet_Int32(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(int32(42))},
			expected{defaultValue: int32(42)},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int32(tt.name, tt.opts...)
			assertFlag[int32](t, fs, tt)
		})
	}
}

func TestFlagSet_GetInt32(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[int32]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42"),
			result[int32]{
				ptrTo(int32(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(int32(42))},
			nil,
			result[int32]{
				ptrTo(int32(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[int32]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int32(tt.name, tt.opts...)
			assertGetFlag[int32](t, fs, tt, fs.GetInt32, fs.GetInt32Ptr)
		})
	}
}

func TestFlagSet_Int64(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(int64(42))},
			expected{defaultValue: int64(42)},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int64(tt.name, tt.opts...)
			assertFlag[int64](t, fs, tt)
		})
	}
}

func TestFlagSet_GetInt64(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[int64]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42"),
			result[int64]{
				ptrTo(int64(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(int64(42))},
			nil,
			result[int64]{
				ptrTo(int64(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[int64]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int64(tt.name, tt.opts...)
			assertGetFlag[int64](t, fs, tt, fs.GetInt64, fs.GetInt64Ptr)
		})
	}
}

func TestFlagSet_Uint(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(uint(42))},
			expected{defaultValue: uint(42)},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint(tt.name, tt.opts...)
			assertFlag[uint](t, fs, tt)
		})
	}
}

func TestFlagSet_GetUint(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[uint]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42"),
			result[uint]{
				ptrTo(uint(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(uint(42))},
			nil,
			result[uint]{
				ptrTo(uint(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[uint]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint(tt.name, tt.opts...)
			assertGetFlag[uint](t, fs, tt, fs.GetUint, fs.GetUintPtr)
		})
	}
}

func TestFlagSet_Uint8(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(uint8(42))},
			expected{defaultValue: uint8(42)},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint8(tt.name, tt.opts...)
			assertFlag[uint8](t, fs, tt)
		})
	}
}

func TestFlagSet_GetUint8(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[uint8]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42"),
			result[uint8]{
				ptrTo(uint8(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(uint8(42))},
			nil,
			result[uint8]{
				ptrTo(uint8(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[uint8]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint8(tt.name, tt.opts...)
			assertGetFlag[uint8](t, fs, tt, fs.GetUint8, fs.GetUint8Ptr)
		})
	}
}

func TestFlagSet_Uint16(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(uint16(42))},
			expected{defaultValue: uint16(42)},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint16(tt.name, tt.opts...)
			assertFlag[uint16](t, fs, tt)
		})
	}
}

func TestFlagSet_GetUint16(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[uint16]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42"),
			result[uint16]{
				ptrTo(uint16(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(uint16(42))},
			nil,
			result[uint16]{
				ptrTo(uint16(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[uint16]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint16(tt.name, tt.opts...)
			assertGetFlag[uint16](t, fs, tt, fs.GetUint16, fs.GetUint16Ptr)
		})
	}
}

func TestFlagSet_Uint32(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(uint32(42))},
			expected{defaultValue: uint32(42)},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint32(tt.name, tt.opts...)
			assertFlag[uint32](t, fs, tt)
		})
	}
}

func TestFlagSet_GetUint32(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[uint32]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42"),
			result[uint32]{
				ptrTo(uint32(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(uint32(42))},
			nil,
			result[uint32]{
				ptrTo(uint32(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[uint32]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint32(tt.name, tt.opts...)
			assertGetFlag[uint32](t, fs, tt, fs.GetUint32, fs.GetUint32Ptr)
		})
	}
}

func TestFlagSet_Uint64(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(uint64(42))},
			expected{defaultValue: uint64(42)},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint64(tt.name, tt.opts...)
			assertFlag[uint64](t, fs, tt)
		})
	}
}

func TestFlagSet_GetUint64(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[uint64]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42"),
			result[uint64]{
				ptrTo(uint64(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(uint64(42))},
			nil,
			result[uint64]{
				ptrTo(uint64(42)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[uint64]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint64(tt.name, tt.opts...)
			assertGetFlag[uint64](t, fs, tt, fs.GetUint64, fs.GetUint64Ptr)
		})
	}
}

func TestFlagSet_Float32(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(float32(42.0))},
			expected{defaultValue: float32(42.0)},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Float32(tt.name, tt.opts...)
			assertFlag[float32](t, fs, tt)
		})
	}
}

func TestFlagSet_GetFloat32(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[float32]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42.0"),
			result[float32]{
				ptrTo(float32(42.0)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(float32(42.0))},
			nil,
			result[float32]{
				ptrTo(float32(42.0)),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[float32]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Float32(tt.name, tt.opts...)
			assertGetFlag[float32](t, fs, tt, fs.GetFloat32, fs.GetFloat32Ptr)
		})
	}
}

func TestFlagSet_Float64(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(42.0)},
			expected{defaultValue: 42.0},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Float64(tt.name, tt.opts...)
			assertFlag[float64](t, fs, tt)
		})
	}
}

func TestFlagSet_GetFloat64(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[float64]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42.0"),
			result[float64]{
				ptrTo(42.0),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(42.0)},
			nil,
			result[float64]{
				ptrTo(42.0),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[float64]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Float64(tt.name, tt.opts...)
			assertGetFlag[float64](t, fs, tt, fs.GetFloat64, fs.GetFloat64Ptr)
		})
	}
}

func TestFlagSet_IntSlice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]int{42})},
			expected{defaultValue: []int{42}},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.IntSlice(tt.name, tt.opts...)
			assertFlag[[]int](t, fs, tt)
		})
	}
}

func TestFlagSet_GetIntSlice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]int]{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			ptrTo("42|1"),
			result[[]int]{
				ptrTo([]int{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42,1"),
			result[[]int]{
				ptrTo([]int{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]int{42, 1})},
			nil,
			result[[]int]{
				ptrTo([]int{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[[]int]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.IntSlice(tt.name, tt.opts...)
			assertGetFlag[[]int](t, fs, tt, fs.GetIntSlice, fs.GetIntSlicePtr)
		})
	}
}

func TestFlagSet_Int8Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]int8{42})},
			expected{defaultValue: []int8{42}},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int8Slice(tt.name, tt.opts...)
			assertFlag[[]int8](t, fs, tt)
		})
	}
}

func TestFlagSet_GetInt8Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]int8]{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			ptrTo("42|1"),
			result[[]int8]{
				ptrTo([]int8{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42,1"),
			result[[]int8]{
				ptrTo([]int8{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]int8{42, 1})},
			nil,
			result[[]int8]{
				ptrTo([]int8{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[[]int8]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int8Slice(tt.name, tt.opts...)
			assertGetFlag[[]int8](t, fs, tt, fs.GetInt8Slice, fs.GetInt8SlicePtr)
		})
	}
}

func TestFlagSet_Int16Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]int16{42})},
			expected{defaultValue: []int16{42}},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int16Slice(tt.name, tt.opts...)
			assertFlag[[]int16](t, fs, tt)
		})
	}
}

func TestFlagSet_GetInt16Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]int16]{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			ptrTo("42|1"),
			result[[]int16]{
				ptrTo([]int16{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42,1"),
			result[[]int16]{
				ptrTo([]int16{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]int16{42, 1})},
			nil,
			result[[]int16]{
				ptrTo([]int16{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[[]int16]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int16Slice(tt.name, tt.opts...)
			assertGetFlag[[]int16](t, fs, tt, fs.GetInt16Slice, fs.GetInt16SlicePtr)
		})
	}
}

func TestFlagSet_Int32Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]int32{42})},
			expected{defaultValue: []int32{42}},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int32Slice(tt.name, tt.opts...)
			assertFlag[[]int32](t, fs, tt)
		})
	}
}

func TestFlagSet_GetInt32Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]int32]{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			ptrTo("42|1"),
			result[[]int32]{
				ptrTo([]int32{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42,1"),
			result[[]int32]{
				ptrTo([]int32{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]int32{42, 1})},
			nil,
			result[[]int32]{
				ptrTo([]int32{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[[]int32]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int32Slice(tt.name, tt.opts...)
			assertGetFlag[[]int32](t, fs, tt, fs.GetInt32Slice, fs.GetInt32SlicePtr)
		})
	}
}

func TestFlagSet_Int64Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]int64{42})},
			expected{defaultValue: []int64{42}},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int64Slice(tt.name, tt.opts...)
			assertFlag[[]int64](t, fs, tt)
		})
	}
}

func TestFlagSet_GetInt64Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]int64]{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			ptrTo("42|1"),
			result[[]int64]{
				ptrTo([]int64{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42,1"),
			result[[]int64]{
				ptrTo([]int64{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]int64{42, 1})},
			nil,
			result[[]int64]{
				ptrTo([]int64{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[[]int64]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Int64Slice(tt.name, tt.opts...)
			assertGetFlag[[]int64](t, fs, tt, fs.GetInt64Slice, fs.GetInt64SlicePtr)
		})
	}
}

func TestFlagSet_UintSlice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]uint{42})},
			expected{defaultValue: []uint{42}},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.UintSlice(tt.name, tt.opts...)
			assertFlag[[]uint](t, fs, tt)
		})
	}
}

func TestFlagSet_GetUintSlice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]uint]{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			ptrTo("42|1"),
			result[[]uint]{
				ptrTo([]uint{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42,1"),
			result[[]uint]{
				ptrTo([]uint{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]uint{42, 1})},
			nil,
			result[[]uint]{
				ptrTo([]uint{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[[]uint]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.UintSlice(tt.name, tt.opts...)
			assertGetFlag[[]uint](t, fs, tt, fs.GetUintSlice, fs.GetUintSlicePtr)
		})
	}
}

func TestFlagSet_Uint8Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]uint8{42})},
			expected{defaultValue: []uint8{42}},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint8Slice(tt.name, tt.opts...)
			assertFlag[[]uint8](t, fs, tt)
		})
	}
}

func TestFlagSet_GetUint8Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]uint8]{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			ptrTo("42|1"),
			result[[]uint8]{
				ptrTo([]uint8{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42,1"),
			result[[]uint8]{
				ptrTo([]uint8{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]uint8{42, 1})},
			nil,
			result[[]uint8]{
				ptrTo([]uint8{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[[]uint8]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint8Slice(tt.name, tt.opts...)
			assertGetFlag[[]uint8](t, fs, tt, fs.GetUint8Slice, fs.GetUint8SlicePtr)
		})
	}
}

func TestFlagSet_Uint16Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]uint16{42})},
			expected{defaultValue: []uint16{42}},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint16Slice(tt.name, tt.opts...)
			assertFlag[[]uint16](t, fs, tt)
		})
	}
}

func TestFlagSet_GetUint16Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]uint16]{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			ptrTo("42|1"),
			result[[]uint16]{
				ptrTo([]uint16{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42,1"),
			result[[]uint16]{
				ptrTo([]uint16{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]uint16{42, 1})},
			nil,
			result[[]uint16]{
				ptrTo([]uint16{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[[]uint16]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint16Slice(tt.name, tt.opts...)
			assertGetFlag[[]uint16](t, fs, tt, fs.GetUint16Slice, fs.GetUint16SlicePtr)
		})
	}
}

func TestFlagSet_Uint32Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]uint32{42})},
			expected{defaultValue: []uint32{42}},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint32Slice(tt.name, tt.opts...)
			assertFlag[[]uint32](t, fs, tt)
		})
	}
}

func TestFlagSet_GetUint32Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]uint32]{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			ptrTo("42|1"),
			result[[]uint32]{
				ptrTo([]uint32{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42,1"),
			result[[]uint32]{
				ptrTo([]uint32{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]uint32{42, 1})},
			nil,
			result[[]uint32]{
				ptrTo([]uint32{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[[]uint32]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint32Slice(tt.name, tt.opts...)
			assertGetFlag[[]uint32](t, fs, tt, fs.GetUint32Slice, fs.GetUint32SlicePtr)
		})
	}
}

func TestFlagSet_Uint64Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]uint64{42})},
			expected{defaultValue: []uint64{42}},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint64Slice(tt.name, tt.opts...)
			assertFlag[[]uint64](t, fs, tt)
		})
	}
}

func TestFlagSet_GetUint64Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]uint64]{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			ptrTo("42|1"),
			result[[]uint64]{
				ptrTo([]uint64{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42,1"),
			result[[]uint64]{
				ptrTo([]uint64{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]uint64{42, 1})},
			nil,
			result[[]uint64]{
				ptrTo([]uint64{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[[]uint64]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Uint64Slice(tt.name, tt.opts...)
			assertGetFlag[[]uint64](t, fs, tt, fs.GetUint64Slice, fs.GetUint64SlicePtr)
		})
	}
}

func TestFlagSet_Float32Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]float32{42.0})},
			expected{defaultValue: []float32{42.0}},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Float32Slice(tt.name, tt.opts...)
			assertFlag[[]float32](t, fs, tt)
		})
	}
}

func TestFlagSet_GetFloat32Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]float32]{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			ptrTo("42.0|2.7|3.14"),
			result[[]float32]{
				ptrTo([]float32{42.0, 2.7, 3.14}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42.0,2.7"),
			result[[]float32]{
				ptrTo([]float32{42.0, 2.7}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]float32{42.0, 2.7})},
			nil,
			result[[]float32]{
				ptrTo([]float32{42.0, 2.7}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[[]float32]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Float32Slice(tt.name, tt.opts...)
			assertGetFlag[[]float32](t, fs, tt, fs.GetFloat32Slice, fs.GetFloat32SlicePtr)
		})
	}
}

func TestFlagSet_Float64Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]float64{42.0})},
			expected{defaultValue: []float64{42.0}},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Float64Slice(tt.name, tt.opts...)
			assertFlag[[]float64](t, fs, tt)
		})
	}
}

func TestFlagSet_GetFloat64Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]float64]{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			ptrTo("42.0|2.7|3.14"),
			result[[]float64]{
				ptrTo([]float64{42.0, 2.7, 3.14}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42.0,2.7"),
			result[[]float64]{
				ptrTo([]float64{42.0, 2.7}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]float64{42.0, 2.7})},
			nil,
			result[[]float64]{
				ptrTo([]float64{42.0, 2.7}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[[]float64]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Float64Slice(tt.name, tt.opts...)
			assertGetFlag[[]float64](t, fs, tt, fs.GetFloat64Slice, fs.GetFloat64SlicePtr)
		})
	}
}

func TestFlagSet_StringSlice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]string{"foo", "bar"})},
			expected{defaultValue: []string{"foo", "bar"}},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.StringSlice(tt.name, tt.opts...)
			assertFlag[[]string](t, fs, tt)
		})
	}
}

func TestFlagSet_GetStringSlice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]string]{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			ptrTo("foo|bar"),
			result[[]string]{
				ptrTo([]string{"foo", "bar"}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42,1"),
			result[[]string]{
				ptrTo([]string{"42", "1"}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]string{"foo", "bar"})},
			nil,
			result[[]string]{
				ptrTo([]string{"foo", "bar"}),
				false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.StringSlice(tt.name, tt.opts...)
			assertGetFlag[[]string](t, fs, tt, fs.GetStringSlice, fs.GetStringSlicePtr)
		})
	}
}

func TestFlagSet_DurationSlice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]time.Duration{time.Minute})},
			expected{defaultValue: []time.Duration{time.Minute}},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.DurationSlice(tt.name, tt.opts...)
			assertFlag[[]time.Duration](t, fs, tt)
		})
	}
}

func TestFlagSet_GetDurationSlice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]time.Duration]{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			ptrTo("42s|1m"),
			result[[]time.Duration]{
				ptrTo([]time.Duration{time.Second * 42, time.Minute}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("42s,1m"),
			result[[]time.Duration]{
				ptrTo([]time.Duration{time.Second * 42, time.Minute}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]time.Duration{time.Second * 42})},
			nil,
			result[[]time.Duration]{
				ptrTo([]time.Duration{time.Second * 42}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("invalid"),
			result[[]time.Duration]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.DurationSlice(tt.name, tt.opts...)
			assertGetFlag[[]time.Duration](t, fs, tt, fs.GetDurationSlice, fs.GetDurationSlicePtr)
		})
	}
}

func TestFlagSet_BoolSlice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]bool{true, false})},
			expected{defaultValue: []bool{true, false}},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.BoolSlice(tt.name, tt.opts...)
			assertFlag[[]bool](t, fs, tt)
		})
	}
}

func TestFlagSet_GetBoolSlice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]bool]{
		{
			"sample",
			[]options.FlagOption{options.Separator("|")},
			ptrTo("t|f"),
			result[[]bool]{
				ptrTo([]bool{true, false}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("F,true"),
			result[[]bool]{
				ptrTo([]bool{false, true}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue([]bool{false, false})},
			nil,
			result[[]bool]{
				ptrTo([]bool{false, false}),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{},
			ptrTo("no,no"),
			result[[]bool]{
				nil,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.BoolSlice(tt.name, tt.opts...)
			assertGetFlag[[]bool](t, fs, tt, fs.GetBoolSlice, fs.GetBoolSlicePtr)
		})
	}
}

func TestFlagSet_Counter(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]options.FlagOption{},
			expected{},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(5)},
			expected{defaultValue: generic.NewCounter(5)},
		},
		{
			"sample",
			[]options.FlagOption{
				options.Description("description"),
				options.Shorthand("s"),
				options.Required(),
				options.Shared(),
			},
			expected{
				description: "description",
				shorthand:   "s",
				required:    true,
				shared:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Counter(tt.name, tt.opts...)
			assertFlag[generic.Counter](t, fs, tt)
		})
	}
}

func TestFlagSet_GetCounter(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[int]{
		{
			"sample",
			[]options.FlagOption{},
			ptrTo(""),
			result[int]{
				ptrTo(1),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(1)},
			ptrTo(""),
			result[int]{
				ptrTo(2),
				false,
			},
		},
		{
			"sample",
			[]options.FlagOption{options.DefaultValue(41)},
			ptrTo(""),
			result[int]{
				ptrTo(42),
				false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagSet{}
			fs.Counter(tt.name, tt.opts...)
			assertGetFlag[int](t, fs, tt, fs.GetCounter, nil)
		})
	}
}
