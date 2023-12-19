package main

import (
	"fmt"
	"image"
	"os"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

// directions	↑ → ↓ ←
// directions	0 1 2 3
var xyDir = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

// possible steps - each element has two, start has four
var ps = map[rune][]image.Point{
	'|': {xyDir[0], xyDir[2]},
	'-': {xyDir[1], xyDir[3]},
	'L': {xyDir[0], xyDir[1]},
	'J': {xyDir[0], xyDir[3]},
	'7': {xyDir[2], xyDir[3]},
	'F': {xyDir[1], xyDir[2]},
	'.': {},
	'S': xyDir,
}

type maze map[image.Point]rune

type loop map[image.Point]int

func read(in string) (maze, image.Point) {
	m := make(map[image.Point]rune)
	start := image.Point{0, 0}
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		for x, r := range strings.TrimSpace(l) {
			m[image.Point{x, y}] = r
			if r == 'S' {
				start = image.Point{x, y}
			}
		}
	}
	return m, start
}

func connects(m map[image.Point]rune, from, to image.Point) bool {
	for _, s := range ps[m[from]] {
		if from.Add(s) == to {
			return true
		}
	}
	return false
}

func p1(f string) int {
	m, pos := read(f)
	p1, _ := giantLoop(m, pos)
	return p1
}

func giantLoop(m maze, pos image.Point) (int, loop) {
	var v loop = loop{pos: 1}
	for {
		steps := ps[m[pos]]
		for _, step := range steps {

			newPos := pos.Add(step)

			if connects(m, pos, newPos) && connects(m, newPos, pos) {
				if v[newPos] == 1 {
					return slices.Max(maps.Values(v)) / 2, v
				}
				if v[newPos] == 0 {
					v[newPos] = v[pos] + 1
					pos = newPos
					break
				}
			}
		}
	}
}

func show(m map[image.Point]rune, v map[image.Point]int) {
	mx, my := maxPoint(m)
	fmt.Println(mx, "x", my, ":")
	for y := 0; y <= my; y++ {
		for x := 0; x <= mx; x++ {
			if v[image.Pt(x, y)] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(m[image.Pt(x, y)]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func maxPoint[V any](m map[image.Point]V) (mx, my int) {
	for p := range m {
		mx = max(mx, p.X)
		my = max(my, p.Y)
	}
	return mx, my
}

func main() {
	in, _ := os.ReadFile("../aoc-inputs/2023/10/input.txt")
	fmt.Println("Day 10: Pipe Maze")
	fmt.Println("\tPart One:", p1(string(in))) // 7107
	fmt.Println("\tPart Two:", p2(string(in))) // 281
}

func p2(f string) int {
	maze, start := read(f)
	_, loop := giantLoop(maze, start)

	repairedQuadLoop := repair(quad(maze), quad(loop))
	if slices.Index(os.Args[1:], "--visual") != -1 {
		showLoop(repairedQuadLoop)
	}

	mx, my := maxPoint(repairedQuadLoop)
	// fill from each corner, quick and dirty
	fill(&repairedQuadLoop, image.Pt(0, 0))
	fill(&repairedQuadLoop, image.Pt(mx, 0))
	fill(&repairedQuadLoop, image.Pt(0, my))
	fill(&repairedQuadLoop, image.Pt(mx, my))

	loop = shrink(repairedQuadLoop)

	if slices.Index(os.Args[1:], "--visual") != -1 {
		showLoop(loop)
	}

	return countZeroes(loop)
}

// quadruples image in size, surrounds every value with
func quad[V int | rune](loop map[image.Point]V) map[image.Point]V {
	quadLoop := make(map[image.Point]V)
	mx, my := maxPoint(loop)
	for x := 0; x <= mx; x++ {
		for y := 0; y <= my; y++ {
			quadLoop[image.Pt(x*2, y*2)] = loop[image.Pt(x, y)]
			quadLoop[image.Pt(x*2+1, y*2)] = 0
			quadLoop[image.Pt(x*2, y*2+1)] = 0
			quadLoop[image.Pt(x*2+1, y*2+1)] = 0
		}
	}
	return quadLoop
}

// connects indirectly, over one point
func overConnects(m map[image.Point]rune, from, to image.Point) bool {
	for _, s := range ps[m[from]] {
		if from.Add(s).Add(s) == to {
			return true
		}
	}
	return false
}

// repairs quadrupled loop
func repair(maze maze, loop loop) loop {
	mx, my := maxPoint(loop)
	for x := 0; x <= mx; x++ {
		for y := 0; y <= my; y++ {

			if overConnects(maze, image.Pt(x-1, y), image.Pt(x+1, y)) &&
				overConnects(maze, image.Pt(x+1, y), image.Pt(x-1, y)) {
				loop[image.Pt(x, y)] = 1
			}

			if overConnects(maze, image.Pt(x, y-1), image.Pt(x, y+1)) &&
				overConnects(maze, image.Pt(x, y+1), image.Pt(x, y-1)) {
				loop[image.Pt(x, y)] = 1
			}

		}
	}
	return loop
}

func fill(v *loop, start image.Point) {
	val, ok := (*v)[start]
	if !ok || val > 0 {
		return
	}
	(*v)[start] = 1
	for i := range xyDir {
		candidate := start.Add(xyDir[i])
		fill(v, candidate)
	}
}

// shrink
func shrink[V int | rune](loop map[image.Point]V) map[image.Point]V {
	mx, my := maxPoint(loop)
	for x := 0; x <= (mx-1)/2; x++ {
		for y := 0; y <= (my-1)/2; y++ {
			delete(loop, image.Pt(x*2+1, y*2))
			delete(loop, image.Pt(x*2, y*2+1))
			delete(loop, image.Pt(x*2+1, y*2+1))
		}
	}
	return loop
}

// count zeroes
func countZeroes(loop map[image.Point]int) int {
	cnt := 0
	mx, my := maxPoint(loop)
	for y := 0; y <= my+1; y++ { // lines
		for x := 0; x <= mx+1; x++ { // cols
			v, ok := loop[image.Pt(x, y)]
			if !ok {
				continue
			}
			if v == 0 {
				cnt++
				continue
			}
		}
	}

	return cnt
}

func showLoop(loop map[image.Point]int) {
	mx, my := maxPoint(loop)
	fmt.Println("Loop: mx=", mx, " my=", my, ":")
	fmt.Println()
	for y := 0; y <= my+1; y++ { // lines
		for x := 0; x <= mx+1; x++ { // cols
			v, ok := loop[image.Pt(x, y)]
			if !ok {
				fmt.Print(" ")
				continue
			}
			if v == 0 {
				fmt.Print("░")
				continue
			}
			fmt.Print("█")
		}
		fmt.Println()
	}
	fmt.Println()
}
