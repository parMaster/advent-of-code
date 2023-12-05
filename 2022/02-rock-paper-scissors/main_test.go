package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PartOne(t *testing.T) {
	require.Equal(t, 15, PartOne("../aoc-inputs/2022/02/input0.txt"))
	require.Equal(t, 15632, PartOne("../aoc-inputs/2022/02/input1.txt"))
}

func Test_PartTwo(t *testing.T) {
	require.Equal(t, 12, PartTwo("../aoc-inputs/2022/02/input0.txt"))
	require.Equal(t, 14416, PartTwo("../aoc-inputs/2022/02/input1.txt"))
}
