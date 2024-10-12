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
	prefix             string
	capitalize         bool
	varNameReplacement replacement
}

func Constructor(opts ...Option) *VarNameConstructor {
	c := &VarNameConstructor{}
	for _, opt := range opts {
		opt.apply(c)
	}
	if c.varNameReplacement == (replacement{}) {
		c.varNameReplacement = replacement{old: "", new: ""}
	}
	return c
}

func (c *VarNameConstructor) setPrefix(p string) {
	c.prefix = p
}

func (c *VarNameConstructor) setCapitalize() {
	c.capitalize = true
}

func (c *VarNameConstructor) setReplacement(oldChar, newChar string) {
	c.varNameReplacement = replacement{old: oldChar, new: newChar}
}

func (c *VarNameConstructor) VarFromFlagName(name string) string {
	varName := name
	if c.prefix != "" {
		varName = fmt.Sprintf("%s_%s", c.prefix, varName)
	}
	varName = strings.ReplaceAll(varName, c.varNameReplacement.old, c.varNameReplacement.new)
	if c.capitalize {
		varName = strings.ToUpper(varName)
	}
	return varName
}
