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

func main() {
	fmt.Println("Day 9: Rope Bridge\n\tPart One:", PartOne("input1.txt"))
}
