package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

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

func showMaze(f Maze, slopes bool) {
	asciiBlocks := map[rune]string{'.': "░░", '#': "██"}
	for y := 0; y < len(f); y++ {
		for x := 0; x < len(f[y]); x++ {
			if ab, ok := asciiBlocks[f[y][x]]; ok {
				fmt.Print(ab)
			} else if slopes {
				fmt.Print(string(f[y][x]) + " ")
			} else {
				fmt.Print(asciiBlocks['.'])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func walkSlopes(maze Maze, start [2]int) int {
	maze[0][1] = '#'
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
			return len(visited) + max(walkSlopes(maze, [2]int{legal[0][0], legal[0][1]}), walkSlopes(maze, [2]int{legal[1][0], legal[1][1]}))
		}
	}
}

// forward:backward stepping map
var moves = map[[2]int][2]int{{0, 1}: {0, -1}, {1, 0}: {-1, 0}, {0, -1}: {0, 1}, {-1, 0}: {1, 0}}

type Q struct {
	pos   [2]int
	dir   [2]int
	steps int
}

func MaxPath(maze Maze) {
	maze[0][1] = '#'

	forkedBefore := map[Q]int{}
	maxResult := 0
	results := []int{}
	start := [2]int{1, 1}
	steps := 1
	direction := [2]int{1, 0}
	visited := [][2]int{start}

	q := NewStack(Q{})
	q.Push(Q{pos: start, dir: direction, steps: steps})

	var sx, sy int
	for !q.IsEmpty() {

		curr := q.Pop()
		sx, sy = curr.pos[0], curr.pos[1]
		steps = curr.steps
		direction = curr.dir

		if sx == len(maze[0])-2 && sy == len(maze)-1 {

			if steps > maxResult {
				maxResult = steps
				log.Println(steps, len(visited), len(forkedBefore))
			}
			results = append(results, steps)
			continue
		}

		visited = visited[:steps]
		visited = append(visited, [2]int{sx, sy})

		legal := [][2]int{}
		for move := range moves {
			// try to move and check
			x, y := sx+move[0], sy+move[1]

			// backwards?
			if move == moves[direction] {
				continue
			}

			// was here before?
			if slices.Index(visited, [2]int{x, y}) != -1 {
				continue
			}

			// not in bounds? or a wall ?
			if x < 0 || x >= len(maze[0]) || y < 0 || y >= len(maze) || maze[y][x] == '#' {
				continue
			}

			// if _, ok := forkedBefore[Q{pos: [2]int{x, y}, dir: move, steps: steps + 1}]; ok { // is this memo key valid??????
			// 	continue
			// }

			// otherwise - legal position
			legal = append(legal, [2]int{x, y})
		}
		for i := range len(legal) {
			direction = [2]int{legal[i][0] - sx, legal[i][1] - sy}
			newQ := Q{pos: legal[i], dir: direction, steps: steps + 1}
			// if len(legal) > 1 {
			// 	forkedBefore[Q{pos: legal[i], dir: direction, steps: steps + 1}] = steps + 1
			// }
			q.Push(newQ)
		}
	}
	log.Println("Max path: ", slices.Max(results))
}

func main() {
	maze := readLines("../aoc-inputs/2023/23/input.txt")
	fmt.Println("Day 23: A Long Walk")
	if slices.Contains(os.Args[1:], "--visual") {
		showMaze(maze, true)
	}
	fmt.Println("\tPart One:", walkSlopes(maze, [2]int{1, 1})) // 2386

	if slices.Contains(os.Args[1:], "--bruteforce") {
		fmt.Println("\tPart Two: will print preliminary results every time new max path is found. It's a really long walk...")
		MaxPath(maze) // 6246
	} else {
		fmt.Println("\tPart Two: (skipped by default, run with a '--bruteforce' option and prepare to wait up to forever)")
	}
}
