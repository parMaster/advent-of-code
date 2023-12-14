package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type set struct {
	lines []string
	cols  []string
}

func read(in string) set {
	lines := []string{}
	cols := []string{}

	lines = strings.Split(strings.TrimSpace(string(in)), "\n")
	for i := range lines {
		for j, c := range lines[i] {
			if len(cols) == 0 {
				cols = make([]string, len(lines[i]))
			}
			cols[j] += string(c)
		}
	}
	return set{lines: lines, cols: cols}
}

func shift(s []byte) int {
	w := 0
	total := len(s)
	for _, part := range strings.Split(strings.TrimSpace(string(s)), "#") {
		pp := []byte(part)
		if len(pp) != 0 {
			slices.Sort(pp)
			slices.Reverse(pp)
			s = slices.DeleteFunc(pp, func(e byte) bool { return e == byte('.') })
			// rocks := len(s)
			// dots := len(pp) - rocks
		}

		for _, p := range pp {
			if p == byte('O') {
				w += total
			}
			total--
			fmt.Print(string(p))
		}
		total--
		fmt.Print("#")
		// OOOO.#.O.. 10
		// OO..#....#  9
		// OO..O##..O  8
		// O..#.OO...  7
		// ........#.  6
		// ..#....#.#  5
		// ..O..#.O.O  4
		// ..O.......  3
		// #....###..  2
		// #....#....  1
		// fmt.Println(string(s), total, w)
	}
	fmt.Println()
	return w
}

func solve(f string) int {
	in, _ := os.ReadFile(f)
	sum := 0
	for _, s := range read(string(in)).cols {
		sum += shift([]byte(s))
	}
	return sum
}

func main() {
	fmt.Println("Day 14: Parabolic Reflector Dish")
	fmt.Println("\tPart One:", solve("../aoc-inputs/2023/14/input.txt")) //
	// fmt.Println("\tPart Two:", solve("../aoc-inputs/2023/14/input.txt", 1)) //
}
