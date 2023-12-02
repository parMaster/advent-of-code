package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Both(t *testing.T) {

	require.Equal(t, 8, PartOne("input0.txt", 12, 13, 14))
	require.Equal(t, 3059, PartOne("input1.txt", 12, 13, 14))

	require.Equal(t, 2286, PartTwo("input0.txt"))
	require.Equal(t, 65371, PartTwo("input1.txt"))
}
