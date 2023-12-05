package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Solution(t *testing.T) {
	one, two := Solve("../aoc-inputs/2022/07/input0.txt")
	require.Equal(t, 95437, one)
	require.Equal(t, 24933642, two)

	one, two = Solve("../aoc-inputs/2022/07/input1.txt")
	require.Equal(t, 1367870, one)
	require.Equal(t, 549173, two)
}
