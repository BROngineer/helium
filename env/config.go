package env

import (
	"fmt"
	"strings"
)

type replacement struct {
	old string
	new string
}

type VarNameConstructor struct {
	prefix      string
	capitalize  bool
	replacement replacement
}

func Constructor(charOld, charNew string, opts ...Option) *VarNameConstructor {
	c := &VarNameConstructor{replacement: replacement{old: charOld, new: charNew}}
	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}

func (c *VarNameConstructor) setPrefix(p string) {
	c.prefix = p
}

func (c *VarNameConstructor) setCapitalize() {
	c.capitalize = true
}

func (c *VarNameConstructor) VarFromFlagName(name string) string {
	varName := strings.ReplaceAll(name, c.replacement.old, c.replacement.new)
	if c.prefix != "" {
		varName = fmt.Sprintf("%s_%s", c.prefix, varName)
	}
	if c.capitalize {
		varName = strings.ToUpper(varName)
	}
	return varName
}
