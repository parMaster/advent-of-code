package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PartOne(t *testing.T) {
	require.Equal(t, 15, PartOne("input0.txt"))
	require.Equal(t, 15632, PartOne("input1.txt"))
}

func Test_PartTwo(t *testing.T) {
	require.Equal(t, 12, PartTwo("input0.txt"))
	require.Equal(t, 14416, PartTwo("input1.txt"))
}
