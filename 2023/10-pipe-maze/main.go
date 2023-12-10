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

func read(in string) (map[image.Point]rune, map[image.Point]int, image.Point) {
	m, v := make(map[image.Point]rune), make(map[image.Point]int)
	start := image.Point{0, 0}
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		for x, r := range strings.TrimSpace(l) {
			m[image.Point{x, y}] = r
			if r == 'S' {
				start = image.Point{x, y}
			}
		}
	}
	v[start] = 1
	return m, v, start
}

func connects(m map[image.Point]rune, from, to image.Point) bool {
	for _, s := range ps[m[from]] {
		if from.Add(s) == to {
			return true
		}
	}
	return false
}

func solve(f string) int {
	m, v, pos := read(f)
	for {
		steps := ps[m[pos]]
		for _, step := range steps {

			newPos := pos.Add(step)

			if connects(m, pos, newPos) && connects(m, newPos, pos) {
				if v[newPos] == 1 {
					show(m, v)
					return slices.Max(maps.Values(v)) / 2
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

func main() {
	in, _ := os.ReadFile("../aoc-inputs/2023/10/input.txt")
	fmt.Println("Day 10: Pipe Maze")
	fmt.Println("\tPart One:", solve(string(in))) // 7107
}

func show(m map[image.Point]rune, v map[image.Point]int) {

	mx, my := 0, 0
	for p := range m {
		mx = max(mx, p.X)
		my = max(my, p.Y)
	}
	fmt.Println(mx, "x", my, ":")
	for x := 0; x <= mx; x++ {
		for y := 0; y <= my; y++ {
			if v[image.Pt(y, x)] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(m[image.Pt(y, x)]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
