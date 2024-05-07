package generic

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const defaultSliceSeparator = ","

type Counter struct {
	count int
}

func NewCounter(v int) Counter {
	return Counter{v}
}

func (c *Counter) update() {
	c.count = c.count + 1
}

func (c *Counter) Count() int {
	return c.count
}

type intFlag interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type uintFlag interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type floatFlag interface {
	~float32 | ~float64
}

type intSliceFlag interface {
	~[]int | ~[]int8 | ~[]int16 | ~[]int32 | ~[]int64
}

type uintSliceFlag interface {
	~[]uint | ~[]uint8 | ~[]uint16 | ~[]uint32 | ~[]uint64
}

type floatSliceFlag interface {
	~[]float32 | ~[]float64
}

type sliceFlag interface {
	~[]string | ~[]bool | ~[]time.Duration | intSliceFlag | uintSliceFlag | floatSliceFlag
}

type Allowed interface {
	~string | ~bool | intFlag | uintFlag | floatFlag | sliceFlag | time.Duration | Counter
}

type Flag[T Allowed] struct {
	name         string
	description  string
	shorthand    string
	required     bool
	shared       bool
	visited      bool
	defaultValue *T
	value        *T
	separator    string
}

func (f *Flag[T]) Value() any {
	if f.value == nil {
		return f.defaultValue
	}
	return f.value
}

func (f *Flag[T]) Name() string {
	return f.name
}

func (f *Flag[T]) Description() string {
	return f.description
}

func (f *Flag[T]) Shorthand() string {
	return f.shorthand
}

func (f *Flag[T]) Separator() string {
	if f.separator == "" {
		return defaultSliceSeparator
	}
	return f.separator
}

func (f *Flag[T]) IsRequired() bool {
	return f.required
}

func (f *Flag[T]) IsShared() bool {
	return f.shared
}

func (f *Flag[T]) IsVisited() bool {
	return f.visited
}

func NewFlag[T Allowed](name string) *Flag[T] {
	return &Flag[T]{name: name}
}

func (f *Flag[T]) SetDescription(description string) {
	f.description = description
}

func (f *Flag[T]) SetShorthand(shorthand string) {
	f.shorthand = shorthand
}

func (f *Flag[T]) SetRequired() {
	f.required = true
}

func (f *Flag[T]) SetShared() {
	f.shared = true
}

func (f *Flag[T]) SetDefaultValue(value any) {
	var v T
	switch any(v).(type) {
	case Counter:
		vv, ok := value.(int)
		if !ok {
			_, _ = fmt.Fprintf(os.Stderr, "counter must be int, got %T", value)
			os.Exit(1)
		}
		v = any(Counter{vv}).(T)
	case T:
		v = value.(T)
	default:
		_, _ = fmt.Fprintf(os.Stderr, "Error: wrong type for default value: got %T, wanted %T\n", value, v)
		os.Exit(1)
	}
	f.defaultValue = &v
}

func (f *Flag[T]) SetSeparator(separator string) {
	f.separator = separator
}

func (f *Flag[T]) Parse(input string) error {
	var (
		t any
		v T
	)
	switch any(f.value).(type) {
	case *string:
		v, _ = any(input).(T)
	case *int:
		d, err := strconv.Atoi(input)
		if err != nil {
			return err
		}
		v = any(d).(T)
	case *int8:
		d, err := strconv.ParseInt(input, 10, 8)
		if err != nil {
			return err
		}
		t = int8(d)
		v = t.(T)
	case *int16:
		d, err := strconv.ParseInt(input, 10, 16)
		if err != nil {
			return err
		}
		t = int16(d)
		v = t.(T)
	case *int32:
		d, err := strconv.ParseInt(input, 10, 32)
		if err != nil {
			return err
		}
		t = int32(d)
		v = t.(T)
	case *int64:
		d, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return err
		}
		v = any(d).(T)
	case *uint:
		d, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			return err
		}
		t = uint(d)
		v = t.(T)
	case *uint8:
		d, err := strconv.ParseUint(input, 10, 8)
		if err != nil {
			return err
		}
		t = uint8(d)
		v = t.(T)
	case *uint16:
		d, err := strconv.ParseUint(input, 10, 16)
		if err != nil {
			return err
		}
		t = uint16(d)
		v = t.(T)
	case *uint32:
		d, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			return err
		}
		t = uint32(d)
		v = t.(T)
	case *uint64:
		d, err := strconv.ParseUint(input, 10, 64)
		if err != nil {
			return err
		}
		v = any(d).(T)
	case *float32:
		d, err := strconv.ParseFloat(input, 32)
		if err != nil {
			return err
		}
		t = float32(d)
		v = t.(T)
	case *float64:
		d, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return err
		}
		v, _ = any(d).(T)
	case *bool:
		var (
			d   bool
			err error
		)
		if input == "" {
			d = true
		} else {
			d, err = strconv.ParseBool(input)
			if err != nil {
				return err
			}
		}
		v, _ = any(d).(T)
	case *time.Duration:
		d, err := time.ParseDuration(input)
		if err != nil {
			return err
		}
		v, _ = any(d).(T)
	case *[]string:
		d := strings.Split(input, f.Separator())
		v, _ = any(d).(T)
	case *[]int:
		d := strings.Split(input, f.Separator())
		s := make([]int, len(d))
		for i, el := range d {
			r, err := strconv.Atoi(el)
			if err != nil {
				return err
			}
			s[i] = r
		}
		t = s
		v = t.(T)
	case *[]int8:
		d := strings.Split(input, f.Separator())
		s := make([]int8, len(d))
		for i, el := range d {
			r, err := strconv.ParseInt(el, 10, 8)
			if err != nil {
				return err
			}
			s[i] = int8(r)
		}
		t = s
		v = t.(T)
	case *[]int16:
		d := strings.Split(input, f.Separator())
		s := make([]int16, len(d))
		for i, el := range d {
			r, err := strconv.ParseInt(el, 10, 16)
			if err != nil {
				return err
			}
			s[i] = int16(r)
		}
		t = s
		v = t.(T)
	case *[]int32:
		d := strings.Split(input, f.Separator())
		s := make([]int32, len(d))
		for i, el := range d {
			r, err := strconv.ParseInt(el, 10, 32)
			if err != nil {
				return err
			}
			s[i] = int32(r)
		}
		t = s
		v = t.(T)
	case *[]int64:
		d := strings.Split(input, f.Separator())
		s := make([]int64, len(d))
		for i, el := range d {
			r, err := strconv.ParseInt(el, 10, 64)
			if err != nil {
				return err
			}
			s[i] = r
		}
		t = s
		v = t.(T)
	case *[]uint:
		d := strings.Split(input, f.Separator())
		s := make([]uint, len(d))
		for i, el := range d {
			r, err := strconv.ParseUint(el, 10, 32)
			if err != nil {
				return err
			}
			s[i] = uint(r)
		}
		t = s
		v = t.(T)
	case *[]uint8:
		d := strings.Split(input, f.Separator())
		s := make([]uint8, len(d))
		for i, el := range d {
			r, err := strconv.ParseUint(el, 10, 8)
			if err != nil {
				return err
			}
			s[i] = uint8(r)
		}
		t = s
		v = t.(T)
	case *[]uint16:
		d := strings.Split(input, f.Separator())
		s := make([]uint16, len(d))
		for i, el := range d {
			r, err := strconv.ParseUint(el, 10, 16)
			if err != nil {
				return err
			}
			s[i] = uint16(r)
		}
		t = s
		v = t.(T)
	case *[]uint32:
		d := strings.Split(input, f.Separator())
		s := make([]uint32, len(d))
		for i, el := range d {
			r, err := strconv.ParseUint(el, 10, 32)
			if err != nil {
				return err
			}
			s[i] = uint32(r)
		}
		t = s
		v = t.(T)
	case *[]uint64:
		d := strings.Split(input, f.Separator())
		s := make([]uint64, len(d))
		for i, el := range d {
			r, err := strconv.ParseUint(el, 10, 64)
			if err != nil {
				return err
			}
			s[i] = r
		}
		t = s
		v = t.(T)
	case *[]float32:
		d := strings.Split(input, f.Separator())
		s := make([]float32, len(d))
		for i, el := range d {
			r, err := strconv.ParseFloat(el, 32)
			if err != nil {
				return err
			}
			s[i] = float32(r)
		}
		t = s
		v = t.(T)
	case *[]float64:
		d := strings.Split(input, f.Separator())
		s := make([]float64, len(d))
		for i, el := range d {
			r, err := strconv.ParseFloat(el, 64)
			if err != nil {
				return err
			}
			s[i] = r
		}
		t = s
		v = t.(T)
	case *[]time.Duration:
		d := strings.Split(input, f.Separator())
		s := make([]time.Duration, len(d))
		for i, el := range d {
			r, err := time.ParseDuration(el)
			if err != nil {
				return err
			}
			s[i] = r
		}
		t = s
		v = t.(T)
	case *[]bool:
		d := strings.Split(input, f.Separator())
		s := make([]bool, len(d))
		for i, el := range d {
			r, err := strconv.ParseBool(el)
			if err != nil {
				return err
			}
			s[i] = r
		}
		t = s
		v = t.(T)
	case *Counter:
		var (
			d = 1
		)
		val := f.Value().(*Counter)
		if val == nil {
			t = Counter{d}
		} else {
			c := val
			c.update()
			t = *c
		}
		v = t.(T)
	}
	f.value = &v
	f.visited = true
	return nil
}
