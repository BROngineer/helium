package flag

import (
	"strings"
	"time"

	"github.com/brongineer/helium/errors"
)

type durationSlice = flag[[]time.Duration]

type DurationSliceFlag struct {
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
	if p.IsSetFromCmd() {
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

func DurationSlice(name string, opts ...Option) *DurationSliceFlag {
	f := newFlag[[]time.Duration](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultDurationSliceParser())
	}
	return &DurationSliceFlag{f}
}
