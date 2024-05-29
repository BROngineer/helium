package flag

import (
	"fmt"
	"strconv"
)

type fbool = flag[bool]

type Bool struct {
	*fbool
}

func (f *Bool) Parse(input string) error {
	var (
		v   bool
		err error
	)
	if f.IsVisited() {
		return fmt.Errorf("flag already parsed")
	}
	switch input {
	case "":
		v = true
	default:
		v, err = strconv.ParseBool(input)
		if err != nil {
			return err
		}
	}
	f.value = &v
	f.visited = true
	return nil
}

func NewBool(name string, opts ...Option) *Bool {
	f := newFlag[bool](name)
	applyForFlag(f, opts...)
	return &Bool{f}
}
