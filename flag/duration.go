package flag

import (
	"time"

	"github.com/brongineer/helium/errors"
)

type duration = flag[time.Duration]

type DurationFlag struct {
	*duration
}

type durationParser struct {
	*embeddedParser
}

func defaultDurationParser() *durationParser {
	return &durationParser{&embeddedParser{}}
}

func (p *durationParser) ParseCmd(input string) (any, error) {
	if p.IsSetFromCmd() {
		return nil, errors.ErrFlagVisited
	}
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	parsed, err := time.ParseDuration(input)
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

func (p *durationParser) ParseEnv(input string) (any, error) {
	var (
		parsed time.Duration
		err    error
	)
	if parsed, err = time.ParseDuration(input); err != nil {
		return nil, err
	}
	return &parsed, err
}

func Duration(name string, opts ...Option) *DurationFlag {
	f := newFlag[time.Duration](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultDurationParser())
	}
	return &DurationFlag{f}
}
