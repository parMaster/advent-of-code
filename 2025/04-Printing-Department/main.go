package main

import (
	"advent-of-code/2025/utils"
	"fmt"
	"image"
	"os"
	"strings"
	"time"
)

// directions   ↖ ↑ ↗ → ↘ ↓ ↙ ←
// directions	0 1 2 3 4 5 6 7
var Dir = []image.Point{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}

func solve(file string) (p1, p2 int) {
	in, _ := os.ReadFile(file)
	g, bounds := ReadGrid(string(in))

	for run := range 135 * 135 {
		q := []image.Point{}
		for p, cell := range g {
			if cell != '@' {
				continue
			}

			adjRolls := 0
			for _, dir := range Dir {
				if adj, ok := g[image.Point.Add(p, dir)]; ok && adj == '@' {
					adjRolls++
				}
			}
			if adjRolls < 4 {
				if run == 0 {
					p1++
				}
				q = append(q, p)
			}
		}
		if len(q) > 0 {
			for _, p := range q {
				g[p] = 'x'
				p2++
			}
		} else {
			break
		}
	}

	g.Render(Dir[0], bounds, map[rune]string{'.': "empty", 'x': "half", '@': "full"})
	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 04: Printing Department")
	p1, p2 := solve("../aoc-inputs/2025/04/input.txt")
	// p1, p2 := solve("input-pub.txt")
	fmt.Println("\tPart One:", p1) // 1344
	fmt.Println("\tPart Two:", p2) // 8112
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}

func ReadGrid(in string) (g utils.Grid, bounds image.Rectangle) {
	g = make(map[image.Point]rune)
	var w, h int
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		w = len(l)
		h = y + 1
		for x, r := range strings.TrimSpace(l) {
			g[image.Point{x, y}] = r
		}
	}
	bounds = image.Rect(0, 0, w, h)
	return
}
