package python_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	python3 "go.nhat.io/python/v3"
)

func TestTypeName(t *testing.T) {
	testCases := []struct {
		scenario string
		value    *python3.Object
		expected string
	}{
		{
			scenario: "bool",
			value:    python3.NewBool(true),
			expected: "bool",
		},
		{
			scenario: "string",
			value:    python3.NewString("hello"),
			expected: "str",
		},
		{
			scenario: "int",
			value:    python3.NewInt(42),
			expected: "int",
		},
		{
			scenario: "int64",
			value:    python3.NewInt64(42),
			expected: "int",
		},
		{
			scenario: "float",
			value:    python3.NewFloat64(3.14),
			expected: "float",
		},
		{
			scenario: "list",
			value:    python3.NewList(1).AsObject(),
			expected: "list",
		},
		{
			scenario: "tuple",
			value:    python3.NewTuple(1).AsObject(),
			expected: "tuple",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			actual := python3.TypeName(tc.value)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
