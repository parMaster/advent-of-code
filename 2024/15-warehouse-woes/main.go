package main

import (
	"fmt"
	"image"
	"os"
	"strings"
	"time"
)

type Grid map[image.Point]rune

func ReadGrid(in string) (g Grid, bounds image.Rectangle, robot image.Point) {
	g = make(map[image.Point]rune)
	var w, h int
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		w = len(l)
		h = y
		for x, r := range strings.TrimSpace(l) {
			if r == '@' {
				robot = image.Pt(x, y)
				g[image.Point{x, y}] = '.'
				continue
			}
			g[image.Point{x, y}] = r
		}
	}
	bounds = image.Rect(0, 0, w, h)
	return
}

var dir = map[rune]image.Point{'^': {0, -1}, '>': {1, 0}, 'v': {0, 1}, '<': {-1, 0}}

func (g Grid) moveBoxes(p image.Point, d rune) bool {
	np := p.Add(dir[d])
	if g[np] == 'O' {
		if g.moveBoxes(np, d) {
			g[p], g[np] = g[np], g[p]
			return true
		}
	}
	if g[np] == '.' {
		g[p], g[np] = g[np], g[p]
		return true
	}
	return false
}

func (grid Grid) p1(robot image.Point, moves string) (res int) {

	for _, m := range moves {
		np := robot.Add(dir[m])
		if grid[np] == '.' {
			robot = np
		}
		if grid[np] == 'O' {
			if grid.moveBoxes(np, m) {
				robot = np
			}
		}
	}

	for p := range grid {
		if grid[p] == 'O' {
			res += 100*p.Y + p.X
		}
	}

	return
}

func solve(file string) (p1, p2 int) {
	in, _ := os.ReadFile(file)
	parts := strings.Split(string(in), "\n\n")
	moves := strings.ReplaceAll(parts[1], "\n", "")
	grid, bounds, robot := ReadGrid(parts[0])

	grid.Show(robot, bounds)
	p1 = grid.p1(robot, moves)

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 15: Warehouse Woes")
	// p1, p2 := solve("input00.txt")
	p1, p2 := solve("../aoc-inputs/2024/15/input.txt")
	fmt.Println("\tPart One:", p1) // 1448589
	fmt.Println("\tPart Two:", p2) //
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}

func (g Grid) Show(r image.Point, bounds image.Rectangle) {
	// fmt.Println(bounds.Max.X, "x", bounds.Max.Y, ":")
	for y := 0; y <= bounds.Max.Y; y++ {
		for x := 0; x <= bounds.Max.X; x++ {
			if r == image.Pt(x, y) {
				fmt.Print("@")
				continue
			}
			fmt.Print(string(g[image.Pt(x, y)]))
		}
		fmt.Println()
	}
	fmt.Println()
}
