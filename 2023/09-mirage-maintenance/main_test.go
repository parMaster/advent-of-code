package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPredict(t *testing.T) {
	require.Equal(t, 18, predict([]int{0, 3, 6, 9, 12, 15}))
	require.Equal(t, 28, predict([]int{1, 3, 6, 10, 15, 21}))
	require.Equal(t, 68, predict([]int{10, 13, 16, 21, 30, 45}))
}

func TestSolve(t *testing.T) {
	require.Equal(t, 1921197370, solve("../aoc-inputs/2023/09/input.txt", predict))
	require.Equal(t, 1124, solve("../aoc-inputs/2023/09/input.txt", predict_past))
}
