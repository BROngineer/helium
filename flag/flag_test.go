package flag

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
	Parser() flagParser
}

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

type result[T any] struct {
	some *T
	err  bool
}

func ptrTo[T any](v T) *T {
	return &v
}

type flagTest struct {
	name     string
	opts     []Option
	expected expected
}

type getFlagTest[T any] struct {
	name   string
	opts   []Option
	input  *string
	wanted result[T]
}

func assertFlag[T any](t *testing.T, f flagPropertyGetter, tt flagTest) {
	assert.NotNil(t, f)
	assert.Equal(t, tt.expected.Description(), f.Description())
	assert.Equal(t, tt.expected.Shorthand(), f.Shorthand())
	assert.Equal(t, tt.expected.Separator(), f.Separator())
	if tt.expected.DefaultValue() != nil {
		actual, ok := f.Value().(*T)
		assert.True(t, ok)
		assert.Equal(t, tt.expected.DefaultValue(), *actual)
	}
	assert.Equal(t, tt.expected.Shared(), f.IsShared())
}

func assertGetFlag[T any](t *testing.T, f flagPropertyGetter, tt getFlagTest[T]) {
	assert.NotNil(t, f)
	if tt.input != nil {
		if tt.wanted.err {
			assert.Error(t, f.FromCommandLine(*tt.input))
			return
		}
		assert.NoError(t, f.FromCommandLine(*tt.input))
		assert.True(t, f.IsVisited())
	}
	assert.Equal(t, *tt.wanted.some, DerefOrDie[T](f.Value()))
	assert.Equal(t, tt.wanted.some, PtrOrDie[T](f.Value()))
}

type custom struct {
	field int
}

type customParser struct {
	*embeddedParser
}

func newCustomParser() *customParser {
	return &customParser{&embeddedParser{}}
}

func (p *customParser) ParseCmd(s string) (any, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}
	return &custom{field: v}, nil
}

func (p *customParser) ParseEnv(_ string) (any, error) {
	return nil, nil
}

func TestFlag_Custom(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{Parser(newCustomParser())},
			expected{},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewCustom[custom](tt.name, tt.opts...)
			assertFlag[custom](t, f, tt)
		})
	}
}

func TestFlag_GetCustom(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[custom]{
		{
			"sample",
			[]Option{},
			ptrTo("test"),
			result[custom]{err: true},
		},
		{
			"sample",
			[]Option{Parser(newCustomParser())},
			ptrTo("10"),
			result[custom]{
				ptrTo(custom{field: 10}),
				false,
			},
		},
		{
			"sample",
			[]Option{Parser(newCustomParser())},
			ptrTo("invalid"),
			result[custom]{err: true},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewCustom[custom](tt.name, tt.opts...)
			assertGetFlag[custom](t, f, tt)
		})
	}
}

func TestFlag_String(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue("test")},
			expected{defaultValue: "test"},
		},
		{
			"sample",
			[]Option{
				Description("description"),
				Shorthand("s"),
				Shared(),
			},
			expected{
				description: "description",
				shorthand:   "s",
				required:    true,
				shared:      true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewString(tt.name, tt.opts...)
			assertFlag[string](t, f, tt)
		})
	}
}

func TestFlag_GetString(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[string]{
		{
			"sample",
			[]Option{},
			ptrTo("test"),
			result[string]{
				ptrTo("test"),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("1"),
			result[string]{
				ptrTo("1"),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue("default")},
			nil,
			result[string]{
				ptrTo("default"),
				false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewString(tt.name, tt.opts...)
			assertGetFlag[string](t, f, tt)
		})
	}
}

func TestFlag_Duration(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue(time.Second)},
			expected{defaultValue: time.Second},
		},
		{
			"sample",
			[]Option{
				Description("description"),
				Shorthand("s"),
				Shared(),
			},
			expected{
				description: "description",
				shorthand:   "s",
				required:    true,
				shared:      true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewDuration(tt.name, tt.opts...)
			assertFlag[time.Duration](t, f, tt)
		})
	}
}

func TestFlag_GetDuration(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[time.Duration]{
		{
			"sample",
			[]Option{},
			ptrTo("1s"),
			result[time.Duration]{
				ptrTo(time.Second),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("1h"),
			result[time.Duration]{
				ptrTo(time.Hour),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(time.Minute * 5)},
			nil,
			result[time.Duration]{
				ptrTo(time.Minute * 5),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[time.Duration]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewDuration(tt.name, tt.opts...)
			assertGetFlag[time.Duration](t, f, tt)
		})
	}
}

func TestFlag_Bool(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue(true)},
			expected{defaultValue: true},
		},
		{
			"sample",
			[]Option{
				Description("description"),
				Shorthand("s"),
				Shared(),
			},
			expected{
				description: "description",
				shorthand:   "s",
				required:    true,
				shared:      true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewBool(tt.name, tt.opts...)
			assertFlag[bool](t, f, tt)
		})
	}
}

func TestFlag_GetBool(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[bool]{
		{
			"sample",
			[]Option{},
			ptrTo("false"),
			result[bool]{
				ptrTo(false),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("1"),
			result[bool]{
				ptrTo(true),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(true)},
			nil,
			result[bool]{
				ptrTo(true),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(false)},
			ptrTo(""),
			result[bool]{
				ptrTo(true),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("no"),
			result[bool]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewBool(tt.name, tt.opts...)
			assertGetFlag[bool](t, f, tt)
		})
	}
}

func TestFlag_Int(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue(42)},
			expected{defaultValue: 42},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt(tt.name, tt.opts...)
			assertFlag[int](t, f, tt)
		})
	}
}

func TestFlag_GetInt(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[int]{
		{
			"sample",
			[]Option{},
			ptrTo("42"),
			result[int]{
				ptrTo(42),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(42)},
			nil,
			result[int]{
				ptrTo(42),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[int]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt(tt.name, tt.opts...)
			assertGetFlag[int](t, f, tt)
		})
	}
}

func TestFlag_Int8(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue(int8(42))},
			expected{defaultValue: int8(42)},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt8(tt.name, tt.opts...)
			assertFlag[int8](t, f, tt)
		})
	}
}

func TestFlag_GetInt8(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[int8]{
		{
			"sample",
			[]Option{},
			ptrTo("42"),
			result[int8]{
				ptrTo(int8(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(int8(42))},
			nil,
			result[int8]{
				ptrTo(int8(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[int8]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt8(tt.name, tt.opts...)
			assertGetFlag[int8](t, f, tt)
		})
	}
}

func TestFlag_Int16(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue(int16(42))},
			expected{defaultValue: int16(42)},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt16(tt.name, tt.opts...)
			assertFlag[int16](t, f, tt)
		})
	}
}

func TestFlag_GetInt16(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[int16]{
		{
			"sample",
			[]Option{},
			ptrTo("42"),
			result[int16]{
				ptrTo(int16(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(int16(42))},
			nil,
			result[int16]{
				ptrTo(int16(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[int16]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt16(tt.name, tt.opts...)
			assertGetFlag[int16](t, f, tt)
		})
	}
}

func TestFlag_Int32(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue(int32(42))},
			expected{defaultValue: int32(42)},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt32(tt.name, tt.opts...)
			assertFlag[int32](t, f, tt)
		})
	}
}

func TestFlag_GetInt32(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[int32]{
		{
			"sample",
			[]Option{},
			ptrTo("42"),
			result[int32]{
				ptrTo(int32(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(int32(42))},
			nil,
			result[int32]{
				ptrTo(int32(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[int32]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt32(tt.name, tt.opts...)
			assertGetFlag[int32](t, f, tt)
		})
	}
}

func TestFlag_Int64(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue(int64(42))},
			expected{defaultValue: int64(42)},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt64(tt.name, tt.opts...)
			assertFlag[int64](t, f, tt)
		})
	}
}

func TestFlag_GetInt64(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[int64]{
		{
			"sample",
			[]Option{},
			ptrTo("42"),
			result[int64]{
				ptrTo(int64(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(int64(42))},
			nil,
			result[int64]{
				ptrTo(int64(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[int64]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt64(tt.name, tt.opts...)
			assertGetFlag[int64](t, f, tt)
		})
	}
}

func TestFlag_Uint(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue(uint(42))},
			expected{defaultValue: uint(42)},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint(tt.name, tt.opts...)
			assertFlag[uint](t, f, tt)
		})
	}
}

func TestFlag_GetUint(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[uint]{
		{
			"sample",
			[]Option{},
			ptrTo("42"),
			result[uint]{
				ptrTo(uint(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(uint(42))},
			nil,
			result[uint]{
				ptrTo(uint(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[uint]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint(tt.name, tt.opts...)
			assertGetFlag[uint](t, f, tt)
		})
	}
}

func TestFlag_Uint8(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue(uint8(42))},
			expected{defaultValue: uint8(42)},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint8(tt.name, tt.opts...)
			assertFlag[uint8](t, f, tt)
		})
	}
}

func TestFlag_GetUint8(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[uint8]{
		{
			"sample",
			[]Option{},
			ptrTo("42"),
			result[uint8]{
				ptrTo(uint8(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(uint8(42))},
			nil,
			result[uint8]{
				ptrTo(uint8(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[uint8]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint8(tt.name, tt.opts...)
			assertGetFlag[uint8](t, f, tt)
		})
	}
}

func TestFlag_Uint16(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue(uint16(42))},
			expected{defaultValue: uint16(42)},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint16(tt.name, tt.opts...)
			assertFlag[uint16](t, f, tt)
		})
	}
}

func TestFlag_GetUint16(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[uint16]{
		{
			"sample",
			[]Option{},
			ptrTo("42"),
			result[uint16]{
				ptrTo(uint16(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(uint16(42))},
			nil,
			result[uint16]{
				ptrTo(uint16(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[uint16]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint16(tt.name, tt.opts...)
			assertGetFlag[uint16](t, f, tt)
		})
	}
}

func TestFlag_Uint32(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue(uint32(42))},
			expected{defaultValue: uint32(42)},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint32(tt.name, tt.opts...)
			assertFlag[uint32](t, f, tt)
		})
	}
}

func TestFlag_GetUint32(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[uint32]{
		{
			"sample",
			[]Option{},
			ptrTo("42"),
			result[uint32]{
				ptrTo(uint32(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(uint32(42))},
			nil,
			result[uint32]{
				ptrTo(uint32(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[uint32]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint32(tt.name, tt.opts...)
			assertGetFlag[uint32](t, f, tt)
		})
	}
}

func TestFlag_Uint64(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue(uint64(42))},
			expected{defaultValue: uint64(42)},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint64(tt.name, tt.opts...)
			assertFlag[uint64](t, f, tt)
		})
	}
}

func TestFlag_GetUint64(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[uint64]{
		{
			"sample",
			[]Option{},
			ptrTo("42"),
			result[uint64]{
				ptrTo(uint64(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(uint64(42))},
			nil,
			result[uint64]{
				ptrTo(uint64(42)),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[uint64]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint64(tt.name, tt.opts...)
			assertGetFlag[uint64](t, f, tt)
		})
	}
}

func TestFlag_Float32(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue(float32(42.0))},
			expected{defaultValue: float32(42.0)},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewFloat32(tt.name, tt.opts...)
			assertFlag[float32](t, f, tt)
		})
	}
}

func TestFlag_GetFloat32(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[float32]{
		{
			"sample",
			[]Option{},
			ptrTo("42.0"),
			result[float32]{
				ptrTo(float32(42.0)),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(float32(42.0))},
			nil,
			result[float32]{
				ptrTo(float32(42.0)),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[float32]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewFloat32(tt.name, tt.opts...)
			assertGetFlag[float32](t, f, tt)
		})
	}
}

func TestFlag_Float64(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue(42.0)},
			expected{defaultValue: 42.0},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewFloat64(tt.name, tt.opts...)
			assertFlag[float64](t, f, tt)
		})
	}
}

func TestFlag_GetFloat64(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[float64]{
		{
			"sample",
			[]Option{},
			ptrTo("42.0"),
			result[float64]{
				ptrTo(42.0),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(42.0)},
			nil,
			result[float64]{
				ptrTo(42.0),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[float64]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewFloat64(tt.name, tt.opts...)
			assertGetFlag[float64](t, f, tt)
		})
	}
}

func TestFlag_IntSlice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]int{42})},
			expected{defaultValue: []int{42}},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewIntSlice(tt.name, tt.opts...)
			assertFlag[[]int](t, f, tt)
		})
	}
}

func TestFlag_GetIntSlice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]int]{
		{
			"sample",
			[]Option{Separator("|")},
			ptrTo("42|1"),
			result[[]int]{
				ptrTo([]int{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("42,1"),
			result[[]int]{
				ptrTo([]int{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]int{42, 1})},
			nil,
			result[[]int]{
				ptrTo([]int{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[[]int]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewIntSlice(tt.name, tt.opts...)
			assertGetFlag[[]int](t, f, tt)
		})
	}
}

func TestFlag_Int8Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]int8{42})},
			expected{defaultValue: []int8{42}},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt8Slice(tt.name, tt.opts...)
			assertFlag[[]int8](t, f, tt)
		})
	}
}

func TestFlag_GetInt8Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]int8]{
		{
			"sample",
			[]Option{Separator("|")},
			ptrTo("42|1"),
			result[[]int8]{
				ptrTo([]int8{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("42,1"),
			result[[]int8]{
				ptrTo([]int8{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]int8{42, 1})},
			nil,
			result[[]int8]{
				ptrTo([]int8{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[[]int8]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt8Slice(tt.name, tt.opts...)
			assertGetFlag[[]int8](t, f, tt)
		})
	}
}

func TestFlag_Int16Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]int16{42})},
			expected{defaultValue: []int16{42}},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt16Slice(tt.name, tt.opts...)
			assertFlag[[]int16](t, f, tt)
		})
	}
}

func TestFlag_GetInt16Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]int16]{
		{
			"sample",
			[]Option{Separator("|")},
			ptrTo("42|1"),
			result[[]int16]{
				ptrTo([]int16{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("42,1"),
			result[[]int16]{
				ptrTo([]int16{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]int16{42, 1})},
			nil,
			result[[]int16]{
				ptrTo([]int16{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[[]int16]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt16Slice(tt.name, tt.opts...)
			assertGetFlag[[]int16](t, f, tt)
		})
	}
}

func TestFlag_Int32Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]int32{42})},
			expected{defaultValue: []int32{42}},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt32Slice(tt.name, tt.opts...)
			assertFlag[[]int32](t, f, tt)
		})
	}
}

func TestFlag_GetInt32Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]int32]{
		{
			"sample",
			[]Option{Separator("|")},
			ptrTo("42|1"),
			result[[]int32]{
				ptrTo([]int32{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("42,1"),
			result[[]int32]{
				ptrTo([]int32{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]int32{42, 1})},
			nil,
			result[[]int32]{
				ptrTo([]int32{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[[]int32]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt32Slice(tt.name, tt.opts...)
			assertGetFlag[[]int32](t, f, tt)
		})
	}
}

func TestFlag_Int64Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]int64{42})},
			expected{defaultValue: []int64{42}},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt64Slice(tt.name, tt.opts...)
			assertFlag[[]int64](t, f, tt)
		})
	}
}

func TestFlag_GetInt64Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]int64]{
		{
			"sample",
			[]Option{Separator("|")},
			ptrTo("42|1"),
			result[[]int64]{
				ptrTo([]int64{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("42,1"),
			result[[]int64]{
				ptrTo([]int64{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]int64{42, 1})},
			nil,
			result[[]int64]{
				ptrTo([]int64{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[[]int64]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewInt64Slice(tt.name, tt.opts...)
			assertGetFlag[[]int64](t, f, tt)
		})
	}
}

func TestFlag_UintSlice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]uint{42})},
			expected{defaultValue: []uint{42}},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUintSlice(tt.name, tt.opts...)
			assertFlag[[]uint](t, f, tt)
		})
	}
}

func TestFlag_GetUintSlice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]uint]{
		{
			"sample",
			[]Option{Separator("|")},
			ptrTo("42|1"),
			result[[]uint]{
				ptrTo([]uint{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("42,1"),
			result[[]uint]{
				ptrTo([]uint{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]uint{42, 1})},
			nil,
			result[[]uint]{
				ptrTo([]uint{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[[]uint]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUintSlice(tt.name, tt.opts...)
			assertGetFlag[[]uint](t, f, tt)
		})
	}
}

func TestFlag_Uint8Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]uint8{42})},
			expected{defaultValue: []uint8{42}},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint8Slice(tt.name, tt.opts...)
			assertFlag[[]uint8](t, f, tt)
		})
	}
}

func TestFlag_GetUint8Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]uint8]{
		{
			"sample",
			[]Option{Separator("|")},
			ptrTo("42|1"),
			result[[]uint8]{
				ptrTo([]uint8{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("42,1"),
			result[[]uint8]{
				ptrTo([]uint8{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]uint8{42, 1})},
			nil,
			result[[]uint8]{
				ptrTo([]uint8{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[[]uint8]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint8Slice(tt.name, tt.opts...)
			assertGetFlag[[]uint8](t, f, tt)
		})
	}
}

func TestFlag_Uint16Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]uint16{42})},
			expected{defaultValue: []uint16{42}},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint16Slice(tt.name, tt.opts...)
			assertFlag[[]uint16](t, f, tt)
		})
	}
}

func TestFlag_GetUint16Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]uint16]{
		{
			"sample",
			[]Option{Separator("|")},
			ptrTo("42|1"),
			result[[]uint16]{
				ptrTo([]uint16{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("42,1"),
			result[[]uint16]{
				ptrTo([]uint16{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]uint16{42, 1})},
			nil,
			result[[]uint16]{
				ptrTo([]uint16{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[[]uint16]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint16Slice(tt.name, tt.opts...)
			assertGetFlag[[]uint16](t, f, tt)
		})
	}
}

func TestFlag_Uint32Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]uint32{42})},
			expected{defaultValue: []uint32{42}},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint32Slice(tt.name, tt.opts...)
			assertFlag[[]uint32](t, f, tt)
		})
	}
}

func TestFlag_GetUint32Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]uint32]{
		{
			"sample",
			[]Option{Separator("|")},
			ptrTo("42|1"),
			result[[]uint32]{
				ptrTo([]uint32{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("42,1"),
			result[[]uint32]{
				ptrTo([]uint32{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]uint32{42, 1})},
			nil,
			result[[]uint32]{
				ptrTo([]uint32{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[[]uint32]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint32Slice(tt.name, tt.opts...)
			assertGetFlag[[]uint32](t, f, tt)
		})
	}
}

func TestFlag_Uint64Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]uint64{42})},
			expected{defaultValue: []uint64{42}},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint64Slice(tt.name, tt.opts...)
			assertFlag[[]uint64](t, f, tt)
		})
	}
}

func TestFlag_GetUint64Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]uint64]{
		{
			"sample",
			[]Option{Separator("|")},
			ptrTo("42|1"),
			result[[]uint64]{
				ptrTo([]uint64{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("42,1"),
			result[[]uint64]{
				ptrTo([]uint64{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]uint64{42, 1})},
			nil,
			result[[]uint64]{
				ptrTo([]uint64{42, 1}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[[]uint64]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewUint64Slice(tt.name, tt.opts...)
			assertGetFlag[[]uint64](t, f, tt)
		})
	}
}

func TestFlag_Float32Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]float32{42.0})},
			expected{defaultValue: []float32{42.0}},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewFloat32Slice(tt.name, tt.opts...)
			assertFlag[[]float32](t, f, tt)
		})
	}
}

func TestFlag_GetFloat32Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]float32]{
		{
			"sample",
			[]Option{Separator("|")},
			ptrTo("42.0|2.7|3.14"),
			result[[]float32]{
				ptrTo([]float32{42.0, 2.7, 3.14}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("42.0,2.7"),
			result[[]float32]{
				ptrTo([]float32{42.0, 2.7}),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]float32{42.0, 2.7})},
			nil,
			result[[]float32]{
				ptrTo([]float32{42.0, 2.7}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[[]float32]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewFloat32Slice(tt.name, tt.opts...)
			assertGetFlag[[]float32](t, f, tt)
		})
	}
}

func TestFlag_Float64Slice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]float64{42.0})},
			expected{defaultValue: []float64{42.0}},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewFloat64Slice(tt.name, tt.opts...)
			assertFlag[[]float64](t, f, tt)
		})
	}
}

func TestFlag_GetFloat64Slice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]float64]{
		{
			"sample",
			[]Option{Separator("|")},
			ptrTo("42.0|2.7|3.14"),
			result[[]float64]{
				ptrTo([]float64{42.0, 2.7, 3.14}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("42.0,2.7"),
			result[[]float64]{
				ptrTo([]float64{42.0, 2.7}),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]float64{42.0, 2.7})},
			nil,
			result[[]float64]{
				ptrTo([]float64{42.0, 2.7}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[[]float64]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewFloat64Slice(tt.name, tt.opts...)
			assertGetFlag[[]float64](t, f, tt)
		})
	}
}

func TestFlag_StringSlice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]string{"foo", "bar"})},
			expected{defaultValue: []string{"foo", "bar"}},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewStringSlice(tt.name, tt.opts...)
			assertFlag[[]string](t, f, tt)
		})
	}
}

func TestFlag_GetStringSlice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]string]{
		{
			"sample",
			[]Option{Separator("|")},
			ptrTo("foo|bar"),
			result[[]string]{
				ptrTo([]string{"foo", "bar"}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("42,1"),
			result[[]string]{
				ptrTo([]string{"42", "1"}),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]string{"foo", "bar"})},
			nil,
			result[[]string]{
				ptrTo([]string{"foo", "bar"}),
				false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewStringSlice(tt.name, tt.opts...)
			assertGetFlag[[]string](t, f, tt)
		})
	}
}

func TestFlag_DurationSlice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]time.Duration{time.Minute})},
			expected{defaultValue: []time.Duration{time.Minute}},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewDurationSlice(tt.name, tt.opts...)
			assertFlag[[]time.Duration](t, f, tt)
		})
	}
}

func TestFlag_GetDurationSlice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]time.Duration]{
		{
			"sample",
			[]Option{Separator("|")},
			ptrTo("42s|1m"),
			result[[]time.Duration]{
				ptrTo([]time.Duration{time.Second * 42, time.Minute}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("42s,1m"),
			result[[]time.Duration]{
				ptrTo([]time.Duration{time.Second * 42, time.Minute}),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]time.Duration{time.Second * 42})},
			nil,
			result[[]time.Duration]{
				ptrTo([]time.Duration{time.Second * 42}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("invalid"),
			result[[]time.Duration]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewDurationSlice(tt.name, tt.opts...)
			assertGetFlag[[]time.Duration](t, f, tt)
		})
	}
}

func TestFlag_BoolSlice(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{Separator("|")},
			expected{
				separator: "|",
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]bool{true, false})},
			expected{defaultValue: []bool{true, false}},
		},
		{
			"sample",
			[]Option{
				Description("description"),
			},
			expected{
				description: "description",
				required:    false,
				shared:      false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewBoolSlice(tt.name, tt.opts...)
			assertFlag[[]bool](t, f, tt)
		})
	}
}

func TestFlag_GetBoolSlice(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[[]bool]{
		{
			"sample",
			[]Option{Separator("|")},
			ptrTo("t|f"),
			result[[]bool]{
				ptrTo([]bool{true, false}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("F,true"),
			result[[]bool]{
				ptrTo([]bool{false, true}),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue([]bool{false, false})},
			nil,
			result[[]bool]{
				ptrTo([]bool{false, false}),
				false,
			},
		},
		{
			"sample",
			[]Option{},
			ptrTo("no,no"),
			result[[]bool]{
				nil,
				true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewBoolSlice(tt.name, tt.opts...)
			assertGetFlag[[]bool](t, f, tt)
		})
	}
}

func TestFlag_Counter(t *testing.T) {
	t.Parallel()
	tests := []flagTest{
		{
			"sample",
			[]Option{},
			expected{},
		},
		{
			"sample",
			[]Option{DefaultValue(5)},
			expected{defaultValue: 5},
		},
		{
			"sample",
			[]Option{
				Description("description"),
				Shorthand("s"),
				Shared(),
			},
			expected{
				description: "description",
				shorthand:   "s",
				required:    true,
				shared:      true,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewCounter(tt.name, tt.opts...)
			assertFlag[int](t, f, tt)
		})
	}
}

func TestFlag_GetCounter(t *testing.T) {
	t.Parallel()
	tests := []getFlagTest[int]{
		{
			"sample",
			[]Option{},
			ptrTo(""),
			result[int]{
				ptrTo(1),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(1)},
			ptrTo(""),
			result[int]{
				ptrTo(2),
				false,
			},
		},
		{
			"sample",
			[]Option{DefaultValue(41)},
			ptrTo(""),
			result[int]{
				ptrTo(42),
				false,
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := NewCounter(tt.name, tt.opts...)
			assertGetFlag[int](t, f, tt)
		})
	}
}
