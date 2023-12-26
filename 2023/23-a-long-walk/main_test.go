package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
)

type Maze [][]rune

func readLines(file string) Maze {
	in, _ := os.ReadFile(file)
	f := Maze{}
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		f = append(f, []rune{})
		for _, r := range strings.TrimSpace(l) {
			f[y] = append(f[y], r)
		}

	}
	return f
}

func showLines(f Maze) {
	asciiBlocks := map[rune]string{'.': "░░", '#': "██"}
	for y := 0; y < len(f); y++ {
		for x := 0; x < len(f[y]); x++ {
			if ab, ok := asciiBlocks[f[y][x]]; ok {
				fmt.Print(ab)
			} else {
				fmt.Print(string(f[y][x]) + " ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func walk(maze Maze, start [2]int) int {
	moves := map[rune][2]int{'^': {0, -1}, '<': {-1, 0}, 'v': {0, 1}, '>': {1, 0}}

	sx, sy := start[0], start[1]
	visited := map[[2]int]bool{}
	for {
		visited[[2]int{sx, sy}] = true
		// at the finish already?
		if sx == len(maze[0])-2 && sy == len(maze)-1 {
			return len(visited)
		}
		legal := [][2]int{}
		for slope, move := range moves {

			// current position is a slope? then move only downhill
			if slices.Index(maps.Keys(moves), maze[sy][sx]) != -1 && maze[sy][sx] != slope {
				continue
			}

			// try to move and check
			x, y := sx+move[0], sy+move[1]

			// going back?
			if _, ok := visited[[2]int{x, y}]; ok {
				continue
			}

			// not in bounds?
			if x < 0 || x >= len(maze[0]) || y < 0 || y >= len(maze) ||
				// or a wall ?
				maze[y][x] == '#' {
				continue
			}

			// is it a slope? downhill slope?
			if slices.Index(maps.Keys(moves), maze[y][x]) != -1 && maze[y][x] != slope {
				continue
			}

			// all checked, legal move
			legal = append(legal, [2]int{x, y})
		}

		// no way, not a finish
		if len(legal) == 0 {
			return 0
		}

		// only one way, just move, don't call walk
		if len(legal) == 1 {
			sx, sy = legal[0][0], legal[0][1]
		}

		// fork and choose the longest route
		if len(legal) == 2 {
			return len(visited) + max(walk(maze, [2]int{legal[0][0], legal[0][1]}), walk(maze, [2]int{legal[1][0], legal[1][1]}))
		}
	}
}

func TestRead(t *testing.T) {
	maze := readLines("../aoc-inputs/2023/23/input.txt")
	require.True(t, len(maze) > 0)

	maze[0][1] = '#'
	showLines(maze)
	start := [2]int{1, 1}

	p1 := walk(maze, start)
	log.Println(p1) //2386
}
