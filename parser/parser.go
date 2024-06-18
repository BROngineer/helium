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

type DefaultParser struct {
	visited      bool
	separator    string
	currentValue any
}

func (p *DefaultParser) SetVisited(v bool) {
	p.visited = v
}

func (p *DefaultParser) IsVisited() bool {
	return p.visited
}

func (p *DefaultParser) SetSeparator(s string) {
	p.separator = s
}

func (p *DefaultParser) Separator() string {
	return p.separator
}

func (p *DefaultParser) SetCurrentValue(v any) {
	p.currentValue = v
}

func (p *DefaultParser) CurrentValue() any {
	return p.currentValue
}
