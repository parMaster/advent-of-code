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

var i200 = `
...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`
var i201 = `
..........
.S------7.
.|F----7|.
.||OOOO||.
.||OOOO||.
.|L-7F-J|.
.|II||II|.
.L--JL--J.
..........
`

var i202 = `
.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
`

func TestFuncs(t *testing.T) {
	m, _ := read(i201)
	w, h := maxPoint(m)
	require.Equal(t, 9, w)
	require.Equal(t, 8, h)

	m, _ = read(i1)

	require.True(t, connects(m, image.Pt(1, 1), image.Pt(1, 0)), "S ↑ L is technically connected")
	require.False(t, connects(m, image.Pt(1, 0), image.Pt(1, 1)), "L ↓ S is not connected")
}

func Test_p1(t *testing.T) {
	require.Equal(t, 4, p1(i0))
	require.Equal(t, 4, p1(i1))
	require.Equal(t, 8, p1(i2))

	in, _ := os.ReadFile("../aoc-inputs/2023/10/input.txt")
	require.Equal(t, 7107, p1(string(in)))
}

func Test_p2(t *testing.T) {

	require.Equal(t, 4, p2(i200))
	require.Equal(t, 4, p2(i201))
	require.Equal(t, 8, p2(i202))

	in, _ := os.ReadFile("../aoc-inputs/2023/10/input.txt")
	require.Equal(t, 281, p2(string(in)))
}
