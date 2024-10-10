package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	flagName string
	options  []Option
	expected string
}

func TestVarNameConstructor_VarFromFlagName(t *testing.T) {
	t.Parallel()

	tests := []testCase{
		{
			name:     "var name as is",
			flagName: "sample-flag",
			options:  []Option{},
			expected: "sample-flag",
		},
		{
			name:     "var name capitalized",
			flagName: "sample-flag",
			options:  []Option{Capitalized()},
			expected: "SAMPLE-FLAG",
		},
		{
			name:     "var name with replacement",
			flagName: "sample-flag",
			options:  []Option{Capitalized(), VarNameReplace("-", "_")},
			expected: "SAMPLE_FLAG",
		},
		{
			name:     "var name with prefix",
			flagName: "sample-flag",
			options:  []Option{Prefix("test")},
			expected: "test_sample-flag",
		},
		{
			name:     "var name with combined options",
			flagName: "sample-flag",
			options:  []Option{Prefix("test"), Capitalized(), VarNameReplace("-", "_")},
			expected: "TEST_SAMPLE_FLAG",
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := Constructor(tt.options...)
			actual := c.VarFromFlagName(tt.flagName)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
