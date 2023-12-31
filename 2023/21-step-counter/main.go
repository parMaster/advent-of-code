package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

// [y][x]true is a stone
type Field [][]bool

func readLines(in string, tile int) (Field, int, int) {
	tile--
	f := Field{}
	sx, sy := 0, 0
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		f = append(f, []bool{})
		for x, r := range strings.TrimSpace(l) {
			if r == '#' {
				f[y] = append(f[y], false)
			} else {
				f[y] = append(f[y], true)
			}
			if r == 'S' {
				sx, sy = x, y
			}
		}

		if tile > 0 {
			fy := slices.Clone(f[y])
			for i := 0; i < tile; i++ {
				f[y] = append(f[y], fy...)
			}
		}
	}
	if tile > 0 {
		fc := slices.Clone(f)
		for i := 0; i < tile; i++ {
			f = append(f, fc...)
		}
		sx = len(f[0]) / 2
		sy = len(f) / 2
	}

	return f, sx, sy
}

func showLines(f Field, sx, sy int) {
	asciiBlocks := []string{"░░", "██", "▒▒"}
	for y := 0; y < len(f); y++ {
		for x := 0; x < len(f[y]); x++ {
			if x == sx && y == sy {
				fmt.Print(asciiBlocks[2])
				continue
			}
			if f[y][x] {
				fmt.Print(asciiBlocks[0])
			} else {
				fmt.Print(asciiBlocks[1])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Queue of coordinates to step next
type Q map[[2]int]bool

func reach(f Field, sx, sy, steps int) []int {
	res := []int{0}
	q := Q{{sx, sy}: true}
	N := 0
	for N < steps && len(q) > 0 {
		nq := Q{}
		for qi := range q {
			x, y := qi[0], qi[1]
			// check around qi:
			for _, move := range [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} {

				// candidate for next step
				cx, cy := x+move[0], y+move[1]

				// stone?
				if !f[cy][cx] {
					continue
				}

				// otherwise legal next step
				nq[[2]int{cx, cy}] = true
			}
		}

		q = maps.Clone(nq)
		N++
		res = append(res, len(nq))
		if N == steps {
			return res
		}
	}
	return []int{}
}

func AllEqual(a []int, s int) bool {
	for i := range a {
		if a[i] != s {
			return false
		}
	}
	return true
}

// prediction func from Day 9 Part 1
func predict(a []int) int {
	if len(a) == 0 || AllEqual(a, 0) {
		return 0
	}

	nexta := []int{}
	for i := 0; i < len(a)-1; i++ {
		nexta = append(nexta, a[i+1]-a[i])
	}

	return a[len(a)-1] + predict(nexta)
}

func main() {
	fmt.Println("Day 21: Step Counter")

	in, _ := os.ReadFile("../aoc-inputs/2023/21/input.txt")
	lines, sx, sy := readLines(string(in), 5) // 5x5 tiles

	if slices.Index(os.Args[1:], "--visual") != -1 {
		showLines(readLines(string(in), 3))
	}
	res := reach(lines, sx, sy, 64)
	fmt.Println("\tPart One:", res[64]) // 3646

	// fmt.Println("Reaching", 65+0*131, 65+1*131, "and", 65+2*131, "steps to build predictions for further steps:")
	res = reach(lines, sx, sy, 65+2*131)

	a := []int{
		res[0*131+65], // 3759
		res[1*131+65], // 33496
		res[2*131+65], // 92857
	}

	// i.e.: a[3*131+65] = predict[a]

	// how many garden plots could the Elf reach in exactly 26501365 steps?
	// 26501365 = 202300 * 131 + 62

	// fmt.Println("Predicting steps till N == 202300")

	for N := 2; N < 202300; N++ {
		a = append(a, predict(a[max(0, len(a)-4):]))
	}

	fmt.Println("\tPart Two:", a[len(a)-1]) // 606188414811259
}
