package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"strconv"
	"strings"
)

var moves map[rune]image.Point = map[rune]image.Point{
	'R': {1, 0},
	'L': {-1, 0},
	'U': {0, -1},
	'D': {0, 1},
}

var xyDir = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

type Field map[image.Point]int

var f Field
var w, h int

func read(in string) (m Field, w int, h int) {

	points := []image.Point{}

	m = make(Field)
	minx, miny := math.MaxInt, math.MaxInt
	pos := image.Point{0, 0}
	for _, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {

		parts := strings.Split(strings.TrimSpace(l), " ")
		if len(parts) == 0 {
			continue
		}

		move := l[0]
		amount, _ := strconv.Atoi(parts[1])
		for i := 0; i < amount; i++ {
			pos = pos.Add(moves[rune(move)])
			points = append(points, pos)

			w = max(w, pos.X)
			minx = min(minx, pos.X)
			h = max(h, pos.Y)
			miny = min(miny, pos.Y)
		}
	}
	w = w + (-minx) + 2
	h = h + (-miny) + 2
	for _, p := range points {
		m[image.Pt(p.X+(-minx)+1, p.Y+(-miny)+1)] = 1
	}

	return
}

func fill(v *Field, start image.Point) {
	_, ok := (*v)[start]

	if !ok {
		(*v)[start] = 2
		for i := range xyDir {
			candidate := start.Add(xyDir[i])

			// (?) in bounds
			if candidate.X > w || candidate.X < 0 || candidate.Y > h || candidate.Y < 0 {
				continue
			}

			fill(v, candidate)
		}
	}
}

func show(m map[image.Point]int, w, h int) {
	asciiBlocks := []string{"░░", "██", "▒▒"}
	fmt.Println(w+1, "x", h+1, ":")
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			fmt.Print(asciiBlocks[m[image.Pt(x, y)]])
		}
		fmt.Println()
	}
	fmt.Println()
}

func p1(file string) int {
	in, _ := os.ReadFile(file)
	f, w, h = read(string(in))
	fill(&f, image.Point{0, 0})
	// show(f, w, h)
	lava := 0
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			if v, ok := f[image.Pt(x, y)]; !ok || v == 1 {
				lava++
			}
		}
	}

	return lava
}

func main() {
	args := os.Args[1:]
	defaultInput := "../aoc-inputs/2023/18/input.txt"
	if len(args) > 0 {
		defaultInput = args[0]
	}

	fmt.Println("Day 18: Lavaduct Lagoon")
	fmt.Println("\tPart One:", p1(defaultInput)) // 62573
	// fmt.Println("\tPart Two:", p2(defaultInput)) //
}
