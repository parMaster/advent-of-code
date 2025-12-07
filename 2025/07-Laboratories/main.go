package main

import (
	"fmt"
	"image"
	"maps"
	"os"
	"slices"
	"strings"
	"time"
)

// https://adventofcode.com/2025/day/7

type Grid map[image.Point]rune

var moves = []image.Point{
	{0, 1},  // down
	{-1, 0}, // left
	{1, 0},  // right
}

func read(file string) (Grid, image.Point, int, int) {
	in, _ := os.ReadFile(file)
	m := Grid{}
	lines := strings.Split(string(in), "\n")
	h, w := len(lines), len(lines[0])
	start := image.Pt(0, 0)
	for y, line := range lines {
		for x, r := range line {
			switch r {
			case 'S':
				start = image.Pt(x, y)
				m[image.Pt(x, y)] = '|'
			case '^':
				m[image.Pt(x, y)] = r
			}
		}
	}
	return m, start, w, h
}

func p2_recursive(file string) (p2Rec int) {
	m, start, _, h := read(file)
	return rec(m, start, h)
}

func solve(file string) (p1, p2 int) {
	m, start, w, h := read(file)
	if slices.Contains(os.Args[1:], "--visual") {
		m.Show(image.Pt(0, 0), image.Rect(0, 0, w, h), map[image.Point]int{})
	}

	// count number of beams in each point, split and combine numbers of beams
	quantumSplit := func(g Grid, start image.Point, h int) (splits, totalBeams int) {
		q := map[image.Point]int{start: 1}
		for range h {
			// m.Show(image.Pt(0, 0), image.Rect(0, 0, w, h), q)
			nq := map[image.Point]int{}
			for curr, beams := range q {
				down := curr.Add(moves[0])
				if v, ok := g[down]; ok && v == '^' {
					splits++
					// splitter
					for _, move := range moves[1:] {
						if sideBeams, ok := nq[down.Add(move)]; ok {
							nq[down.Add(move)] = sideBeams + beams
						} else {
							nq[down.Add(move)] = beams
						}

					}
				} else {
					// no splitter, could be a bunch of beams
					if v, ok := nq[down]; ok {
						nq[down] = beams + v
					} else {
						nq[down] = beams
					}
				}
			}
			q = maps.Clone(nq)
		}
		for _, beams := range q {
			totalBeams += beams
		}
		return splits, totalBeams
	}

	return quantumSplit(m, start, h)
}

func main() {
	start := time.Now()
	fmt.Println("Day 07: Laboratories")
	p1, p2 := solve("../aoc-inputs/2025/07/input.txt")
	// p1, p2 := solve("input-pub.txt")
	fmt.Println("\tPart One:", p1) // 1628
	fmt.Println("\tPart Two:", p2) // 27055852018812
	fmt.Printf("Done in %.6f seconds \n", time.Since(start).Seconds())
	start = time.Now()
	p2Rec := p2_recursive("../aoc-inputs/2025/07/input.txt")
	fmt.Println("\tPart Two (recursive):", p2Rec) // 27055852018812
	fmt.Printf("Done in %.6f seconds \n", time.Since(start).Seconds())
}

func (g Grid) Show(r image.Point, bounds image.Rectangle, nq map[image.Point]int) {
	fmt.Println(bounds.Max.X, "x", bounds.Max.Y, ":")
	for y := 0; y <= bounds.Max.Y; y++ {
		fmt.Printf(" %2.0d ", y)
		for x := 0; x <= bounds.Max.X; x++ {
			if v, ok := g[image.Pt(x, y)]; ok {
				fmt.Printf(" %s ", string(v))
			} else {
				if v, ok := nq[image.Pt(x, y)]; ok {
					fmt.Printf(" %d ", v)
				} else {
					fmt.Print(" . ")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

var memo map[image.Point]int = map[image.Point]int{}

// recursive part 2 with memoization
func rec(m Grid, beam image.Point, h int) int {
	if beam.Y == h {
		return 1
	}
	if v, ok := memo[beam]; ok {
		return v
	}
	down := beam.Add(moves[0])
	if v, ok := m[down]; ok && v == '^' {
		dr, dl := down.Add(moves[1]), down.Add(moves[2])
		res := rec(m, dl, h) + rec(m, dr, h)
		memo[beam] = res
		return res
	}

	res := rec(m, down, h)
	memo[beam] = res
	return res
}
