package flag

import (
	"strconv"
	"strings"

	"github.com/brongineer/helium/errors"
)

type boolSlice = flag[[]bool]

type BoolSliceFlag struct {
	*boolSlice
}

type boolSliceParser struct {
	*embeddedParser
}

func defaultBoolSliceParser() *boolSliceParser {
	return &boolSliceParser{&embeddedParser{}}
}

func (p *boolSliceParser) ParseCmd(input string) (any, error) {
	var empty string
	if input == empty {
		return nil, errors.ErrNoValueProvided
	}
	s := strings.Split(input, p.Separator())
	parsed := make([]bool, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseBool(el)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, v)
	}
	if p.IsSetFromCmd() {
		stored := DerefOrDie[[]bool](p.CurrentValue())
		parsed = append(stored, parsed...)
	}
	return &parsed, nil
}

func (p *boolSliceParser) ParseEnv(value string) (any, error) {
	s := strings.Split(value, p.Separator())
	parsed := make([]bool, 0, len(s))
	for _, el := range s {
		v, err := strconv.ParseBool(el)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, v)
	}
	return &parsed, nil
}

func BoolSlice(name string, opts ...Option) *BoolSliceFlag {
	f := newFlag[[]bool](name)
	applyForFlag(f, opts...)
	if f.Parser() == nil {
		f.setParser(defaultBoolSliceParser())
	}
	return &BoolSliceFlag{f}
}
