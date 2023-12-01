package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CalibrationValues(t *testing.T) {

	tests := []struct {
		name     string
		inp      string
		expected int
	}{
		{
			name:     "primitive input",
			inp:      "input0.txt",
			expected: 142,
		},
		{
			name:     "real input",
			inp:      "input.txt",
			expected: 53974,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := PartOne(tt.inp)
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}

func Test_Calibration_Spelled(t *testing.T) {
	tests := []struct {
		name     string
		inp      string
		expected int
	}{
		{
			name:     "part 2 - primitive input",
			inp:      "input20.txt",
			expected: 281,
		},
		{
			name:     "part 2 - real input",
			inp:      "input2.txt",
			expected: 52840, // 52736 - 52844
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := PartTwo(tt.inp)
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}
