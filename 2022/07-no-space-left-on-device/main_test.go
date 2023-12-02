package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Solution(t *testing.T) {
	one, two := Solve("input0.txt")
	require.Equal(t, 95437, one)
	require.Equal(t, 24933642, two)

	one, two = Solve("input1.txt")
	require.Equal(t, 1367870, one)
	require.Equal(t, 549173, two)
}
