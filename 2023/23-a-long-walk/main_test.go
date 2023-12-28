package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestP0(t *testing.T) {
	maze := readLines("../aoc-inputs/2023/23/input0.txt")
	require.True(t, len(maze) > 0)

	t.Skip()
	maze[0][1] = '#'
	showMaze(maze, true)
}

func Test(t *testing.T) {
	legal := [][2]int{{1, 1}, {2, 2}}
	legal = legal[:1]
	require.Equal(t, 1, len(legal))
	require.Equal(t, [2]int{1, 1}, legal[0])
}

func TestP1(t *testing.T) {
	t.Skip()
	maze := readLines("../aoc-inputs/2023/23/input.txt")
	require.True(t, len(maze) > 0)

	maze[0][1] = '#'
	showMaze(maze, true)
	start := [2]int{1, 1}

	p1 := walkSlopes(maze, start)
	log.Println(p1) //2386
	require.Equal(t, 2386, p1)
}

// Basically three versions of failed breadth-first attempts
// demostrating masive combinatorial explosions

// func walkMaze(maze Maze, start [2]int, visited map[[2]int]bool) map[[2]int]bool {
// 	moves := map[rune][2]int{'^': {0, -1}, '<': {-1, 0}, 'v': {0, 1}, '>': {1, 0}}

// 	sx, sy := start[0], start[1]
// 	for {
// 		visited[[2]int{sx, sy}] = true
// 		// at the finish already?
// 		if sx == len(maze[0])-2 && sy == len(maze)-1 {
// 			return visited
// 		}
// 		legal := [][2]int{}
// 		for _, move := range moves {
// 			// try to move and check
// 			x, y := sx+move[0], sy+move[1]

// 			// going back?
// 			if _, ok := visited[[2]int{x, y}]; ok {
// 				continue
// 			}

// 			// not in bounds?
// 			if x < 0 || x >= len(maze[0]) || y < 0 || y >= len(maze) ||
// 				// or a wall ?
// 				maze[y][x] == '#' {
// 				continue
// 			}

// 			// all checked, legal move
// 			legal = append(legal, [2]int{x, y})
// 		}

// 		// no way, not a finish
// 		if len(legal) == 0 {
// 			return map[[2]int]bool{}
// 		}

// 		// fmt.Println(legal)

// 		// only one way, just move, don't call walk
// 		if len(legal) == 1 {
// 			sx, sy = legal[0][0], legal[0][1]
// 		}

// 		// was here with better result?
// 		if prevBest, ok := memo[[2]int{sx, sy}]; ok && prevBest > len(visited) {
// 			return map[[2]int]bool{}
// 		}
// 		memo[[2]int{sx, sy}] = len(visited)

// 		// fork and choose the longest route
// 		if len(legal) == 2 {
// 			left := map[[2]int]bool{}
// 			right := map[[2]int]bool{}

// 			// deep copy maps
// 			for k, v := range visited {
// 				left[[2]int{k[0], k[1]}] = v
// 				right[[2]int{k[0], k[1]}] = v
// 			}

// 			leftVisited := walkMaze(maze, [2]int{legal[0][0], legal[0][1]}, left)
// 			rightVisited := walkMaze(maze, [2]int{legal[1][0], legal[1][1]}, right)
// 			if len(leftVisited) > len(rightVisited) {
// 				return leftVisited
// 			}
// 			return rightVisited
// 		}
// 	}
// }

//
// Another Fail:
//
// func p2() {
// 	maze := readLines("../aoc-inputs/2023/23/input.txt")
// 	// require.True(t, len(maze) > 0)

// 	start := [2]int{1, 1}

// 	q := []Q{{[][2]int{}, start}}
// 	results := []int{}
// 	moves := map[rune][2]int{'^': {0, -1}, '<': {-1, 0}, 'v': {0, 1}, '>': {1, 0}}
// 	maze[0][1] = '#'
// 	// showMaze(maze, false)
// 	// visited := map[[2]int]bool{}
// 	i := 0
// 	for len(q) > 0 {
// 		nq := []Q{}
// 		for _, current := range q {

// 			current.visited = append(current.visited, current.pos)

// 			sx, sy := current.pos[0], current.pos[1]
// 			if sx == len(maze[0])-2 && sy == len(maze)-1 {
// 				// at the finish
// 				results = append(results, len(current.visited))
// 				continue
// 			}

// 			legal := [][2]int{}
// 			for _, move := range moves {
// 				// try to move and check
// 				x, y := sx+move[0], sy+move[1]

// 				// was here before?
// 				if slices.Index(current.visited, [2]int{x, y}) != -1 {
// 					continue
// 				}

// 				if prevBest, ok := memo[[2]int{sx, sy}]; ok && prevBest > len(current.visited) {
// 					continue
// 				}

// 				// not in bounds? or a wall ?
// 				if x < 0 || x >= len(maze[0]) || y < 0 || y >= len(maze) || maze[y][x] == '#' {
// 					continue
// 				}
// 				// otherwise - legal position
// 				legal = append(legal, [2]int{x, y})
// 			}

// 			if len(legal) > 0 {
// 				// was here with better result?
// 				if prevBest, ok := memo[[2]int{sx, sy}]; ok && prevBest > len(current.visited) {
// 					continue
// 				}
// 				memo[[2]int{sx, sy}] = len(current.visited)
// 			}

// 			// more than one? fork the walks
// 			if len(legal) > 0 {

// 				for _, lm := range legal {
// 					forkedVisited := [][2]int{}
// 					// deep copy
// 					// for k, v := range current.visited {
// 					// 	forkedVisited[k] = v
// 					// }

// 					forkedVisited = slices.Clone(current.visited)
// 					// forkedVisited = maps.Clone(current.visited)
// 					nq = append(nq, Q{forkedVisited, [2]int{lm[0], lm[1]}})
// 				}
// 				// break
// 			}
// 			// log.Println(nq)
// 		}
// 		q = nq
// 		if i%100 == 0 {
// 			log.Println(len(q))
// 		}
// 		i++
// 	}
// 	// t.SkipNow()
// 	// p2 := walkMaze(maze, [2]int{1, 1}, map[[2]int]bool{})
// 	log.Println(slices.Max(results)) //
// }
//
//

// Fail:
//
//
// func TestP2(t *testing.T) {
// 	maze := readLines("../aoc-inputs/2023/23/input.txt")
// 	require.True(t, len(maze) > 0)

// 	start := [2]int{1, 1}

// 	q := []Q{{map[[2]int]bool{}, start}}
// 	results := []int{}
// 	moves := map[rune][2]int{'^': {0, -1}, '<': {-1, 0}, 'v': {0, 1}, '>': {1, 0}}
// 	maze[0][1] = '#'
// 	// showMaze(maze, false)
// 	// visited := map[[2]int]bool{}
// 	for len(q) > 0 {
// 		nq := []Q{}
// 		for _, current := range q {

// 			sx, sy := current.pos[0], current.pos[1]
// 			current.visited[[2]int{sx, sy}] = true

// 			if sx == len(maze[0])-2 && sy == len(maze)-1 {
// 				// at the finish
// 				results = append(results, len(current.visited))
// 				continue
// 			}

// 			legal := [][2]int{}
// 			for _, move := range moves {
// 				// try to move and check
// 				x, y := sx+move[0], sy+move[1]
// 				// going back?
// 				if _, ok := current.visited[[2]int{x, y}]; ok {
// 					continue
// 				}
// 				// not in bounds?
// 				if x < 0 || x >= len(maze[0]) || y < 0 || y >= len(maze) ||
// 					// or a wall ?
// 					maze[y][x] == '#' {
// 					continue
// 				}
// 				// otherwise - legal position
// 				legal = append(legal, [2]int{x, y})
// 			}

// 			// forking here?
// 			if len(legal) > 1 {
// 				// was here with better result?
// 				if prevBest, ok := memo[[2]int{sx, sy}]; ok && prevBest > len(current.visited) {
// 					continue
// 				}
// 				memo[[2]int{sx, sy}] = len(current.visited)
// 			}

// 			// more than one? fork the walks
// 			if len(legal) > 0 {

// 				for _, lm := range legal {
// 					forkedVisited := map[[2]int]bool{}
// 					// deep copy maps
// 					for k, v := range current.visited {
// 						forkedVisited[[2]int{k[0], k[1]}] = v
// 					}
// 					nq = append(nq, Q{forkedVisited, [2]int{lm[0], lm[1]}})
// 				}
// 			}
// 			// log.Println(nq)
// 			q = nq
// 		}

// 	}
// 	// t.SkipNow()
// 	// p2 := walkMaze(maze, [2]int{1, 1}, map[[2]int]bool{})
// 	log.Println(slices.Max(results)) //
// }
