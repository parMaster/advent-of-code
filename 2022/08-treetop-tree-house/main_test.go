package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Solve(t *testing.T) {
	visible, maxScore := Solve("../aoc-inputs/2022/08/input0.txt")
	require.Equal(t, 21, visible)
	require.Equal(t, 8, maxScore)

	visible, maxScore = Solve("../aoc-inputs/2022/08/input1.txt")
	require.Equal(t, 1789, visible)
	require.Equal(t, 314820, maxScore)
}
