package flag

import (
	"strconv"
)

type counter = flag[int]

type Counter struct {
	*counter
}

type counterParser struct {
	*embeddedParser
}

func defaultCounterParser() *counterParser {
	return &counterParser{&embeddedParser{}}
}

func (p *counterParser) ParseCmd(input string) (any, error) {
	var (
		empty  string
		parsed int
		err    error
	)
	current := DerefOrDie[int](p.CurrentValue())
	if input == empty {
		parsed = current + 1
	} else {
		parsed, err = strconv.Atoi(input)
		if err != nil {
			return nil, err
		}
	}
	return &parsed, nil
}

func (p *counterParser) ParseEnv(input string) (any, error) {
	var (
		parsed int
		err    error
	)
	if parsed, err = strconv.Atoi(input); err != nil {
		return nil, err
	}
	return &parsed, nil
}

func NewCounter(name string, opts ...Option) *Counter {
	f := newFlag[int](name)
	applyForFlag(f, opts...)
	if f.defaultValue == nil {
		f.setDefaultValue(0)
	}
	f.value = f.defaultValue
	if f.Parser() == nil {
		f.setParser(defaultCounterParser())
	}
	return &Counter{f}
}
