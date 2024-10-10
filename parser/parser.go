package parser

import "github.com/brongineer/helium/errors"

type FlagParser interface {
	SetFromEnv(bool)
	SetFromCmd(bool)
	SetSeparator(string)
	SetCurrentValue(any)
	IsSetFromEnv() bool
	IsSetFromCmd() bool
	Separator() string
	CurrentValue() any
	ParseCmd(string) (any, error)
	ParseEnv(string) (any, error)
}

type EmbeddedParser struct {
	fromEnv      bool
	fromCmd      bool
	separator    string
	currentValue any
}

func (p *EmbeddedParser) SetFromEnv(v bool) {
	p.fromEnv = v
}

func (p *EmbeddedParser) IsSetFromEnv() bool {
	return p.fromEnv
}

func (p *EmbeddedParser) SetFromCmd(v bool) {
	p.fromCmd = v
}

func (p *EmbeddedParser) IsSetFromCmd() bool {
	return p.fromCmd
}

func (p *EmbeddedParser) SetSeparator(s string) {
	p.separator = s
}

func (p *EmbeddedParser) Separator() string {
	return p.separator
}

func (p *EmbeddedParser) SetCurrentValue(v any) {
	p.currentValue = v
}

func (p *EmbeddedParser) CurrentValue() any {
	return p.currentValue
}

func (p *EmbeddedParser) ParseCmd(_ string) (any, error) {
	return nil, errors.ErrCmdParserIsNotImplemented
}

func (p *EmbeddedParser) ParseEnv(_ string) (any, error) {
	return nil, errors.ErrEnvParserIsNotImplemented
}
