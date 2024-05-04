package helium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlagSet_Bool(t *testing.T) {
	t.Parallel()

	fs := FlagSet{}
	fs.Bool("sample")
	assert.NotNil(t, fs.flagByName("sample"))
}

func TestFlagSet_GetBool(t *testing.T) {
	t.Parallel()

	fs := FlagSet{}
	fs.Bool("sample")
	f := fs.flagByName("sample")
	_ = f.Parse("true")
	assert.Equal(t, fs.GetBool("sample"), true)
}

func TestFlagSet_Counter(t *testing.T) {
	t.Parallel()

	fs := FlagSet{}
	fs.Counter("sample")
	assert.NotNil(t, fs.flagByName("sample"))
}

func TestFlagSet_GetCounter(t *testing.T) {
	t.Parallel()

	fs := FlagSet{}
	fs.Counter("sample")
	f := fs.flagByName("sample")
	_ = f.Parse("")
	assert.Equal(t, fs.GetCounter("sample"), uint64(1))
	_ = f.Parse("")
	assert.Equal(t, fs.GetCounter("sample"), uint64(2))
	_ = f.Parse("")
	assert.Equal(t, fs.GetCounter("sample"), uint64(3))
}
