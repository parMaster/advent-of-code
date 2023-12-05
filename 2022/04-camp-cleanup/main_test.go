package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PartOne(t *testing.T) {
	require.Equal(t, 2, PartOne("../aoc-inputs/2022/04/input0.txt"))
	require.Equal(t, 431, PartOne("../aoc-inputs/2022/04/input1.txt"))
}

func Test_PartTwo(t *testing.T) {
	require.Equal(t, 4, PartTwo("../aoc-inputs/2022/04/input0.txt"))
	require.Equal(t, 823, PartTwo("../aoc-inputs/2022/04/input1.txt"))
}
