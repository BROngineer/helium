package flag

import (
	"strings"
	"time"

	"github.com/brongineer/helium/errors"
)

type durationSlice = flag[[]time.Duration]

type DurationSlice struct {
	*durationSlice
}

type durationSliceParser struct {
	*embeddedParser
}

func defaultDurationSliceParser() *durationSliceParser {
	return &durationSliceParser{&embeddedParser{}}
}

func (p *durationSliceParser) ParseCmd(input string) (any, error) {
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	s := strings.Split(input, p.Separator())
	parsed := make([]time.Duration, 0, len(s))
	for _, el := range s {
		v, err := time.ParseDuration(el)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, v)
	}
	if p.IsVisited() {
		stored := DerefOrDie[[]time.Duration](p.CurrentValue())
		parsed = append(stored, parsed...)
	}
	return &parsed, nil
}

func (p *durationSliceParser) ParseEnv(input string) (any, error) {
	s := strings.Split(input, p.Separator())
	parsed := make([]time.Duration, 0, len(s))
	for _, el := range s {
		v, err := time.ParseDuration(el)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, v)
	}
	return &parsed, nil
}

func NewDurationSlice(name string, opts ...Option) *DurationSlice {
	f := newFlag[[]time.Duration](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultDurationSliceParser())
	}
	return &DurationSlice{f}
}
