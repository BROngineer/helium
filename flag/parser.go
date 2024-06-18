package flag

type embeddedParser struct {
	visited      bool
	separator    string
	currentValue any
}

func (p *embeddedParser) SetVisited(v bool) {
	p.visited = v
}

func (p *embeddedParser) IsVisited() bool {
	return p.visited
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
