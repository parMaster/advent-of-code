package main

import (
	"image"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var i0 = `
.....
.S-7.
.|.|.
.L-J.
.....
`

var i1 = `
-L|F7
7S-7|
L|7||
-L-J|
L|-JF
`

var i2 = `
..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`

func Test_Solve(t *testing.T) {

	m, _, _ := read(i1)

	require.True(t, connects(m, image.Pt(1, 1), image.Pt(1, 0)), "S ↑ L is technically connected")
	require.False(t, connects(m, image.Pt(1, 0), image.Pt(1, 1)), "L ↓ S is not connected")

	require.Equal(t, 4, solve(i0))
	require.Equal(t, 4, solve(i1))
	require.Equal(t, 8, solve(i2))

	in, _ := os.ReadFile("../aoc-inputs/2023/10/input.txt")
	require.Equal(t, 7107, solve(string(in)))

}
