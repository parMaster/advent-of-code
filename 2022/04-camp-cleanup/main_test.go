package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PartOne(t *testing.T) {
	require.Equal(t, 2, PartOne("input0.txt"))
	require.Equal(t, 431, PartOne("input1.txt"))
}

func Test_PartTwo(t *testing.T) {
	require.Equal(t, 4, PartTwo("input0.txt"))
	require.Equal(t, 823, PartTwo("input1.txt"))
}
