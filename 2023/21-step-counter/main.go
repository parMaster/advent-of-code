package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

type Stones map[image.Point]bool

func read(in string) (Stones, image.Point, int, int) {
	m := make(map[image.Point]bool)
	mx, my := 0, 0
	start := image.Point{0, 0}
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		for x, r := range strings.TrimSpace(l) {
			mx = max(mx, x)
			my = max(my, y)
			if r == '#' {
				m[image.Point{x, y}] = true
			}
			if r == 'S' {
				start = image.Point{x, y}
			}
		}
	}
	return m, start, mx, my
}

type Queue map[image.Point]bool

var moves []image.Point = []image.Point{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

func main() {
	in, _ := os.ReadFile("../aoc-inputs/2023/21/input.txt")

	stones, start, w, h := read(string(in))

	show(stones, w, h)

	N := 0
	var q Queue = Queue{start: true}
	for N < 64 && len(q) > 0 {

		var nq Queue = Queue{}
		for i := range q {
			// check around q[i]:
			for _, m := range moves {

				// candidate for next step
				c := i.Add(m)

				// in bounds?
				if c.X > w || c.X < 0 || c.Y > w || c.Y < 0 {
					continue
				}

				// stone ?
				if v, ok := stones[c]; ok && v {
					continue
				}

				// otherwise legal next step
				nq[i.Add(m)] = true
			}
		}
		q = maps.Clone(nq)
		N++
		// log.Println(N, len(maps.Keys(q)))
	}
	log.Println("Part One:", len(q)) // 3646
}

func show(m map[image.Point]bool, w, h int) {
	if slices.Index(os.Args[1:], "--visual") == -1 {
		return
	}
	asciiBlocks := []string{"░░", "██", "▒▒"}
	fmt.Println(w+1, "x", h+1, ":")
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			if m[image.Pt(x, y)] {
				fmt.Print(asciiBlocks[1])
			} else {
				fmt.Print(asciiBlocks[0])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
