package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name       string
	flagName   string
	replaceOld string
	replaceNew string
	options    []Option
	expected   string
}

func TestVarNameConstructor_VarFromFlagName(t *testing.T) {
	t.Parallel()

	tests := []testCase{
		{
			name:       "var name as is",
			flagName:   "sample-flag",
			replaceOld: "",
			replaceNew: "",
			options:    []Option{},
			expected:   "sample-flag",
		},
		{
			name:       "var name capitalized",
			flagName:   "sample-flag",
			replaceOld: "",
			replaceNew: "",
			options:    []Option{Capitalized()},
			expected:   "SAMPLE-FLAG",
		},
		{
			name:       "var name with replacement",
			flagName:   "sample-flag",
			replaceOld: "-",
			replaceNew: "_",
			options:    []Option{Capitalized()},
			expected:   "SAMPLE_FLAG",
		},
		{
			name:       "var name with prefix",
			flagName:   "sample-flag",
			replaceOld: "",
			replaceNew: "",
			options:    []Option{Prefix("test")},
			expected:   "test_sample-flag",
		},
		{
			name:       "var name with combined options",
			flagName:   "sample-flag",
			replaceOld: "-",
			replaceNew: "_",
			options:    []Option{Prefix("test"), Capitalized()},
			expected:   "TEST_SAMPLE_FLAG",
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := Constructor(tt.replaceOld, tt.replaceNew, tt.options...)
			actual := c.VarFromFlagName(tt.flagName)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
