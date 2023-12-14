package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/maps"
)

type set struct {
	lines []string
	cols  []string
}

func read(f string) (sets []set) {
	in, _ := os.ReadFile(f)
	for _, m := range strings.Split(string(in), "\n\n") {
		lines := []string{}
		cols := []string{}

		lines = strings.Split(strings.TrimSpace(string(m)), "\n")

		for i := range lines {
			for j, c := range lines[i] {
				if len(cols) == 0 {
					cols = make([]string, len(lines[i]))
				}
				cols[j] += string(c)
			}
		}

		sets = append(sets, set{lines: lines, cols: cols})
	}
	return sets
}

// returns true if a slice of lines can be folded in the middle
func reflected(lines []string) bool {
	if len(lines)%2 != 0 || len(lines) == 0 {
		return false
	}

	for i := range lines[:len(lines)/2] {
		if lines[i] != lines[len(lines)-1-i] {
			return false
		}
	}

	return true
}

func findMirrors(s []string) map[int]bool {
	mirrors := map[int]bool{}
	for i := range s {
		if reflected(s[0:i]) {
			mirrors[(i+1)/2] = true
			// fmt.Println(0, "x", i)
		}

		if reflected(s[i:]) {
			mirrors[len(s)-(len(s)-i)/2] = true
			// fmt.Println(i, "x", len(s))
		}

	}
	return mirrors
}

func p1(f string) int {
	sets := read(f)
	sum := 0
	for _, s := range sets {
		m := findMirrors(s.lines)
		for _, r := range maps.Keys(m) {
			sum += r * 100
		}
		// fmt.Println(m)
		m = findMirrors(s.cols)
		for _, r := range maps.Keys(m) {
			sum += r
		}
		// fmt.Println(m)
	}
	// fmt.Println(sum)
	return sum
}

func main() {
	fmt.Println("Day 13: Point of Incidence")
	fmt.Println("\tPart One:", p1("../aoc-inputs/2023/13/input.txt")) // 36041
}
