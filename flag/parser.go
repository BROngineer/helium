package flag

import "github.com/brongineer/helium/errors"

type embeddedParser struct {
	fromEnv      bool
	fromCmd      bool
	separator    string
	currentValue any
}

func (p *embeddedParser) SetFromEnv(v bool) {
	p.fromEnv = v
}

func (p *embeddedParser) IsSetFromEnv() bool {
	return p.fromEnv
}

func (p *embeddedParser) SetFromCmd(v bool) {
	p.fromCmd = v
}

func (p *embeddedParser) IsSetFromCmd() bool {
	return p.fromCmd
}

func (p *embeddedParser) SetSeparator(s string) {
	p.separator = s
}

func (p *embeddedParser) Separator() string {
	return p.separator
}

func (p *embeddedParser) SetCurrentValue(v any) {
	p.currentValue = v
}

func (p *embeddedParser) CurrentValue() any {
	return p.currentValue
}

func (p *embeddedParser) ParseCmd(_ string) (any, error) {
	return nil, errors.ErrCmdParserIsNotImplemented
}

func (p *embeddedParser) ParseEnv(_ string) (any, error) {
	return nil, errors.ErrEnvParserIsNotImplemented
}
