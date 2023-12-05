package main

import (
	"fmt"
	"image"
	"os"
	"slices"
	"strconv"
	"strings"
)

// directions   ↖ ↑ ↗ → ↘ ↓ ↙ ←
// directions	0 1 2 3 4 5 6 7
var d = []image.Point{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}

var dir = map[string]image.Point{"U": d[1], "R": d[3], "D": d[5], "L": d[7]}

// returns true if t is adjacent (in the d square around) to h
func gap(h, t image.Point) bool {
	for _, p := range d {
		if t == h.Add(p) {
			return false
		}
	}
	return true && h != t
}

func pair(s string) (string, int) {
	ss := strings.Fields(strings.TrimSpace(s))
	n, _ := strconv.Atoi(ss[1])
	return string(ss[0][0]), n
}

func PartOne(file string) int {
	input, _ := os.ReadFile(file)
	h, t := image.Point{}, image.Point{}
	visited := []string{t.String()}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, l := range lines {

		move, size := pair(l)

		for i := 0; i < size; i++ {
			prev := h
			h = h.Add(dir[move])
			if gap(h, t) {
				t = prev
				visited = append(visited, t.String())
			}
		}

	}

	slices.Sort(visited)
	visited = slices.Compact(visited)
	return len(visited)
}

/** Part Two */

func abs(d image.Point) image.Point {
	if d.X < 0 {
		d.X = -d.X
	}

	if d.Y < 0 {
		d.Y = -d.Y
	}
	return d
}

func d1(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

func PartTwo(file string) int {
	visited := map[image.Point]bool{}
	input, _ := os.ReadFile(file)
	k := [10]image.Point{} // knots (a head + 9 knots)

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, l := range lines {
		move, size := pair(l)

		for i := 0; i < size; i++ {

			k[0] = k[0].Add(dir[move])

			for j := 1; j < len(k); j++ {

				d := k[j-1].Sub(k[j])
				if abs(d).X > 1 || abs(d).Y > 1 {
					k[j] = k[j].Add(image.Point{d1(d.X), d1(d.Y)})
				}
			}
			visited[k[len(k)-1]] = true

		}
	}

	return len(visited)
}

func main() {
	fmt.Println("Day 9: Rope Bridge\n\tPart One:", PartOne("../aoc-inputs/2022/09/input1.txt"))
	fmt.Println("\tPart Two:", PartTwo("../aoc-inputs/2022/09/input1.txt"))
}
