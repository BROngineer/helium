package flagset

import (
	"github.com/brongineer/helium/env"
)

type Builder struct {
	fs *FlagSet
}

func New(opts ...env.Option) *Builder {
	envConstructor := env.Constructor(opts...)
	return &Builder{
		fs: &FlagSet{envVarBinder: envConstructor},
	}
}

func (b *Builder) BindFlag(f flagItem) *Builder {
	b.fs.addFlag(f)
	return b
}

func (b *Builder) Build() *FlagSet {
	return b.fs
}
