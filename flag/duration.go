package flag

import "time"

type duration = flag[time.Duration]

type Duration struct {
	*duration
}

func (f *Duration) Parse(input string) error {
	v, err := time.ParseDuration(input)
	if err != nil {
		return err
	}
	f.value = &v
	f.visited = true
	return nil
}

func NewDuration(name string, opts ...Option) *Duration {
	f := newFlag[time.Duration](name)
	applyForFlag(f, opts...)
	return &Duration{f}
}
