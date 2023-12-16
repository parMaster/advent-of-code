package main

import (
	"fmt"
	"image"
	"os"
	"slices"
	"strings"
)

// directions	↑ → ↓ ←
// directions	0 1 2 3
var xyDir = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

// new direction == reflects[mirror][current direction]
var reflects = map[rune]map[image.Point]image.Point{
	'/': {
		xyDir[0]: xyDir[1],
		xyDir[1]: xyDir[0],
		xyDir[2]: xyDir[3],
		xyDir[3]: xyDir[2],
	},
	'\\': {
		xyDir[0]: xyDir[3],
		xyDir[1]: xyDir[2],
		xyDir[3]: xyDir[0],
		xyDir[2]: xyDir[1],
	},
}

// []directions == forks[-|]
var forks = map[rune][]image.Point{
	'-': {xyDir[1], xyDir[3]},
	'|': {xyDir[0], xyDir[2]},
}

type Beam struct {
	pos image.Point
	dir image.Point
}

// visited[position]dirrections to check cycles and count answer
var visited map[image.Point][]image.Point = map[image.Point][]image.Point{}

type Field map[image.Point]rune

func read(in string) (m Field, w int, h int) {
	m = make(map[image.Point]rune)
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		w = len(l) - 1
		h = y
		for x, r := range strings.TrimSpace(l) {
			m[image.Point{x, y}] = r
		}
	}
	return
}

// Part one without recursion
func p1(f string) int {
	in, _ := os.ReadFile(f)
	field, w, h := read(string(in))

	var beams = []Beam{{pos: image.Point{0, 0}, dir: xyDir[1]}}
	progressed := true
	for progressed {
		progressed = false
		for beamid, beam := range beams {
			// (?) check out of bounds
			if beam.pos.X > w || beam.pos.X < 0 || beam.pos.Y > h || beam.pos.Y < 0 {
				continue
			}

			// (?) check if there was such beam before?
			if v, ok := visited[beam.pos]; ok && slices.Index(v, beam.dir) != -1 {
				continue
			}

			// no? then visited
			visited[beam.pos] = append(visited[beam.pos], beam.dir)
			progressed = true

			// 2. reflects?
			if v, ok := reflects[field[beam.pos]][beam.dir]; ok {
				beam.dir = v
			}

			// 3. forks?
			if fork, ok := forks[field[beam.pos]]; ok {
				beam.dir = fork[0]
				beams = append(beams, Beam{pos: beam.pos.Add(fork[1]), dir: fork[1]})
			}
			// saving changes
			beam.pos = beam.pos.Add(beam.dir)
			beams[beamid] = beam
		}
	}

	return len(visited)
}

func main() {
	fmt.Println("Day 16: The floor will be lava")
	fmt.Println("\tPart One:", p1("../aoc-inputs/2023/16/input.txt")) // 6795
	// fmt.Println("\tPart Two:", p2("../aoc-inputs/2023/16/input.txt")) //
}
