package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PartOne(t *testing.T) {
	require.Equal(t, 157, PartOne("input0.txt"))
	require.Equal(t, 8202, PartOne("input1.txt"))
}

func Test_PartTwo(t *testing.T) {
	require.Equal(t, 70, PartTwo("input0.txt"))
	require.Equal(t, 2864, PartTwo("input1.txt"))
}
