package main

import (
	"image"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PartOne(t *testing.T) {

	require.True(t, gap(image.Point{0, 0}, image.Point{2, 0}))
	require.True(t, gap(image.Point{0, 0}, image.Point{20, 0}))
	require.False(t, gap(image.Point{0, 0}, image.Point{0, 0}))
	require.False(t, gap(image.Point{0, 0}, image.Point{1, 0}))
	require.False(t, gap(image.Point{0, 0}, image.Point{1, 1}))
	require.True(t, gap(image.Point{0, 0}, image.Point{1, 2}))

	require.Equal(t, 13, PartOne("input0.txt"))
	require.Equal(t, 6522, PartOne("input1.txt"))
}
