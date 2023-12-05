package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PartOne(t *testing.T) {

	require.Equal(t, "CMZ", PartOne("../aoc-inputs/2022/05/input0.txt"))
	require.Equal(t, "CNSZFDVLJ", PartOne("../aoc-inputs/2022/05/input1.txt"))

}

func Test_PartTwo(t *testing.T) {

	require.Equal(t, "MCD", PartTwo("../aoc-inputs/2022/05/input0.txt"))
	require.Equal(t, "QNDWLMGNS", PartTwo("../aoc-inputs/2022/05/input1.txt"))

}

func Test_Stack(t *testing.T) {
	s := NewStack()
	s.push("0")
	require.Equal(t, []string{"0"}, s.items)
	s.push("1")
	s.push("2")
	s.push("3")
	require.Equal(t, []string{"0", "1", "2", "3"}, s.items)

	s.pop()
	require.Equal(t, []string{"0", "1", "2"}, s.items)
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	require.Equal(t, []string{}, s.items)
}

func Test_MStacks(t *testing.T) {
	m := NewMStacks(5)
	require.Equal(t, 5, len(m.stacks))

	m.stacks[0].push("23")
}
