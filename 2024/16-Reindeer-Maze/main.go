package main

import (
	"fmt"
	"image"
	"maps"
	"os"
	"slices"
	"strings"
	"time"
)

type Grid map[image.Point]rune

func ReadGrid(in string) (g Grid, bounds image.Rectangle, start, end image.Point) {
	g = make(map[image.Point]rune)
	var w, h int
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		w = len(l)
		h = y + 1
		for x, r := range strings.TrimSpace(l) {
			if r == 'S' {
				start = image.Pt(x, y)
				g[start] = '.'
				continue
			}
			if r == 'E' {
				end = image.Pt(x, y)
				g[end] = '.'
				continue
			}
			g[image.Point{x, y}] = r
		}
	}
	bounds = image.Rect(0, 0, w, h)
	return
}

var dir = []image.Point{
	{1, 0},
	{0, -1},
	{-1, 0},
	{0, 1},
}

type Q struct {
	pos   image.Point
	dir   int
	score int
	path  map[image.Point]int
}

func (g Grid) bestScore(start, end image.Point, bounds image.Rectangle) (int, int) {
	scores := map[image.Point]int{start: 0}
	bestScore := 1<<32 - 1
	bestPath := map[int][]image.Point{}

	q := []Q{{start, 0, 0, map[image.Point]int{start: 0}}}
	for len(q) > 0 {
		// fmt.Println(len(q))
		curr := q[0]
		q = q[1:]

		for i := range 4 {
			newDir := (curr.dir + i) % 4
			newPos := curr.pos.Add(dir[newDir])
			rot := curr.dir - newDir
			if rot < 0 {
				rot = -rot
			}

			rot = min(rot, 1)
			newScore := (curr.score + rot*1000) + 1

			if g[newPos] == '#' {
				continue
			}

			if dir[newDir].Add(dir[curr.dir]) == image.Pt(0, 0) {
				continue
			}

			if _, ok := curr.path[newPos]; ok {
				continue
			}

			if posScore, ok := scores[newPos]; ok && posScore < newScore {
				continue
			}

			if newPos == end && newScore <= bestScore {

				curr.path[newPos] = 1
				bestScore = newScore
				bestPath[newScore] = []image.Point{}

				continue
			}

			newPath := map[image.Point]int{}
			maps.Copy(newPath, curr.path)
			newPath[newPos] = newDir

			q = append(q, Q{newPos, newDir, newScore, newPath})
			scores[newPos] = newScore
		}
	}

	fmt.Println(scores)
	g.Show(bounds, bestPath[bestScore])

	return bestScore, 0
}

func solve(file string) (p1, p2 int) {
	input, _ := os.ReadFile(file)
	g, bounds, start, end := ReadGrid(string(input))
	fmt.Println(bounds)
	// g, _, start, end := ReadGrid(string(input))

	p1, _ = g.bestScore(start, end, bounds)
	// _, p2 = g.bestScore(start, end, bounds, p1)

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 16: Reindeer Maze")
	// p1, p2 := solve("input1.txt")
	p1, p2 := solve("../aoc-inputs/2024/16/input.txt")
	fmt.Println("\tPart One:", p1) //
	fmt.Println("\tPart Two:", p2) // 440 low
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}

func (g Grid) Show(bounds image.Rectangle, qm []image.Point) {
	fmt.Println(bounds.Max.X, "x", bounds.Max.Y, ":")

	dirs := []rune{
		'>', '^', '<', 'v', 'O',
	}

	for y := 0; y <= bounds.Max.Y; y++ {
		for x := 0; x <= bounds.Max.X; x++ {
			if slices.Contains(qm, image.Pt(x, y)) {
				fmt.Print(string(dirs[4]))
				continue
			}
			fmt.Print(string(g[image.Pt(x, y)]))
		}
		fmt.Println()
	}
	fmt.Println()
}

// func (g Grid) allBestScore(start, end image.Point) int {
// 	scores := map[image.Point]int{start: 0}
// 	bestScore := 1<<32 - 1
// 	// bestPaths := map[int][]image.Point{}

// 	q := []Q{{start, 0, 0, map[image.Point]int{{0, 0}: 0}}}
// 	for len(q) > 0 {
// 		// fmt.Println("len(q)", len(q))
// 		curr := q[0]
// 		q = q[1:]

// 		for i := range 4 {
// 			newDir := (curr.dir + i) % 4
// 			newPos := curr.pos.Add(dir[newDir])
// 			rot := curr.dir - newDir
// 			if rot < 0 {
// 				rot = -rot
// 			}

// 			rot = min(rot, 1)
// 			newScore := (curr.score + rot*1000) + 1

// 			// if newScore > bestScore {
// 			// 	continue
// 			// }
// 			if g[newPos] == '#' {
// 				continue
// 			}

// 			// if dir[newDir].Add(dir[curr.dir]) == image.Pt(0, 0) {
// 			// 	continue
// 			// }

// 			if _, ok := curr.path[newPos]; ok {
// 				continue
// 			}

// 			if posScore, ok := scores[newPos]; ok && posScore < newScore {
// 				continue
// 			}

// 			if newPos == end && newScore <= bestScore {
// 				fmt.Println("found", newScore)
// 				// fmt.Println("found", curr.path)
// 				curr.path[newPos] = 1

// 				// if _, ok := bestPaths[newScore]; !ok {
// 				// 	bestPaths[newScore] = []image.Point{}
// 				// }
// 				// for k := range curr.path {
// 				// 	if !slices.Contains(bestPaths[newScore], k) {
// 				// 		bestPaths[newScore] = append(bestPaths[newScore], k)
// 				// 	}
// 				// }

// 				bestScore = newScore
// 				continue
// 			}

// 			newPath := map[image.Point]int{}
// 			// maps.Copy(newPath, curr.path)
// 			// newPath[newPos] = newDir
// 			// fmt.Println(newPos)
// 			q = append(q, Q{newPos, newDir, newScore, newPath})
// 			scores[newPos] = newScore
// 		}
// 	}
// 	// g.Show(bounds, bestPaths[bestScore])

// 	return bestScore
// }

// func (g Grid) bestScore(start, end image.Point, bounds image.Rectangle, best int) (int, int) {
// 	scores := map[image.Point]int{start: 0}
// 	bestScore := 1<<32 - 1
// 	bestPaths := map[int][]image.Point{}

// 	q := []Q{{start, 0, 0, map[image.Point]int{start: 0}}}
// 	for len(q) > 0 {
// 		// fmt.Println(len(q))
// 		curr := q[0]
// 		q = q[1:]

// 		for i := range 4 {
// 			newDir := (curr.dir + i) % 4
// 			newPos := curr.pos.Add(dir[newDir])
// 			rot := curr.dir - newDir
// 			if rot < 0 {
// 				rot = -rot
// 			}

// 			rot = min(rot, 1)
// 			newScore := (curr.score + rot*1000) + 1

// 			if g[newPos] == '#' {
// 				continue
// 			}

// 			if dir[newDir].Add(dir[curr.dir]) == image.Pt(0, 0) {
// 				continue
// 			}

// 			if _, ok := curr.path[newPos]; ok {
// 				continue
// 			}

// 			if posScore, ok := scores[newPos]; ok && posScore < newScore {
// 				continue
// 			}

// 			if best == 0 {
// 			} else {
// 				if best < newScore {
// 					continue
// 				}
// 			}

// 			if newPos == end && newScore <= bestScore {

// 				curr.path[newPos] = 1
// 				if newScore == best {
// 					if _, ok := bestPaths[newScore]; !ok {
// 						bestPaths[newScore] = []image.Point{}
// 					}
// 					for k := range curr.path {
// 						if !slices.Contains(bestPaths[newScore], k) {
// 							bestPaths[newScore] = append(bestPaths[newScore], k)
// 						}
// 					}
// 				}

// 				bestScore = newScore
// 				continue
// 			}

// 			newPath := map[image.Point]int{}
// 			// if best > 0 {
// 			maps.Copy(newPath, curr.path)
// 			newPath[newPos] = newDir
// 			// }
// 			// fmt.Println(newPos)

// 			q = append(q, Q{newPos, newDir, newScore, newPath})
// 			scores[newPos] = newScore
// 		}
// 	}

// 	fmt.Println(scores)
// 	g.Show(bounds, bestPaths[bestScore])

// 	return bestScore, len(bestPaths[bestScore])
// }
