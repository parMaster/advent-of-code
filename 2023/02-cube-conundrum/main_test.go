package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Both(t *testing.T) {

	require.Equal(t, 8, PartOne("../aoc-inputs/2023/02/input0.txt", 12, 13, 14))
	require.Equal(t, 3059, PartOne("../aoc-inputs/2023/02/input1.txt", 12, 13, 14))

	require.Equal(t, 2286, PartTwo("../aoc-inputs/2023/02/input0.txt"))
	require.Equal(t, 65371, PartTwo("../aoc-inputs/2023/02/input1.txt"))
}
