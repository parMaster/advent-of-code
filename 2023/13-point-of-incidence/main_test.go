package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReflected(t *testing.T) {
	testCases := []struct {
		name     string
		lines    []string
		expected bool
	}{
		{
			name:     "Empty lines",
			lines:    []string{},
			expected: false,
		},
		{
			name:     "Odd number of lines",
			lines:    []string{"", "", ""},
			expected: false,
		},
		{
			name:     "Even number of lines",
			lines:    []string{"", "", "", ""},
			expected: true,
		},
		{
			name:     "Reflected lines",
			lines:    []string{"...", "###", "###", "..."},
			expected: true,
		},
		{
			name:     "Not reflected lines",
			lines:    []string{".#.", "###", "###", "..."},
			expected: false,
		},
		{
			name:     "not reflected",
			lines:    []string{"#.##..##.", "..#.##.#."},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := reflected(tc.lines, 0)
			require.Equal(t, tc.expected, actual, tc.name)
		})
	}

	// one smudge, which is allowed
	require.True(t, reflected([]string{".#.", "###", "###", "..."}, 1))
	// not two
	require.False(t, reflected([]string{".#.", "###", "###", "..."}, 2))
}
