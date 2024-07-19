package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type float32Slice = flag[[]float32]

type Float32SliceFlag struct {
	*float32Slice
}

type float32SliceParser struct {
	*embeddedParser
}

func defaultFloat32SliceParser() *float32SliceParser {
	return &float32SliceParser{&embeddedParser{}}
}

func (p *float32SliceParser) ParseCmd(input string) (any, error) {
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	s := strings.Split(input, p.Separator())
	parsed := make([]float32, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseFloat(el, 32)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, float32(v))
	}
	if p.IsSetFromCmd() {
		stored := DerefOrDie[[]float32](p.CurrentValue())
		parsed = append(stored, parsed...)
	}
	return &parsed, nil
}

func (p *float32SliceParser) ParseEnv(input string) (any, error) {
	s := strings.Split(input, p.Separator())
	parsed := make([]float32, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseFloat(el, 32)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, float32(v))
	}
	return &parsed, nil
}

func Float32Slice(name string, opts ...Option) *Float32SliceFlag {
	f := newFlag[[]float32](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultFloat32SliceParser())
	}
	return &Float32SliceFlag{f}
}
