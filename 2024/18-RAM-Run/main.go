package main

import (
	"encoding/json"
	"fmt"
	"image"
	"maps"
	"os"
	"strings"
	"time"
)

// directions	→ ↓ ← ↑
// directions	0 1 2 3
var XYDir = []image.Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

type Grid map[image.Point]rune

func solve(file string, size int) (p1 int, p2 string) {
	in, _ := os.ReadFile(file)
	m := [][]int{}
	json.Unmarshal([]byte("[["+strings.ReplaceAll(string(in), "\n", "],[")+"]]"), &m)
	grid, w, h := Grid{}, 0, 0
	for _, v := range m[:size] {
		grid[image.Pt(v[0], v[1])] = '#'
		w, h = max(w, v[0]), max(h, v[1])
	}
	for x := range w + 1 {
		for y := range h + 1 {
			if _, ok := grid[image.Pt(x, y)]; !ok {
				grid[image.Pt(x, y)] = '.'
			}
		}
	}
	bounds := image.Rect(0, 0, w, h)

	grid.Show(image.Pt(0, 0), bounds)

	minPath := func(grid Grid) int {
		visited := map[image.Point]int{image.Pt(0, 0): 0}
		nextQueue := map[image.Point]int{image.Pt(0, 0): 0}

		run := 0
		for {
			run++
			// progress from every queue item
			queue := maps.Clone(nextQueue)
			nextQueue = map[image.Point]int{}
			for from, score := range queue {

				// try to step in every direction
				for _, dir := range XYDir {
					next := from.Add(dir)
					nextScore := score + 1

					// wall
					if cell, ok := grid[next]; ok {
						if cell == '#' {
							continue
						}
					} else {
						// bounds
						continue
					}

					// visited with better path before?
					if bestScore, ok := visited[next]; ok && bestScore <= nextScore {
						continue
					}

					// finish reached
					if next == image.Pt(w, h) {
						return nextScore
					}

					// new path with best score
					visited[next] = nextScore
					nextQueue[next] = nextScore
				}
			}
			// path doesn't exist
			if len(nextQueue) == 0 {
				return -1
			}
		}
		// return -1
	}

	p1 = minPath(grid)

	// p2
	for _, v := range m[size:] {
		grid[image.Pt(v[0], v[1])] = '#'
		minp := minPath(grid)
		if minp == -1 {
			p2 = fmt.Sprintf("%d,%d", v[0], v[1])
			grid.Show(image.Pt(0, 0), bounds)
			return
		}
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 18: RAM Run")
	// p1, p2 := solve("input-pub.txt", 12)
	p1, p2 := solve("../aoc-inputs/2024/18/input.txt", 1024)
	fmt.Println("\tPart One:", p1) // 304
	fmt.Println("\tPart Two:", p2) // 50,28
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}

func (g Grid) Show(r image.Point, bounds image.Rectangle) {
	fmt.Println(bounds.Max.X, "x", bounds.Max.Y, ":")
	for y := 0; y <= bounds.Max.Y; y++ {
		for x := 0; x <= bounds.Max.X; x++ {
			if v, ok := g[image.Pt(x, y)]; ok {
				fmt.Printf("%c", v)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
