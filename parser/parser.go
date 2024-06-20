package parser

type FlagParser interface {
	SetVisited(bool)
	SetSeparator(string)
	SetCurrentValue(any)
	IsVisited() bool
	Separator() string
	CurrentValue() any
	ParseCmd(string) (any, error)
	ParseEnv(string) (any, error)
}

type EmbeddedParser struct {
	visited      bool
	separator    string
	currentValue any
}

func (p *EmbeddedParser) SetVisited(v bool) {
	p.visited = v
}

func (p *EmbeddedParser) IsVisited() bool {
	return p.visited
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
